package main

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	raven "github.com/getsentry/raven-go"
	"github.com/oysterprotocol/lambda-node/hooknode/services"
)

type hooknodeReq struct {
	Provider string               `json:"provider"`
	Chunks   []services.IotaChunk `json:"chunks"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse request body.
	var req hooknodeReq
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		raven.CaptureError(err, nil)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	// TODO: Validate params.

	// PoW + Broadcast
	if err := services.AttachAndBroadcast(req.Provider, &req.Chunks); err != nil {
		raven.CaptureError(err, nil)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "Success!", StatusCode: 200}, nil
}

func main() {
	raven.SetDSN(os.Getenv("SENTRY_DSN"))

	raven.CapturePanic(func() {
		lambda.Start(handler)
	}, nil)
}
