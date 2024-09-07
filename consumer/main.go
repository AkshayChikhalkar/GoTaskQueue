package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/akshaychikhalkar/GoTaskQueue/tasks" // Import the compiled gRPC protobuf

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedTaskServiceServer
}

func (s *server) SendTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	log.Printf("Received task: Type: %d, Value: %d", req.TaskType, req.TaskValue)

	// Simulate processing (sleep for task_value milliseconds)
	time.Sleep(time.Duration(req.TaskValue) * time.Millisecond)

	// TODO: Update the task's state in the database to "done"

	return &pb.TaskResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
