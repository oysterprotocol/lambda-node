package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/oysterprotocol/lambda-node/hooknode/services"
)

type hooknodeReq struct {
	Provider string               `json:"provider"`
	Chunks   []services.IotaChunk `json:"chunks"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO: Validate params.

	// Parse request body.
	var req hooknodeReq
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 200}, nil
	}

	// PoW + Broadcast
	if err := services.AttachAndBroadcast(req.Provider, &req.Chunks); err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 200}, nil
	}

	return events.APIGatewayProxyResponse{Body: "Success!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
