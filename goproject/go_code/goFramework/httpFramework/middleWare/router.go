package middleWare

import "strings"

//import "go/ast"

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {

	return &router{

		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}

}

func paresPattern(pattern string) []string {

	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {

			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts

}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {

	parts := paresPattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}
