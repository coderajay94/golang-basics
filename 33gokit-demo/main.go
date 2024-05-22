package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	transporthttp "github.com/go-kit/kit/transport/http"
)

var ErrorEmpty = errors.New("length is zero")

type StringService interface {
	UpperCase(st string) (string, error)
	Count(st string) int
}

type stringService struct{}

type upperCaseRequest struct {
	S string `json:"s"`
}

type upperCaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err, omitempty"`
}

type countRequest struct {
	S string `json:"s`
}

type countResponse struct {
	S     string `json:"s"`
	Count int    `json:"count"`
}

func (s stringService) UpperCase(st string) (string, error) {
	if len(st) == 0 {
		return st, ErrorEmpty
	} else {
		return strings.ToUpper(st), nil
	}
}

func (s stringService) Count(st string) int {
	return len(st)
}

func makeUpperCaseEndpoint(service stringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(upperCaseRequest)
		res, err := service.UpperCase(req.S)
		if err != nil {
			return upperCaseResponse{res, err.Error()}, err
		}
		return upperCaseResponse{res, ""}, nil
	}
}

func makeCountEndpoint(service stringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		count := service.Count(req.S)
		return countResponse{req.S, count}, nil
	}
}

func main() {

	service := stringService{}

	fmt.Println("starting go-kit tutorial...")

	upperCaseHandler := transporthttp.NewServer(makeUpperCaseEndpoint(service), decodeUpperCaseRequest, encodeResponse)
	countHandler := transporthttp.NewServer(makeCountEndpoint(service), decodeCountRequest, encodeResponse)

	http.Handle("/uppercase", upperCaseHandler)
	http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeCountRequest(context context.Context, r *http.Request) (interface{}, error) {
	var request countRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpperCaseRequest(context context.Context, r *http.Request) (interface{}, error) {
	var request upperCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(context context.Context, r http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(r).Encode(response)
}
