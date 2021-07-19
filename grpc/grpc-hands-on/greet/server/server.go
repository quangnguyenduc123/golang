package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"pb"
)

type server struct{}

func (s server) SquareRoot(ctx context.Context, request *pb.SquareRootRequest) (*pb.SquareRootResponse, error) {
	fmt.Println("Start Server")
	number := request.GetNumber()
	if number<0{
		return nil, status.Error(codes.InvalidArgument, "Positive number is required")
	}
	return &pb.SquareRootResponse{Number: math.Sqrt(float64(number))}, nil
}

func (s server) ComputeAverage(server pb.CalculateService_ComputeAverageServer) error {
	fmt.Println("Start receiving message from client")
	sum := int64(0)
	numbers := int64(0)
	for  {
		req, err := server.Recv()
		if err == io.EOF{
			result := float64(sum/numbers)
			return server.SendAndClose(&pb.ComputeAverageResponse{Number: result})
		}
		if err != nil && err != io.EOF{
			log.Fatalf("Reading stream has error %v", err)
		}
		sum += req.GetNumber()
		numbers++

	}
}

func (s server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &pb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (s server) GreetManyTimes(request *pb.GreetManyTimesRequest, stream pb.GreetService_GreetManyTimesServer) error {
	fmt.Println("GreetManyTimes function was invoked")
	firstName := request.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello" + firstName + "number :" + strconv.Itoa(i)
		res := &pb.GreetManytimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000*time.Millisecond)
	}
	return nil
}

func (s server) LongGreet(greetServer pb.GreetService_LongGreetServer) error {
	fmt.Println("longGreet function was invoked")
	result := "Hello "
	for  {
		req, err := greetServer.Recv()
		if err == io.EOF {
			// reached the end of stream
			result = strings.TrimSuffix(result, ",")
			return greetServer.SendAndClose(&pb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil && err != io.EOF{
			log.Fatalf("Reading stream has error %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += firstName + ","
	}
}

func (s server) GreetEveryone(everyoneServer pb.GreetService_GreetEveryoneServer) error {
	panic("implement me")
}

func (s server) GreetWithDeadline(ctx context.Context, request *pb.GreetWithDeadlineRequest) (*pb.GreetWithDeadlineResponse, error) {
	panic("implement me")
}

func main()  {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &server{})

	pb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}