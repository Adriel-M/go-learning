package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Worker struct {
	id int
	jobsAccepted int
}

func (self *Worker) process(c chan int) {
	// forever
	for {
		data := <- c
		self.jobsAccepted++
		fmt.Printf("Worker %d got %d\n", self.id, data)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	workersToUse := 16
	c := make(chan int, 10)
	workers := make([]*Worker, workersToUse)
	for i := 0; i < workersToUse; i++ {
		worker := &Worker{id: i, jobsAccepted: 0}
		workers[i] = worker
		go worker.process(c)
	}

	for i := 0; i < 100; i++ {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 100)
	}
	for _, worker := range(workers) {
		fmt.Printf("Worker %d did %d jobs\n", worker.id, worker.jobsAccepted)
	}
}