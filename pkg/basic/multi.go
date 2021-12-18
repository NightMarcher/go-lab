package basic

import (
	"fmt"
	"sync"
	"time"
)

func manualWorker(id int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "is working", j)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("worker", id, "finished job", j)
		res <- j * j
	}
}

// pool
func DemoPool() {
	fmt.Println()
	const numJobs = 5
	jobs := make(chan int, numJobs)
	res := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go manualWorker(w, jobs, res)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for i := 0; i < numJobs; i++ {
		<-res
	}
}

func autoWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

// WaitGroup
func DemoWaitGroup() {
	fmt.Println()
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go autoWorker(i, &wg)
	}
	wg.Wait()
}
