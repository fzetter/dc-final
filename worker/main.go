package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	//"github.com/anthonynsimon/bild/effect"

	pb "github.com/fzetter/dc-final/proto"
	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/sub"
	"go.nanomsg.org/mangos/protocol/rep"
	"google.golang.org/grpc"

	// register transports
	_ "go.nanomsg.org/mangos/transport/all"
)

var (
	defaultRPCPort = 50051
	busy = 0
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

var (
	publisherAddress  = "tcp://localhost:40899"
	requestAdress  		= "tcp://localhost:50899"
	workerName        = ""
	tags              = ""
)

/*
   Die
*/
func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

/*
   Date
*/
func date() string {
	return time.Now().Format(time.ANSIC)
}


/*
   Say Hello
	 Implements helloworld.GreeterServer
*/
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("RPC: Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

/*
   Grayscale
*/
func (s *server) GrayscaleFilter(ctx context.Context, in *pb.JobRequest) (*pb.JobReply, error) {
  return &pb.JobReply{Message: "Grayscale Filter " + in.GetName()}, nil
}

/*
   Blur
*/
func (s *server) BlurFilter(ctx context.Context, in *pb.JobRequest) (*pb.JobReply, error) {

	log.Println("*************************")
	log.Println(in.GetName())
	log.Println(in.GetWorkloadId())
	log.Println(in.GetFilter())
	log.Println("*************************")

  return &pb.JobReply{Message: "Blur Filter " + in.GetName()}, nil
}

/*
   Init
*/
func init() {
	flag.StringVar(&publisherAddress, "controller", "tcp://localhost:40899", "Controller address")
	flag.StringVar(&workerName, "worker-name", "hard-worker", "Worker Name")
	flag.StringVar(&tags, "tags", "gpu,superCPU,largeMemory", "Comma-separated worker tags")
}

/*
   Message Passing Socket Configuration
*/
func MessagePassingConfig() (subSock mangos.Socket, repSock mangos.Socket) {
	// PubSub
	var err error
	if subSock, err = sub.NewSocket(); err != nil {die("Can't get new sub socket: %s", err.Error()) }
	log.Printf("Connecting to controller on: %s", publisherAddress)
	if err = subSock.Dial(publisherAddress); err != nil { die("Can't dial on sub socket: %s", err.Error()) }
	err = subSock.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil { die("Cannot subscribe: %s", err.Error()) }

	// ReqRep
	if repSock, err = rep.NewSocket(); err != nil { die("Can't get new rep socket: %s", err) }
	if err = repSock.Dial(requestAdress); err != nil { die("Can't dial on sub socket: %s", err.Error()) }

	return subSock, repSock
}

/*
   Subscribe
*/
func Subscribe(sock mangos.Socket) {
	var msg []byte
	var err error
	if msg, err = sock.Recv(); err != nil { die("Cannot recv: %s", err.Error()) }
	log.Printf("Message-Passing: Worker(%s): Received %s\n", workerName, string(msg))
}

/*
   Reply
*/
func Reply(sock mangos.Socket) {

	var msg []byte
	var err error

	msg, err = sock.Recv()
	if err != nil { die("Cannot receive on rep socket: %s", err.Error()) }

	if string(msg) == "WORKER-STATUS" {
		fmt.Printf("Worker: Sending Status\n")
		var msg = []byte(fmt.Sprintf("%d", busy))
		err = sock.Send(msg)
		if err != nil { die("Can't send reply: %s", err.Error()) }
	}

}

/*
   Join Cluster
*/
func joinCluster() {

	subSock, repSock := MessagePassingConfig()

	for {
		Subscribe(subSock)
		Reply(repSock)
	}

}

/*
   Obtain Available Port
*/
func getAvailablePort() int {
	port := defaultRPCPort
	for {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			port = port + 1
			continue
		}
		ln.Close()
		break
	}
	return port
}

/*
   Main
*/
func main() {
	flag.Parse()

	// Message Passing Config
	go joinCluster()

	// Setup Worker RPC Server
	rpcPort := getAvailablePort()
	log.Printf("Starting RPC Service on localhost:%v", rpcPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", rpcPort))
	if err != nil { log.Fatalf("Failed to listen: %v", err) }

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil { log.Fatalf("Failed to serve: %v", err) }

}
