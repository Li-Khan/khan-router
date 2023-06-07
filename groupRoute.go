package khanRouter

// Implements the functionality for creating and registering route groups with shared middleware.

import "net/http"

// GroupRoute structure represents a group of routes with a common path prefix and shared middleware functions.
type GroupRoute struct {
	// pattern - The path pattern prefix for the group of routes.
	pattern string
	// middleware - A list of middleware functions to be applied to all routes in the group.
	middlewares []func(handler http.Handler) http.Handler
}

// RegisterGroupRoute registers a new route group with a specified pattern and optional middleware functions.
// Parameters:
// pattern (string): The base path pattern for the route group.
// m (...func(handler http.Handler) http.Handler): Optional middleware functions to be applied to all routes in the group.
func RegisterGroupRoute(pattern string, m ...func(handler http.Handler) http.Handler) *GroupRoute {
	return &GroupRoute{
		pattern:     getValidPattern(pattern),
		middlewares: m,
	}
}

// RegisterGroupRoute registers a nested route group under the current group,
// inheriting the parent group's pattern and middleware functions.
func (gr *GroupRoute) RegisterGroupRoute(pattern string, m ...func(handler http.Handler) http.Handler) *GroupRoute {
	return &GroupRoute{
		pattern:     gr.pattern + getValidPattern(pattern),
		middlewares: append(gr.middlewares, m...),
	}
}

// RegisterRouteGET registers a new route with the HTTP GET method under the current group,
// inheriting the parent group's pattern and middleware functions.
func (gr *GroupRoute) RegisterRouteGET(pattern string, handler http.HandlerFunc) *Route {
	pattern = gr.pattern + getValidPattern(pattern)

	route := &Route{
		handler:     handler,
		original:    handler,
		pattern:     pattern,
		method:      http.MethodGet,
		group:       gr,
		middlewares: gr.middlewares,
	}

	for i := len(route.middlewares) - 1; i >= 0; i-- {
		route.handler = route.middlewares[i](route.handler)
	}

	registerRoute(pattern, http.MethodGet, route)

	return route
}

// RegisterRoutePOST registers a new route with the HTTP POST method under the current group,
// inheriting the parent group's pattern and middleware functions.
func (gr *GroupRoute) RegisterRoutePOST(pattern string, handler http.HandlerFunc) *Route {
	pattern = gr.pattern + getValidPattern(pattern)

	route := &Route{
		handler:     handler,
		original:    handler,
		pattern:     pattern,
		method:      http.MethodPost,
		group:       gr,
		middlewares: gr.middlewares,
	}

	for i := len(route.middlewares) - 1; i >= 0; i-- {
		route.handler = route.middlewares[i](route.handler)
	}

	registerRoute(pattern, http.MethodPost, route)

	return route
}

// RegisterRoutePUT registers a new route with the HTTP PUT method under the current group,
// inheriting the parent group's pattern and middleware functions.
func (gr *GroupRoute) RegisterRoutePUT(pattern string, handler http.HandlerFunc) *Route {
	pattern = gr.pattern + getValidPattern(pattern)

	route := &Route{
		handler:     handler,
		original:    handler,
		pattern:     pattern,
		method:      http.MethodPut,
		group:       gr,
		middlewares: gr.middlewares,
	}

	for i := len(route.middlewares) - 1; i >= 0; i-- {
		route.handler = route.middlewares[i](route.handler)
	}

	registerRoute(pattern, http.MethodPut, route)

	return route
}

// RegisterRouteDELETE registers a new route with the HTTP DELETE method under the current group,
// inheriting the parent group's pattern and middleware functions.
func (gr *GroupRoute) RegisterRouteDELETE(pattern string, handler http.HandlerFunc) *Route {
	pattern = gr.pattern + getValidPattern(pattern)

	route := &Route{
		handler:     handler,
		original:    handler,
		pattern:     pattern,
		method:      http.MethodDelete,
		group:       gr,
		middlewares: gr.middlewares,
	}

	for i := len(route.middlewares) - 1; i >= 0; i-- {
		route.handler = route.middlewares[i](route.handler)
	}

	registerRoute(pattern, http.MethodDelete, route)

	return route
}

// RegisterRouteOPTIONS registers a new route with the HTTP OPTIONS method under the current group,
// inheriting the parent group's pattern and middleware functions.
func (gr *GroupRoute) RegisterRouteOPTIONS(pattern string, handler http.HandlerFunc) *Route {
	pattern = gr.pattern + getValidPattern(pattern)

	route := &Route{
		handler:     handler,
		original:    handler,
		pattern:     pattern,
		method:      http.MethodOptions,
		group:       gr,
		middlewares: gr.middlewares,
	}

	for i := len(route.middlewares) - 1; i >= 0; i-- {
		route.handler = route.middlewares[i](route.handler)
	}

	registerRoute(pattern, http.MethodOptions, route)

	return route
}
