package main

import (
	"os"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	raven "github.com/getsentry/raven-go"
	"github.com/oysterprotocol/lambda-node/hooknode/services"
)

type hooknodeReq struct {
	Provider string               `json:"provider"`
	Chunks   []services.IotaChunk `json:"chunks"`
}

func handler(req hooknodeReq) (events.APIGatewayProxyResponse, error) {
	// TODO: Validate params.

	// PoW + Broadcast
	if err := services.AttachAndBroadcast(req.Provider, &req.Chunks); err != nil {
		raven.CaptureError(err, nil)

		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "Success!", StatusCode: 200}, nil
}

func main() {
	// Allow running on multiple cores. Kinda weird that this is manual?
	// https://golang.org/pkg/runtime/#GOMAXPROCS
	runtime.GOMAXPROCS(runtime.NumCPU())

	raven.SetDSN(os.Getenv("SENTRY_DSN"))

	raven.CapturePanic(func() {
		lambda.Start(handler)
	}, nil)
}
