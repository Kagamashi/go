package main

import (
    "fmt"
    "time"
)

/* WORKER POOLS


*/

/* worker of which we'll run several concurrent instances
these worker will receive work on the jobs channel
and send the corresponding results on results channel  */
func worker_p(id int, jobs <-chan int, results chan<- int) { 
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func worker_pools() {

    const numJobs = 5  // we create 2 channels for sending and receiving job results
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {  //start up 3 workers, initially blocked because there are no jobs yet
        go worker_p(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {  // here we send 5 jobs and then close that channel to indicate that's all the work we have
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ {  // we collect all the results of the work
        <-results  // this also ensures that the worker goroutines have finished; an alternative would be to use WaitGroups
    }
		/*	worker 1 started	job 1
				worker 2 started	job 2


		*/


		/* Our running program shows the 5 jobs being executed by various worker.
		The program only takes about 2 seconds despite doing about 5 seconds of total work because there are 3 workers operating concurrently */
}
