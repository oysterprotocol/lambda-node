package services

import "github.com/oysterprotocol/lambda-node/hooknode/types"

type S3 struct{}

// ChunkStore interface
func (i *S3) AdaptReqChunks(chunks []types.ReqChunk) []types.IotaChunk {
	result := make([]types.IotaChunk, len(chunks))
	// TODO: Fetch Message from s3
	for i, reqChk := range chunks {
		result[i] = types.IotaChunk{
			Address: reqChk.Address,
			Value:   reqChk.Value,
			Message: reqChk.Message, // This is null, need to fetch!
			Tag:     reqChk.Tag,
		}
	}

	return result
}
