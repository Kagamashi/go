package main

import (
	"fmt"
	"sync"
	"time"
)

/* WAIT GROUPS
sync.WaitGroup is used to wait for multiple goroutines to complete.
- Creating a WaitGroup:
1. Declare sync.WaitGroup variable
2. Use wg.Add(n) to specify the number of goroutines to wait for
*/

func worker_wg(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func waitgroups() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker_wg(i)
		}()
	}

	wg.Wait()

}
