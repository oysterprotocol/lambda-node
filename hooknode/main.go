package main

import (
	"log"
	"os"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	raven "github.com/getsentry/raven-go"
	"github.com/oysterprotocol/lambda-node/hooknode/services"
	"github.com/oysterprotocol/lambda-node/hooknode/types"
)

func handler(req types.HooknodeReq) (events.APIGatewayProxyResponse, error) {
	// TODO: Validate params.
	log.Printf("Handler request: %s with %d of chunks\n", req.Provider, len(req.Chunks))

	var chkStore types.ChunkStore
	switch req.StoreType {
	case "s3":
		chkStore = &services.S3{}
	default:
		chkStore = &services.Iota{}
	}

	chunks := chkStore.AdaptReqChunks(req.Chunks)

	// PoW + Broadcast
	if err := services.AttachAndBroadcast(req.Provider, &chunks); err != nil {
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
