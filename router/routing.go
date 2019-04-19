package router

import (
	"net/http"
	"net/url"
)

//Handle helps us reference Handler functions easier.
type Handle func(http.ResponseWriter, *http.Request, url.Values)

//Router holds our routing trie.
type Router struct {
	tree     *node
	notFound  Handle
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
		component:    "/",
		isNamedParam: false,
		methods:      make(map[string]Handle),
	}

	return &Router{
		tree:     &node,
		notFound: notFoundHandler,
	}
}

func (n *node) addNode(method string, path string, handler Handle) {

}

func (n *node) traverse(components []string, params url.Values) (*node, string){
	component := components[0]
	if len(n.children <= 0 && n.component == component) return n, component
	
}
