package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/mbellgb/avg/pkg/pb"
	service "github.com/mbellgb/avg/pkg/svc"
)

// Set is a collection of all of the endpoints
type Set struct {
	MeanEndpoint endpoint.Endpoint
}

// New creates a new set
func New(svc service.Service) Set {
	var meanEndpoint endpoint.Endpoint
	{
		meanEndpoint = MakeMeanEndpoint(svc)
	}
	return Set{
		MeanEndpoint: meanEndpoint,
	}
}

// Mean calculates the mean
func (s *Set) Mean(ctx context.Context, values []int32) (float32, error) {
	resp, err := s.MeanEndpoint(ctx, &pb.IntArrayRequest{Data: values})
	if err != nil {
		return 0, err
	}
	response := resp.(*pb.FloatResponse)
	return response.Result, nil
}

// MakeMeanEndpoint creates endplint object
func MakeMeanEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.IntArrayRequest)
		result, err := svc.Mean(ctx, req.Data)
		return &pb.FloatResponse{Result: result, Error: errToStr(err)}, nil
	}
}

func strToErr(str string) error {
	if str == "" {
		return nil
	}
	return errors.New(str)
}

func errToStr(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
