package router

import (
	"net/http"
	"net/url"
	"strings"
)

//Handle helps us reference Handler functions easier.
type Handle func(http.ResponseWriter, *http.Request, url.Values)

//Router holds our routing trie.
type Router struct {
	tree            *node
	notFoundHandler Handle
}

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]Handle
}

//NewRouter creates a new router and registers a redirect URL to be used in case we can't find a path. Generally a 404 page.
func NewRouter(notFoundHandler Handle) *Router {
	node := node{
		component:    "root",
		isNamedParam: false,
		methods:      make(map[string]Handle),
	}

	return &Router{
		tree:            &node,
		notFoundHandler: notFoundHandler,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form

	handleArray := strings.Split(req.URL.Path, "/")

	if handleArray[1] == "" {
		handler := r.getHandler(r.tree, req.Method)
		go handler(w, req, params)
	} else {
		n, _ := r.tree.traverseTree(handleArray[1:], params)

		handler := r.getHandler(n, req.Method)
		go handler(w, req, params)
	}
}

//Handle registers routes into the router.
func (r *Router) Handle(method, path string, handler Handle) {
	if path[0] != '/' {
		panic("Path must start with a /")
	}
	r.tree.addNode(method, path, handler)
}

func (r *Router) getHandler(node *node, method string) Handle {
	handler := node.methods[method]
	if handler != nil {
		return handler
	} else {
		return r.notFoundHandler
	}
}

func makeNewNode(component string) *node {
	return &node{
		component:    component,
		isNamedParam: false,
		methods:      make(map[string]Handle),
	}
}

func (n *node) addNode(method string, path string, handler Handle) {
	components := strings.Split(path, "/")[1:]
	if components[0] == "" {
		if n.component == "root" {
			n.methods[method] = handler
		}
	}

	aNode, component := n.traverseTree(components, nil)
	components = removePrecedingElements(component, components)
	lastNode := aNode.addAllChildrenAndReturnFinalNode(components)
	lastNode.methods[method] = handler

}

func (n *node) addAllChildrenAndReturnFinalNode(components []string) *node {
	if len(components) <= 0 {
		return n
	}

	component := components[0]
	newNode := makeNewNode(component)
	if len(component) > 0 && component[0] == ':' {
		newNode.isNamedParam = true
	}

	nextComponents := components[1:]
	return newNode.addAllChildrenAndReturnFinalNode(nextComponents)
}

func (n *node) traverseTree(components []string, params url.Values) (*node, string) {
	component := components[0]
	if len(n.children) <= 0 {
		return n, component
	}

	validChild := n.getValidChildAndAddParams(component, params)
	if validChild == nil {
		return n, component
	}

	nextComponents := components[1:]
	if len(nextComponents) > 0 {
		return validChild.traverseTree(nextComponents, params)
	}

	return n, component
}

func (n *node) addParamIfNamedParam(params url.Values, component string) {
	if n.isNamedParam {
		params.Add(n.component[1:], component)
	}
}

func (n *node) doesMatchComponent(component string) bool {
	if n.component == component {
		return true
	}

	return false
}

func (n *node) getValidChildAndAddParams(component string, params url.Values) *node {
	for _, child := range n.children {
		if child.doesMatchComponent(component) || child.isNamedParam {
			child.addParamIfNamedParam(params, component)
			return child
		}
	}

	return nil
}

func removePrecedingElements(element string, array []string) []string {
	for index, value := range array {
		if value == element {
			return array[index+1:]
		}
	}

	return nil
}
