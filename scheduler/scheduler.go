package scheduler

import (
	"context"
	"log"
	"time"

	pb "github.com/fzetter/dc-final/proto"
	"google.golang.org/grpc"
)

//const (
//	address     = "localhost:50051"
//	defaultName = "world"
//)

type Job struct {
	Address string
	RPCName string
	WorkloadId string
	Filter string
}

/*
   Schedule
*/
func schedule(job Job) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(job.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil { log.Fatalf("Did not connect: %v", err) }
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Blur
	if (job.Filter == "blur") {
		t, err := c.BlurFilter(ctx, &pb.JobRequest{Name: job.RPCName, WorkloadId: job.WorkloadId, Filter: job.Filter})
		if err != nil { log.Fatalf("Could not greet: %v", err) }
		log.Printf("Blur Res: %s", t.GetMessage())

	// Grayscale
	} else if (job.Filter == "grayscale") {

		s, err := c.GrayscaleFilter(ctx, &pb.JobRequest{Name: job.RPCName, WorkloadId: job.WorkloadId, Filter: job.Filter})
		if err != nil { log.Fatalf("Could not greet: %v", err) }
		log.Printf("Grayscale Res: %s", s.GetMessage())

	// Hello
	} else {
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: job.RPCName})
		if err != nil { log.Fatalf("Could not greet: %v", err) }
		log.Printf("Hello Res: %s", r.GetMessage())
	}


}

/*
   Start
*/
func Start(jobs chan Job) error {
	for {
		job := <-jobs
		schedule(job)
	}
	return nil
}
