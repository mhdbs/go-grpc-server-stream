package main

import (
	"fmt"
	"log"
	greet "mhdbs/go-grpc-server-stream/pb"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello" + firstName + "Number" + strconv.Itoa(i)
		res := &greet.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)

	}
	return nil
}

func main() {
	fmt.Println("Server program")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Println("Failed to listen ", err)
	}
	s := grpc.NewServer()
	greet.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
