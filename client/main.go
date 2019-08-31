package main

import (
	"context"
	"fmt"
	"io"
	"log"
	greet "mhdbs/go-grpc-server-stream/pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client program")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
		defer cc.Close()
	}
	c := greet.NewGreetServiceClient(cc)
	doServerStreaming(c)

}

func doServerStreaming(c greet.GreetServiceClient) {
	fmt.Println("server streaming roc")
	req := &greet.GreetManyTimesRequest{
		Greeting: &greet.Greeting{
			FirstName: "bilal",
			LastName:  "mohammed",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greetmany times RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("End of the stream ")
			break
		}
		if err != nil {
			log.Fatalf("Error while  getting the stream %v", err)
		}
		fmt.Println("Response from greet many times %v", msg.GetResult())
	}

}
