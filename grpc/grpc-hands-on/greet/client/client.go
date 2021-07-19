package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"pb"
	"time"
)

func main(){
	fmt.Println("Hello world")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err!=nil{
		log.Fatal("Could not connect to server")
	}
	defer conn.Close()

	//c := pb.NewGreetServiceClient(conn)

	cal := pb.NewCalculateServiceClient(conn)

	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	//doCalculateService(cal)
	handleError(cal)
}

func doUnary(c pb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &pb.GreetRequest{
		Greeting: &pb.Greeting{
			FirstName: "Quang",
			LastName:  "Nguyen Duc",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming (c pb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &pb.GreetManyTimesRequest{
		Greeting: &pb.Greeting{
			FirstName: "Quang",
			LastName: "Nguyen Duc",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimesRPC %v", err)
	}
	for  {
		msg, err := resStream.Recv()
		if err == io.EOF{
			// we have reached the end of stream
			break;
		}
		if err != io.EOF && err != nil  {
			log.Fatalf("error while reading GreetManyTimesRPC %v", err)
		}
		fmt.Printf("Response from Server Streaming: %v", msg.GetResult())
	}
}

func doClientStreaming ( c pb.GreetServiceClient){
	fmt.Println("Starting to do a Client Streaming RPC...")
	requests := []*pb.LongGreetRequest{
		&pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: "Quang1",
				LastName: "Nguyen Duc",
			},
		},
		&pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: "Quang2",
				LastName: "Nguyen Duc",
			},
		},
		&pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: "Quang3",
				LastName: "Nguyen Duc",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err !=nil {
		log.Fatalf("error while call LongGreetRPC %v", err)
	}
	for _, request := range requests {
		fmt.Println("Sending")
		stream.Send(request)
		time.Sleep(1000*time.Millisecond)
	}
	res, err :=stream.CloseAndRecv()
	if err !=nil {
		log.Fatalf("error while receiving LongGreetRPC %v", err)
	}

	fmt.Println(res)
}

func doCalculateService(cal pb.CalculateServiceClient){
	fmt.Println("Starting to do a Client Calculate RPC...")
	numbers := []int64{3, 5, 9, 54, 23}
	stream, err := cal.ComputeAverage(context.Background())
	if err !=nil {
		log.Fatalf("error while call LongGreetRPC %v", err)
	}
	for _, request := range numbers {
		fmt.Println("Sending")
		stream.Send(&pb.ComputeAverageRequest{Number: request})
		time.Sleep(1000*time.Millisecond)
	}
	res, err :=stream.CloseAndRecv()
	if err !=nil {
		log.Fatalf("error while receiving LongGreetRPC %v", err)
	}

	fmt.Println(res)
}

func handleError(cal pb.CalculateServiceClient){
	fmt.Println("Starting to do a Client SquareRoot RPC...")
	res, err := cal.SquareRoot(context.Background(), &pb.SquareRootRequest{Number: -3})

	if err != nil {
		// check error is from grpc handle or not
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
		} else{
			log.Fatalf("Error from calling Server %v", err)
		}
	}else{
		fmt.Println(res.GetNumber())
	}

}