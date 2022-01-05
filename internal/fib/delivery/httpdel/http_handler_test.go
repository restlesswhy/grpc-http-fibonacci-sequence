package httpdel

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/mock"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/models"
	"github.com/stretchr/testify/require"
)

type FibResponse struct {
	Seq map[int32]string `json:"seq"`
}

func TestFibHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFibUC := mock.NewMockUseCase(ctrl)
	handler := NewFibHandler(mockFibUC)

	inputBody := `{"from": 3, "to": 11}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/seq", bytes.NewBufferString(inputBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	handleFunc := handler.Get()
	resultSeq := models.FibSeq{
		Seq: make(map[int32]string),
	}
	resultSeq.Seq[1] = "1"
	
	
	mockFibUC.EXPECT().GetSeq(gomock.Any(), gomock.Any(), gomock.Any()).Return(resultSeq, nil)
	
	errR := handleFunc(c)
	var fibResponse FibResponse

	respByte, err := ioutil.ReadAll(rec.Body)
	require.NoError(t, err)
	_ = json.Unmarshal(respByte, &fibResponse)

	require.NoError(t, errR)
	require.Equal(t, 200, rec.Code)
	require.Equal(t, resultSeq.Seq, fibResponse.Seq)
}
