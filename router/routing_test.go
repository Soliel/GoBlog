package router

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func createTestTree() *node {
	return &node{
		component:          "root",
		isNamedParam:       false,
		httpMethodHandlers: make(map[string]Handler),
		children: []*node{
			&node{
				component:          "test1",
				isNamedParam:       false,
				httpMethodHandlers: make(map[string]Handler),
				children: []*node{
					&node{
						component:          ":testParam",
						isNamedParam:       true,
						httpMethodHandlers: make(map[string]Handler),
					},
				},
			},
		},
	}
}

//Two nodes are equal if all their fields and children are equal
func compareTwoNodes(leftNode *node, rightNode *node) bool {
	if leftNode.component != rightNode.component {
		return false
	} else if leftNode.isNamedParam != rightNode.isNamedParam {
		return false
	} else if !reflect.DeepEqual(leftNode.httpMethodHandlers, rightNode.httpMethodHandlers) {
		return false
	} else {
		leftCount := len(leftNode.children)
		rightCount := len(rightNode.children)
		if leftCount != rightCount {
			return false
		}

		if leftCount == 0 {
			return true //Base case end of recursion.
		}

		for i := 0; i < leftCount; i++ {
			if !compareTwoNodes(leftNode.children[i], rightNode.children[i]) {
				return false
			}
		}

		return true
	}
}

//We must test this helper function in order to use it.
func Test_compareTwoNodes(t *testing.T) {
	testHandler := func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request, params url.Values) {
		return
	}

	changedRightComponent := createTestTree()
	changedRightNamedParam := createTestTree()
	changedRightMethods := createTestTree()
	changedRightChildNode := createTestTree().children[0]
	changedRightComponent.component = "different"
	changedRightNamedParam.isNamedParam = true
	changedRightMethods.httpMethodHandlers[http.MethodGet] = testHandler
	changedRightChildNode.children[0].isNamedParam = true

	type args struct {
		leftNode  *node
		rightNode *node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Equivalent Nodes",
			args: args{
				leftNode:  createTestTree(),
				rightNode: createTestTree(),
			},
			want: true,
		},
		{
			name: "Test Different Component in Top Level Nodes",
			args: args{
				leftNode:  createTestTree(),
				rightNode: changedRightComponent,
			},
			want: false,
		},
		{
			name: "Test Different isNamedParam in Top Level Nodes",
			args: args{
				leftNode:  createTestTree(),
				rightNode: changedRightNamedParam,
			},
			want: false,
		},
		{
			name: "Test Different Handler Maps in Top Level Nodes",
			args: args{
				leftNode:  createTestTree(),
				rightNode: changedRightMethods,
			},
			want: false,
		},
		{
			name: "Test Different Child Nodes",
			args: args{
				leftNode:  createTestTree(),
				rightNode: changedRightChildNode,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := compareTwoNodes(tt.args.leftNode, tt.args.rightNode)
			if result != tt.want {
				t.Errorf("compareTwoNodes got: %v, want: %v", result, tt.want)
			}
		})
	}
}

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

func TestRouter_Handle(t *testing.T) {
	notFound := func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request, params url.Values) {
		return
	}

	testHandle := func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request, params url.Values) {
		return
	}

	router := NewRouter(notFound)

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
		{
			name:       "Test add handler",
			thisRouter: router,
			args: args{
				httpMethod: http.MethodGet,
				path:       "/",
				handler:    testHandle,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.thisRouter.Handle(tt.args.httpMethod, tt.args.path, false, tt.args.handler)

			splitPath := strings.Split(tt.args.path, "/")[1:]
			node, _ := tt.thisRouter.tree.traverseTree(splitPath, nil)

			gotHandler := tt.thisRouter.getHandler(node, tt.args.httpMethod)
			if fmt.Sprintf("%p", gotHandler) != fmt.Sprintf("%p", tt.args.handler) { //Because functions are never equal unless nil we just test memory address by converting to a string.
				t.Errorf("Handle() returned: %v, want: %v", gotHandler, tt.args.handler)
			}
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
		{
			name: "testCorrectNode",
			args: args{
				component: "A neat Component",
			},
			want: &node{
				component:          "A neat Component",
				isNamedParam:       false,
				httpMethodHandlers: make(map[string]Handler),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeDefaultNode(tt.args.component); !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("makeDefaultNode() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func Test_node_addNode(t *testing.T) {
	newTree := makeDefaultNode("root")

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
		{
			name:     "test Default Add Node",
			thisNode: newTree,
			args: args{
				httpMethod: http.MethodGet,
				path:       "/test/test2/test3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.thisNode.addNode(tt.args.httpMethod, tt.args.path, false, tt.args.handler)
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
	tree := createTestTree()

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
		{
			name:     "testFirstNodeTraverse",
			thisNode: tree,
			args: args{
				components: []string{
					"test1",
				},
				params: make(url.Values),
			},
			want: &node{
				component:          "test1",
				isNamedParam:       false,
				httpMethodHandlers: make(map[string]Handler),
				children: []*node{
					&node{
						component:          ":testParam",
						isNamedParam:       true,
						httpMethodHandlers: make(map[string]Handler),
					},
				},
			},
			want1: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.thisNode.traverseTree(tt.args.components, tt.args.params)
			if !compareTwoNodes(got, tt.want) {
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
			if got := removePrecedingElements(tt.args.element, tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeThisAndPrecedingElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
