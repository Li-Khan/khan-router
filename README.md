# Khan Router

Khan Router is a lightweight Go package that provides a simple and flexible HTTP router implementation. It allows you to register routes with different HTTP methods and handle incoming requests efficiently.

## Features

- Easy registration of routes with various HTTP methods (GET, POST, PUT, DELETE, OPTIONS)
- Support for route groups to organize related routes
- Middleware support for preprocessing requests
- Efficient routing using a map-based lookup

## Installation

Use the following command to install the HTTP Router package:

```sh
go get github.com/Li-Khan/khanRouter
```

# Usage
1. Import the HTTP Router package into your Go code:

    ```go
    import "github.com/Li-Khan/khanRouter"
    ```

2. Create a new router instance:
    ```go
    router := khanRouter.NewRouter()
    ```

3. Register routes using the available HTTP methods:

    ```go
    router.RegisterRouteGET("/users", handleGetUsers)
    router.RegisterRoutePOST("/users", handleCreateUser)
    ```

4. Implement the handler functions for your routes:

    ```go
    func handleGetUsers(w http.ResponseWriter, r *http.Request) {
        // Handle GET /users request
    }

    func handleCreateUser(w http.ResponseWriter, r *http.Request) {
        // Handle POST /users request
    }
    ```

5. Start the HTTP server and pass the router as the handler:

    ```go
    http.ListenAndServe(":8080", router)
    ```

## Route Groups

You can create route groups to organize related routes under a common base path and apply middleware to the entire group. Here's an example:

```go
group := khanRouter.RegisterGroupRoute("/api", MiddlewareAuthentication, MiddlewareJSON)
group.RegisterRouteGET("/users", handleGetUsers) // GET /api/users
group.RegisterRoutePOST("/users", handleCreateUser).Middleware(
    MiddlewareRole,
) // POST /api/users
```

## Middleware

Middleware functions can be added to individual routes or route groups. They are executed before the main handler function and can be used for request preprocessing, authentication, logging, etc.

```go
router.RegisterRouteGET("/orders", getOrders).Middleware(
    MiddlewareAuthentication,
    MiddlewareJSON,
)
```

## Contributing

Contributions are welcome! If you find any issues or want to enhance the HTTP Router package, feel free to open a pull request.