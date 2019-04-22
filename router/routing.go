package router

import (
	"net/http"
	"net/url"
	"strings"
)

//Handler helps us reference Handler functions easier.
type Handler func(http.ResponseWriter, *http.Request, url.Values)

//Router holds our routing trie.
type Router struct {
	tree            *node
	notFoundHandler Handler
}

type node struct {
	children           []*node
	component          string
	isNamedParam       bool
	httpMethodHandlers map[string]Handler
}

//NewRouter creates a new router and registers a redirect URL to be used in case we can't find a path. Generally a 404 page.
func NewRouter(notFoundHandler Handler) *Router {
	node := node{
		component:          "root",
		isNamedParam:       false,
		httpMethodHandlers: make(map[string]Handler),
	}

	return &Router{
		tree:            &node,
		notFoundHandler: notFoundHandler,
	}
}

func (thisRouter *Router) ServeHTTP(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	httpRequest.ParseForm()
	params := httpRequest.Form

	handleArray := strings.Split(httpRequest.URL.Path, "/")

	if handleArray[1] == "" {
		handler := thisRouter.getHandler(thisRouter.tree, httpRequest.Method)
		go handler(httpResponseWriter, httpRequest, params)
	} else {
		n, _ := thisRouter.tree.traverseTree(handleArray[1:], params)

		handler := thisRouter.getHandler(n, httpRequest.Method)
		go handler(httpResponseWriter, httpRequest, params)
	}
}

//Handle registers routes into the router.
func (thisRouter *Router) Handle(httpMethod string, path string, handler Handler) {
	if path[0] != '/' {
		panic("Path must start with a /")
	}

	thisRouter.tree.addNode(httpMethod, path, handler)
}

func (thisRouter *Router) getHandler(node *node, httpMethod string) Handler {
	handler := node.httpMethodHandlers[httpMethod]
	if handler != nil {
		return handler
	} else {
		return thisRouter.notFoundHandler
	}
}

func makeDefaultNode(component string) *node {
	return &node{
		component:          component,
		isNamedParam:       false,
		httpMethodHandlers: make(map[string]Handler),
	}
}

func (thisNode *node) addNode(httpMethod string, path string, handler Handler) {
	componentsWithoutLeadingZero := strings.Split(path, "/")[1:]
	if componentsWithoutLeadingZero[0] == "" {
		if thisNode.component == "root" {
			thisNode.httpMethodHandlers[httpMethod] = handler
			return
		}
	}

	lastNodeInTree, componentAtLastNode := thisNode.traverseTree(componentsWithoutLeadingZero, nil)
	componentsAfterLastNode := removeThisAndPrecedingElements(componentAtLastNode, componentsWithoutLeadingZero)
	newNode := lastNodeInTree.addAllChildrenAndReturnFinalNode(componentsAfterLastNode)
	newNode.httpMethodHandlers[httpMethod] = handler
}

func (thisNode *node) addAllChildrenAndReturnFinalNode(components []string) *node {
	if len(components) <= 0 {
		return thisNode
	}

	firstComponent := components[0]
	newNode := makeDefaultNode(firstComponent)
	if len(firstComponent) > 0 && firstComponent[0] == ':' {
		newNode.isNamedParam = true
	}

	nextComponents := components[1:]
	return newNode.addAllChildrenAndReturnFinalNode(nextComponents)
}

func (thisNode *node) traverseTree(components []string, params url.Values) (*node, string) {
	firstComponent := components[0]
	if len(thisNode.children) <= 0 {
		return thisNode, firstComponent
	}

	validChild := thisNode.getValidChildAndAddParams(firstComponent, params)
	if validChild == nil {
		return thisNode, firstComponent
	}

	nextComponents := components[1:]
	if len(nextComponents) > 0 {
		return validChild.traverseTree(nextComponents, params)
	}

	return validChild, firstComponent
}

func (thisNode *node) getValidChildAndAddParams(component string, params url.Values) *node {
	for _, child := range thisNode.children {
		if child.doesMatchComponent(component) || child.isNamedParam {
			return child
		} else if child.isNamedParam {
			params.Add(thisNode.component[1:], component)
			return child
		}
	}

	return nil
}

func (thisNode *node) doesMatchComponent(component string) bool {
	if thisNode.component == component {
		return true
	}

	return false
}

func removeThisAndPrecedingElements(element string, array []string) []string {
	for index, value := range array {
		if value == element {
			return array[index+1:]
		}
	}

	return nil
}
