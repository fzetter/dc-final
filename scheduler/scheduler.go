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

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: job.RPCName})
	if err != nil { log.Fatalf("Could not greet: %v", err) }
	log.Printf("Scheduler: RPC respose from %s : %s", job.Address, r.GetMessage())

	s, err := c.GrayscaleFilter(ctx, &pb.JobRequest{Name: job.RPCName, WorkloadId: job.WorkloadId, Filter: job.Filter})
	if err != nil { log.Fatalf("Could not greet: %v", err) }
	log.Printf("Greeting: %s", s.GetMessage())

	t, err := c.BlurFilter(ctx, &pb.JobRequest{Name: job.RPCName, WorkloadId: job.WorkloadId, Filter: job.Filter})
	if err != nil { log.Fatalf("Could not greet: %v", err) }
	log.Printf("Greeting: %s", t.GetMessage())
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
