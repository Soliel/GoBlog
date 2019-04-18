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
	redirect string
}

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]Handle
}

//NewRouter creates a new router and registers a redirect URL to be used in case we can't find a path. Generally a 404 page.
func NewRouter(redirectURL string) *Router {
	node := node{
		component:    "root",
		isNamedParam: false,
		methods:      make(map[string]Handle),
	}

	return &Router{
		tree:     &node,
		redirect: redirectURL,
	}
}

func (n *node) addNode(method string, path string, handler Handle) {

}

func (n *node) traverse()
