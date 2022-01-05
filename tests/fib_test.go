package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/httpdel"
)

type FibResponse struct {
	Seq map[int32]string `json:"seq"`
}

func (s *FibTestSuite) TestHttpGetSequence() {
	router := echo.New()
	httpdel.MapRoutes(router.Group("/api"), s.httpHandler)
	r := s.Require()

	inputBody := `{"from": 1, "to": 100}`

	req, _ := http.NewRequest("GET", "/api/seq", bytes.NewBuffer([]byte(inputBody)))
	req.Header.Set("Content-type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	
	r.Equal(http.StatusOK, resp.Result().StatusCode)

	var fibResponse FibResponse

	respByte, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	_ = json.Unmarshal(respByte, &fibResponse)

	resExpect, err := s.usecase.GetSeq(context.Background(), 1, 100)
	s.NoError(err) 

	r.Equal(resExpect.Seq, fibResponse.Seq)
}

func (s *FibTestSuite) TestGrpcGetSequence() {

}

func (s *FibTestSuite) TestHttpGetSequence_TryCache() {
	router := echo.New()
	httpdel.MapRoutes(router.Group("/api"), s.httpHandler)
	r := s.Require()

	inputBody := `{"from": 1, "to": 100}`

	req, _ := http.NewRequest("GET", "/api/seq", bytes.NewBuffer([]byte(inputBody)))
	req.Header.Set("Content-type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	r.Equal(http.StatusOK, resp.Result().StatusCode)

	exist := s.cache.Exists(context.Background(), "5")
	r.True(exist)
	
	time.Sleep(6 * time.Second)

	exist = s.cache.Exists(context.Background(), "5")
	r.False(exist)
}

func (s *FibTestSuite) TestHttpGetSequence_TryWithoutCache() {
	s.cfg.Redis.Caching = false
	router := echo.New()
	httpdel.MapRoutes(router.Group("/api"), s.httpHandler)
	r := s.Require()

	inputBody := `{"from": 1, "to": 100}`

	req, _ := http.NewRequest("GET", "/api/seq", bytes.NewBuffer([]byte(inputBody)))
	req.Header.Set("Content-type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	r.Equal(http.StatusOK, resp.Result().StatusCode)

	exist := s.cache.Exists(context.Background(), "5")
	r.False(exist)
}