package svc_test

import (
	"context"
	"github.com/mbellgb/avg/pkg/svc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	service svc.Service
)

func init() {
	service = svc.NewBasicService()
}

func Test_MeanNoItemsReturns0(t *testing.T) {
	inputs := []int32{}
	result, _ := service.Mean(context.Background(), inputs)
	assert.Equal(t, float32(0.0), result)
}

func Test_CalculatesMean(t *testing.T) {
	inputs := []int32{1, 2, 3, 4}
	result, _ := service.Mean(context.Background(), inputs)
	assert.Equal(t, float32(2.5), result)
}

func Test_CalculatesMeanInclNegatives(t *testing.T) {
	inputs := []int32{-1, -2, 3, 4}
	result, _ := service.Mean(context.Background(), inputs)
	assert.Equal(t, float32(1), result)
}
