package router

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewRouter(t *testing.T) {
	notFound := func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request, params url.Values) {
		return
	}

	type args struct {
		notFoundHandler Handler
	}
	tests := []struct {
		name string
		args args
		want *Router
	}{
		{
			name: "Test correct new router",
			args: args{
				notFoundHandler: notFound,
			},
			want: &Router{
				tree: &node{
					component:          "root",
					isNamedParam:       false,
					httpMethodHandlers: make(map[string]Handler),
				},

				notFoundHandler: notFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRouter(tt.args.notFoundHandler)
			want := tt.want
			if !reflect.DeepEqual(*got.tree, *want.tree) || got.notFoundHandler == nil {
				t.Errorf("NewRouter() = Tree: %v, want Tree: %v\n", *got.tree, *want.tree)
			}
		})
	}
}

func TestRouter_ServeHTTP(t *testing.T) {
	type args struct {
		httpResponseWriter http.ResponseWriter
		httpRequest        *http.Request
	}
	tests := []struct {
		name       string
		thisRouter *Router
		args       args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.thisRouter.ServeHTTP(tt.args.httpResponseWriter, tt.args.httpRequest)
		})
	}
}

func TestRouter_Handle(t *testing.T) {
	type args struct {
		httpMethod string
		path       string
		handler    Handler
	}
	tests := []struct {
		name       string
		thisRouter *Router
		args       args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.thisRouter.Handle(tt.args.httpMethod, tt.args.path, tt.args.handler)
		})
	}
}

func TestRouter_getHandler(t *testing.T) {
	type args struct {
		node       *node
		httpMethod string
	}
	tests := []struct {
		name       string
		thisRouter *Router
		args       args
		want       Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.thisRouter.getHandler(tt.args.node, tt.args.httpMethod); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router.getHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeDefaultNode(t *testing.T) {
	type args struct {
		component string
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeDefaultNode(tt.args.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeDefaultNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_addNode(t *testing.T) {
	type args struct {
		httpMethod string
		path       string
		handler    Handler
	}
	tests := []struct {
		name     string
		thisNode *node
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.thisNode.addNode(tt.args.httpMethod, tt.args.path, tt.args.handler)
		})
	}
}

func Test_node_addAllChildrenAndReturnFinalNode(t *testing.T) {
	type args struct {
		components []string
	}
	tests := []struct {
		name     string
		thisNode *node
		args     args
		want     *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.thisNode.addAllChildrenAndReturnFinalNode(tt.args.components); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.addAllChildrenAndReturnFinalNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_traverseTree(t *testing.T) {
	type args struct {
		components []string
		params     url.Values
	}
	tests := []struct {
		name     string
		thisNode *node
		args     args
		want     *node
		want1    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.thisNode.traverseTree(tt.args.components, tt.args.params)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.traverseTree() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("node.traverseTree() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_node_getValidChildAndAddParams(t *testing.T) {
	type args struct {
		component string
		params    url.Values
	}
	tests := []struct {
		name     string
		thisNode *node
		args     args
		want     *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.thisNode.getValidChildAndAddParams(tt.args.component, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.getValidChildAndAddParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_doesMatchComponent(t *testing.T) {
	type args struct {
		component string
	}
	tests := []struct {
		name     string
		thisNode *node
		args     args
		want     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.thisNode.doesMatchComponent(tt.args.component); got != tt.want {
				t.Errorf("node.doesMatchComponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeThisAndPrecedingElements(t *testing.T) {
	type args struct {
		element string
		array   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeThisAndPrecedingElements(tt.args.element, tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeThisAndPrecedingElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
