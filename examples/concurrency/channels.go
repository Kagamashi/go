package main

import "fmt"

/* CHANNELS



 */

func channels() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
