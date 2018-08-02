package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/oysterprotocol/lambda-node/hooknode/services"
)

type Response struct {
	Message string `json:"message"`
}

type hooknodeReq struct {
	Provider string               `json:"provider"`
	Chunks   []services.IotaChunk `json:"chunks"`
}

func handler(req hooknodeReq) (Response, error) {
	// TODO: Validate params.

	if err := services.AttachAndBroadcast(req.Provider, &req.Chunks); err != nil {
		return Response{Message: err.Error()}, err
	}

	return Response{Message: "Success!"}, nil
}

func main() {
	lambda.Start(handler)
}
