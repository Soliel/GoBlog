package router

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewRouter(t *testing.T) {
	notFound := func(w http.ResponseWriter, r *http.Request, params url.Values) {
		return
	}

	type args struct {
		notFoundHandler Handle
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
					component:    "root",
					isNamedParam: false,
					methods:      make(map[string]Handle),
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
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		r    *Router
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.ServeHTTP(tt.args.w, tt.args.req)
		})
	}
}

func TestRouter_Handle(t *testing.T) {
	type args struct {
		method  string
		path    string
		handler Handle
	}
	tests := []struct {
		name string
		r    *Router
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Handle(tt.args.method, tt.args.path, tt.args.handler)
		})
	}
}

func TestRouter_getHandler(t *testing.T) {
	type args struct {
		node   *node
		method string
	}
	tests := []struct {
		name string
		r    *Router
		args args
		want Handle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.getHandler(tt.args.node, tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router.getHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeNewNode(t *testing.T) {
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
			if got := makeNewNode(tt.args.component); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeNewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_addNode(t *testing.T) {
	type args struct {
		method  string
		path    string
		handler Handle
	}
	tests := []struct {
		name string
		n    *node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.addNode(tt.args.method, tt.args.path, tt.args.handler)
		})
	}
}

func Test_node_addAllChildrenAndReturnFinalNode(t *testing.T) {
	type args struct {
		components []string
	}
	tests := []struct {
		name string
		n    *node
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.addAllChildrenAndReturnFinalNode(tt.args.components); !reflect.DeepEqual(got, tt.want) {
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
		name  string
		n     *node
		args  args
		want  *node
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.n.traverseTree(tt.args.components, tt.args.params)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.traverseTree() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("node.traverseTree() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_node_addParamIfNamedParam(t *testing.T) {
	type args struct {
		params    url.Values
		component string
	}
	tests := []struct {
		name string
		n    *node
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.addParamIfNamedParam(tt.args.params, tt.args.component)
		})
	}
}

func Test_node_doesMatchComponent(t *testing.T) {
	type args struct {
		component string
	}
	tests := []struct {
		name string
		n    *node
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.doesMatchComponent(tt.args.component); got != tt.want {
				t.Errorf("node.doesMatchComponent() = %v, want %v", got, tt.want)
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
		name string
		n    *node
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.getValidChildAndAddParams(tt.args.component, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.getValidChildAndAddParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removePrecedingElements(t *testing.T) {
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
			if got := removePrecedingElements(tt.args.element, tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removePrecedingElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
