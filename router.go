package khanRouter

// Contains the main router implementation and the logic for routing incoming requests.

import (
	"log"
	"net/http"
)

// Router structure implements the http.Handler interface
// and handles incoming HTTP requests by matching the request's path and method to registered routes.
type Router struct {
	// NotFoundHandler The handler to be called when no matching route is found for a requested path.
	NotFoundHandler http.Handler
	// MethodNotAllowedHandler The handler to be called when a route exists for the requested path
	// but does not support the specified HTTP method.
	MethodNotAllowedHandler http.Handler
}

// NewRouter creates a new instance of the Router structure with default handlers for not found routes and method not allowed responses.
func NewRouter() *Router {
	return &Router{
		NotFoundHandler:         http.NotFoundHandler(),
		MethodNotAllowedHandler: http.HandlerFunc(methodNotAllowedHandler),
	}
}

// ServeHTTP Implements the http.Handler interface for the Router structure
// and handles incoming HTTP requests by calling the appropriate handler based on the request's path and method.
func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.getHandler(request).ServeHTTP(writer, request)
}

// getHandler Retrieves the appropriate handler for a given HTTP request by matching the request's path
// and method against the registered routes.
// If no matching route is found, it returns the NotFoundHandler.
// If a route exists but does not support the requested method, it returns the MethodNotAllowedHandler.
func (r *Router) getHandler(request *http.Request) http.Handler {
	routes, ok := routesMap[request.URL.Path]
	if !ok {
		return r.NotFoundHandler
	}
	route, ok := routes[request.Method]
	if !ok {
		return r.MethodNotAllowedHandler
	}
	return route.handler
}

// methodNotAllowedHandler Handles the case when a route exists for a requested path
// but does not support the specified HTTP method. It returns a "409 Method Not Allowed" response.
func methodNotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "409 method not allowed", http.StatusMethodNotAllowed)
}

// RegisterRouteGET registers a new route with the HTTP GET method.
func (r *Router) RegisterRouteGET(pattern string, handler http.HandlerFunc) *Route {
	pattern = getValidPattern(pattern)
	route := &Route{
		handler:  handler,
		original: handler,
		pattern:  pattern,
		method:   http.MethodGet,
	}

	registerRoute(pattern, http.MethodGet, route)

	return route
}

// RegisterRoutePOST registers a new route with the HTTP POST method.
func (r *Router) RegisterRoutePOST(pattern string, handler http.HandlerFunc) *Route {
	pattern = getValidPattern(pattern)
	route := &Route{
		handler:  handler,
		original: handler,
		pattern:  pattern,
		method:   http.MethodPost,
	}

	registerRoute(pattern, http.MethodPost, route)

	return route
}

// RegisterRoutePUT registers a new route with the HTTP PUT method.
func (r *Router) RegisterRoutePUT(pattern string, handler http.HandlerFunc) *Route {
	pattern = getValidPattern(pattern)
	route := &Route{
		handler:  handler,
		original: handler,
		pattern:  pattern,
		method:   http.MethodPut,
	}

	registerRoute(pattern, http.MethodPut, route)

	return route
}

// RegisterRouteDELETE registers a new route with the HTTP DELETE method.
func (r *Router) RegisterRouteDELETE(pattern string, handler http.HandlerFunc) *Route {
	pattern = getValidPattern(pattern)
	route := &Route{
		handler:  handler,
		original: handler,
		pattern:  pattern,
		method:   http.MethodDelete,
	}

	registerRoute(pattern, http.MethodDelete, route)

	return route
}

// RegisterRouteOPTIONS registers a new route with the HTTP OPTIONS method.
func (r *Router) RegisterRouteOPTIONS(pattern string, handler http.HandlerFunc) *Route {
	pattern = getValidPattern(pattern)
	route := &Route{
		handler:  handler,
		original: handler,
		pattern:  pattern,
		method:   http.MethodOptions,
	}

	registerRoute(pattern, http.MethodOptions, route)

	return route
}

// registerRoute registers a new route by adding it to the routesMap based on the specified pattern and HTTP method.
// It checks for duplicate routes and panics if a route with the same pattern and method already exists.
func registerRoute(pattern string, method string, route *Route) {
	routes, ok := routesMap[pattern]
	if ok {
		_, routeOk := routes[method]
		if routeOk {
			log.Panicf("route '%s %s' already exists", method, pattern)
		}
	}
	routesMutex.Lock()
	routes = make(map[string]*Route)
	routes[method] = route
	routesMap[pattern] = routes
	routesMutex.Unlock()
}
