package main

import "fmt"

/* SELECT DEFAULT
Basic sends and receives on channels are blocking.
However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
*/

func select_default() {
	messages := make(chan string)
	signals := make(chan bool)

	select { //non-blocking receive
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	// no message received

	msg := "hi"
	select { // non-blocking send
	case messages <- msg: // msg cannot be sent to the messages channel, because it has no buffer and there is no receiver
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	// no message sent

	select { // we can use multiple cases above the default clause to implement a multi-way non-blocking select
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
	// no activity

}
