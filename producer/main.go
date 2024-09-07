package main

import (
    "context"
    "log"
    "math/rand"
    "time"

    pb "github.com/akshaychikhalkar/GoTaskQueue/tasks"  // Import the compiled gRPC protobuf

    "google.golang.org/grpc"
)

const (
    address = "localhost:50051"  // Address of the gRPC server (Consumer)
)

func main() {
    // Connect to gRPC server
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewTaskServiceClient(conn)

    // Generate random tasks and send to the consumer
    for {
        taskType := rand.Intn(10)
        taskValue := rand.Intn(100)

        req := &pb.TaskRequest{
            TaskType:  int32(taskType),
            TaskValue: int32(taskValue),
        }

        res, err := client.SendTask(context.Background(), req)
        if err != nil {
            log.Fatalf("Error sending task: %v", err)
        }
        log.Printf("Task sent: %v, Success: %v", req, res.Success)
        time.Sleep(2 * time.Second)
    }
}
