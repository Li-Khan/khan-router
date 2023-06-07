package khanRouter

// Defines the structure for individual routes and allows adding middleware to them.

import (
	"net/http"
	"sync"
)

var (
	// routesMutex is a mute used to ensure the safety of accessing the shared routesMap map during write operations.
	routesMutex = &sync.Mutex{}
	// routesMap store routes.
	routesMap = make(map[string]map[string]*Route)
)

// Route structure stores information related to a registered route
// and provides functionality for adding middleware to the route.
type Route struct {
	// handler - The main handler function for the route.
	handler http.Handler
	// original - The original handler function before applying any middleware.
	original http.Handler
	// pattern -  The path pattern for the route.
	pattern string
	// method - The HTTP method associated with the route.
	method string
	// middlewares - A list of middleware functions to be applied to the route's handler.
	middlewares []func(handler http.Handler) http.Handler
	// group - The route group that the route belongs to (if any).
	group *GroupRoute
}

// Middleware adds one or more middleware functions to the route.
// These middleware functions will be executed before the main handler function.
// Parameters:
// m (...func(handler http.Handler) http.Handler): One or more middleware functions to be added to the route.
func (r *Route) Middleware(m ...func(handler http.Handler) http.Handler) {
	r.handler = r.original
	r.middlewares = append(r.middlewares, m...)
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		r.handler = r.middlewares[i](r.handler)
	}
}
