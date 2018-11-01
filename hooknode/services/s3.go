package services

import (
	"sync"

	"github.com/oysterprotocol/lambda-node/hooknode/types"
)

type S3 struct{}

// ChunkStore interface
func (i *S3) AdaptReqChunks(chunks []types.ReqChunk) []types.IotaChunk {
	var sem sync.WaitGroup
	sem.Add(len(chunks)) // Semaphore the size of the number of chunks.

	result := make([]types.IotaChunk, len(chunks))
	// TODO: Fetch Message from s3
	for i, reqChk := range chunks {
		go func(i int, reqChk types.ReqChunk) {
			result[i] = types.IotaChunk{
				Address: reqChk.Address,
				Value:   reqChk.Value,
				Message: fetchMsg(reqChk.StoreUrl), // This is null, need to fetch!
				Tag:     reqChk.Tag,
			}

			sem.Done()
		}(i, reqChk)
	}

	sem.Wait()

	return result
}

func fetchMsg(url string) string {
	return url // TODO: Actually fetch from s3
}
