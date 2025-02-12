package main

import (
	"fmt"
	"net/http"
)

/* HTTP SERVER
net/http package also allows you to create HTTP servers to handle incoming HTTP requests
- http.HandleFunc - registers handlers to respond to different URL patterns
- http.ListenAndServe - starts the server on a specified port
- http.Request - to access incoming request data
- http.ResponseWriter - send responses
- encoding/json to send JSON responses
- http.Server for advanced settings: timeouts, TLS, custom middleware
*/

/*
	Handler is an object implementing http.Handler interface

Common way to write handler is by using http.HandlerFunc adapter on functions with appropriate signature
Functions serving as handlers take http.Responsewriter and http.Request as arguments
Response writer is used to fill in the HTTP response
*/
func hello(w http.ResponseWriter, req *http.Request) { // req is incoming HTTP request
	fmt.Fprintf(w, "hello\n")
}

// This handler reads all HTTP request headers and echoing them into the response body
func headers(w http.ResponseWriter, req *http.Request) { // responds to requests made to the /headers URL path

	for name, headers := range req.Header { // iterates over all headers in incoming HTTP request
		for _, h := range headers { // for each header it writes the header name and value to HTTP response
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func http_server() {

	http.HandleFunc("/hello", hello)     // Registering handlers on server routes using http.HandleFunc function
	http.HandleFunc("/headers", headers) // It sets up default router in net/http package and takes a function as an argument

	http.ListenAndServe(":8090", nil) // Finally we call ListenAndServe with port and a handler, nil tells it to use default router we set up with previous commands

	/*
		$ go run http_server.go & // Run server in the background
		$ curl localhost:8090/hello
		hello
	*/

}
