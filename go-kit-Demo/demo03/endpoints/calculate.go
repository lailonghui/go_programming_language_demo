/*
@Time : 2021/3/5 15:26
@Author : lai
@Description :
@File : calculate
*/
package endpoints

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"lai.com/go_programming_language_demo/go-kit-Demo/demo03/services"
)

// ArithmeticRequest define request struct
type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

// ArithmeticResponse define response struct
type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

// MakeArithmeticEndpoint make endpoints
func MakeArithmeticEndpoint(svc services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)

		var (
			res, a, b int
		)
		a = req.A
		b = req.B
		switch req.RequestType {
		case "Add":
			res = svc.Add(a, b)
		case "Subtract":
			res = svc.Subtract(a, b)
		case "Multiply":
			res = svc.Multiply(a, b)
		case "Divide":
			res, err = svc.Divide(a, b)
		default:
			return nil, errors.New("ErrInvalidRequestType")
		}

		return ArithmeticResponse{Result: res, Error: err}, nil
	}
}
