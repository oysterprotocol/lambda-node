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
	Message string `json:"message"`
	Tag     string `json:"tag"`
}

type ChunkStore interface {
	AdaptReqChunks(chunks []ReqChunk) []IotaChunk
}
