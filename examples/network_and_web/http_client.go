package main

import (
	"bufio"
	"fmt"
	"net/http"
)

/* HTTP CLIENT
net/http package provides built-in HTTP client to send HTTP requests and receive responses from servers
- http.Get/http.Post - Simple GET/POST requests
- http.NewRequest - custom requests, more control over headers, methods and request bodies
- http.Client - timeouts
*/

func http_client() {

	resp, err := http.Get("https://gobyexample.com") // Issue HTTP GET request to a server
	if err != nil {                                  // http.Get is a convenient shortcut around creating http.Client object and calling it Get method (it uses http.DefaultClient object with default settings)
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status) // Print the HTTP response status

	scanner := bufio.NewScanner(resp.Body) // Print the first 5 lines of the response body
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	/*
		Response status: 200 OK
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8">
				<title>Go by Example</title>
	*/

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
