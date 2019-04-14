package server

import (
	"context"
	"fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	avgendpoint "github.com/mbellgb/avg/pkg/endpoint"
	pb "github.com/mbellgb/avg/pkg/pb"
	avgservice "github.com/mbellgb/avg/pkg/svc"
	"google.golang.org/grpc"
	"net"
)

type grpcServer struct {
	mean grpctransport.Handler
}

// NewGRPCServer creates a server
func NewGRPCServer(endpoints avgendpoint.Set) pb.AvgServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{
		mean: grpctransport.NewServer(
			endpoints.MeanEndpoint,
			encodeGRPCIntArrayRequest,
			encodeGRPCFloatResponse,
			options...,
		),
	}
}

func (s *grpcServer) Mean(ctx context.Context, req *pb.IntArrayRequest) (*pb.FloatResponse, error) {
	_, rep, err := s.mean.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FloatResponse), nil
}

// NewGRPCClient creates a gRPC client connection, represented as a service
func NewGRPCClient(conn *grpc.ClientConn) avgservice.Service {
	options := []grpctransport.ClientOption{}

	var meanEndpoint endpoint.Endpoint
	{
		meanEndpoint = grpctransport.NewClient(
			conn,
			"pb.Avg",
			"Mean",
			encodeGRPCIntArrayRequest,
			encodeGRPCFloatResponse,
			pb.FloatResponse{},
			options...,
		).Endpoint()
	}

	return &avgendpoint.Set{
		MeanEndpoint: meanEndpoint,
	}
}

func encodeGRPCIntArrayRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.IntArrayRequest)
	return req, nil
}

func encodeGRPCFloatResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.FloatResponse)
	return res, nil
}

// Start starts the server
func Start(host, port string) {
	addr := fmt.Sprintf("%s:%s", host, port)
	var (
		service   = avgservice.New()
		endpoints = avgendpoint.New(service)
		server    = NewGRPCServer(endpoints)
	)
	grpcListener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error starting listener: ", err)
	}
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterAvgServer(baseServer, server)
	baseServer.Serve(grpcListener)
	grpcListener.Close()
}
