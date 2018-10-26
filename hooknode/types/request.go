package types

type ReqChunk struct {
	Address  string `json:"address"`
	Value    int    `json:"value"`
	Message  string `json:"message"`
	StoreUrl string `json:"storeUrl"`
	Tag      string `json:"tag"`
}

type HooknodeReq struct {
	Provider  string     `json:"provider"`
	StoreType string     `json:"storeType"`
	Chunks    []ReqChunk `json:"chunks"`
}

type IotaChunk struct {
	Address string `json:"address"`
	Value   int    `json:"value"`
	Tag     string `json:"tag"`

	// The following should be a discriminated union; only 1 will be non-null.
	// Maybe in Go 2.0 https://github.com/golang/go/issues/19412
	Message  string `json:"message"`
	StoreUrl string `json:"storeUrl"`
}

type ChunkStore interface {
	AdaptReqChunks(chunks []ReqChunk) []IotaChunk
}
