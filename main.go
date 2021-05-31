package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"github.com/fzetter/dc-final/api"
	"github.com/fzetter/dc-final/controller"
	"github.com/fzetter/dc-final/scheduler"
)

func main() {
	log.Println("Welcome to the Distributed and Parallel Image Processing System")

	// Channels
	adminAccess := make(chan string)
	jobs := make(chan scheduler.Job)

	// Start Controller
	go controller.Start(adminAccess, jobs)

	// Start Scheduler
	go scheduler.Start(jobs)

	// Send sample jobs
	sampleJob := scheduler.Job{Address: "localhost:50051", RPCName: "hello"}

	// API
	go api.Start(adminAccess)

	for {
		sampleJob.RPCName = fmt.Sprintf("hello-%v", rand.Intn(10000))
		jobs <- sampleJob
		time.Sleep(time.Second * 5)
	}
}
