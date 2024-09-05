/*
	PACKAGES

- package is a collection of related Go files that help organize code and promote reuse
- packages provide a namespace to avoid naming conflicts between functions variables and other identifiers
- packages can be imported into other Go programs using 'import' keyword
- the main package (package main) s used for executable programs while other packages provide reusable libraries
*/
package main

/* IMPORTS
- used to include external or standard library packages in Go programs
- this allows for reusability of pre-build functionality and code

*/
import (
	"context"       // Provides context for managing timeouts and cancellations
	"encoding/json" // Provides functions for encoding and decoding JSON data
	"errors"        // Provides functions to create and manipulate errors
	"fmt"           // Provides formatted I/O functions
	"io"            // Provides basic I/O primitives
	"log"           // Provides simple logging functions
	"math"          // Provides basic constants and mathematical functions
	"net/http"      // Provides HTTP client and server implementations
	"os"            // Provides functions for interacting with the operating system
	"strings"       // Provides functions for manipulating strings
	"sync"          // Provides concurrency primitives like Mutex and WaitGroup
	"time"          // Provides functionality for working with dates and times
)

func main() {
	fmt.Println("Hello, Go!")               // Print a message to the console
	errors.New("error message")             // Create a new error with a message
	math.Sqrt(16)                           // Compute the square root of 16
	time.Now()                              // Get the current date and time
	http.ListenAndServe(":8080", nil)       // Start an HTTP server on port 8080
	os.Open("file.txt")                     // Open a file named "file.txt"
	strings.Contains("GoLang", "Go")        // Check if the string "Go" is contained in "GoLang"
	io.Copy(dst, src)                       // Copy data from src to dst
	sync.Mutex{}                            // Create a new Mutex for synchronization
	context.WithTimeout(ctx, time.Second*5) // Create a context with a 5-second timeout
	json.Marshall(struct{})                 // (Note: should be `json.Marshal`, this is a typo) Encode a struct to JSON
	log.Println("Error occurred")           // Log an error message
}
