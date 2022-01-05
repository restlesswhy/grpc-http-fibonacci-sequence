package grpcdel

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/mock"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/models"
	fiboService "github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/proto"
	"github.com/stretchr/testify/require"
)

func TestFibMicroservice_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFibUC := mock.NewMockUseCase(ctrl)
	fibMicroservice := NewFibMicroservice(mockFibUC)

	reqValue := &fiboService.FiboRequest{
		From: 1,
		To: 10,
	}

	t.Run("Get", func(t *testing.T) {

		// seq := mode
		resultSeq := models.FibSeq{
			Seq: make(map[int32]string),
		}
		

		mockFibUC.EXPECT().GetSeq(gomock.Any(), gomock.Any(), gomock.Any()).Return(resultSeq, nil)
		
		response, err := fibMicroservice.Get(context.Background(), reqValue)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Equal(t, resultSeq.Seq, response.Result)
	})
}