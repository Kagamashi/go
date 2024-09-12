package main

import (
	"fmt"
	"sync"
)

/* MUTEXES
Mutexes (Mutual Exclusions) are used to safely access shared data across multiple goroutines by allowing only one goroutine (locks) to access a critical section of code at a time.
- Using sync.Mutex provides Lock() and Unlock() methods to control access to shared resources.
var mu sync.Mutex
var counter int
mu.Lock()
counter++ // Critical section
mu.Unlock()

- syncRWMutex is a variant that allows multiple goroutines to read but only one to write.
RLock() and RUnlock()
- Used to protect shared variables, counters, resources from race conditions when accessed by multiple goroutines
*/

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func mutexes() {
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
