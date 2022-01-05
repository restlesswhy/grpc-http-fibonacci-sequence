package httpdel

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/mock"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/models"
	"github.com/stretchr/testify/require"
)

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
	exp := `{"1":"1"}
`

	mockFibUC.EXPECT().GetSeq(gomock.Any(), gomock.Any(), gomock.Any()).Return(resultSeq, nil)

	err := handleFunc(c)
	fmt.Println(rec.Body.String())
	require.NoError(t, err)
	require.Equal(t, 200, rec.Code)
	require.Equal(t, exp, rec.Body.String())
}