package main

import (
	"fmt"
	"time"
)

func worker(workerId int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("GOROUTINE_%d: Finished file_%d.txt", workerId, job)
	}
}

func main() {
	startProcess := time.Now()

	concurrent := 50
	files := 50

	jobs := make(chan int, files)
	results := make(chan string)

	for workerId := 0; workerId <= concurrent; workerId++ {
		go worker(workerId+1, jobs, results)
	}

	for i :=0; i < files; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < files; i++ {
		fmt.Println(<-results)
	}

	close(results)

	endProcess := time.Now()
	stringResult := fmt.Sprintf("Total time in seconds using %d goroutines: %f", concurrent, endProcess.Sub(startProcess).Seconds())
	fmt.Println(stringResult)
}