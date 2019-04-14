package svc

import (
	"context"
)

// Service defines the endpoints of avg service
type Service interface {
	Mean(context.Context, []int32) (float32, error)
}

// New creates a new service
func New() Service {
	var svc Service
	{
		svc = NewBasicService()
	}
	return svc
}

type basicService struct{}

// NewBasicService creates a basic implementation
// of avg service
func NewBasicService() Service {
	return basicService{}
}

func (svc basicService) Mean(_ context.Context, nums []int32) (float32, error) {
	// TODO: implement
	return 0.0, nil
}
