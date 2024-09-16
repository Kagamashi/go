package main

import (
	"fmt"
	"net/http"
	"time"
)

/* CONTEXT
HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation.
A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.
*/

func hello_c(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context() // context.Context is created for each request by the net/http machinery and is available with the Context() method
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select { // Wait for few seconds before sending a reply to the client
	case <-time.After(10 * time.Second): // This could simulate some work the server is doing
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err() // Contexts Err() method returns an error that explains why Done() channel was closed
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func context() {

	http.HandleFunc("/hello", hello_c) // Register our handler on "/hello" route
	http.ListenAndServe(":8090", nil)

	/* Output:
	$ go run context.go &
	$ curl localhost:8090/hello
	server: hello handler started
	^C // Simulate client request to /hello by hitting Ctrl+C shortly after starting the server
	server: context canceled
	server: hello handler ended
	*/

}
