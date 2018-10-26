package services

import (
	"github.com/iotaledger/giota"
	"github.com/oysterprotocol/lambda-node/hooknode/types"
)

const seed = giota.Trytes("OYSTERPRLOYSTERPRLOYSTERPRLOYSTERPRLOYSTERPRLOYSTERPRLOYSTERPRLOYSTERPRLOYSTERPRL")
const security = 1
const value = int64(0)
const depth = int64(1)
const mwm = int64(6)

type Iota struct{}

// ChunkStore interface
func (i *Iota) AdaptReqChunks(chunks []types.ReqChunk) []types.IotaChunk {
	result := make([]types.IotaChunk, len(chunks))
	for i, reqChk := range chunks {
		result[i] = types.IotaChunk{
			Address: reqChk.Address,
			Value:   reqChk.Value,
			Message: reqChk.Message,
			Tag:     reqChk.Tag,
		}
	}

	return result
}

func AttachAndBroadcast(provider string, chunks *[]types.IotaChunk) error {
	// TODO: Add logging to segment so we know what's going on.

	api := giota.NewAPI(provider, nil)
	oysterTag, err := giota.ToTrytes("OYSTERHOOKNODE")
	if err != nil {
		return err
	}
	_, powFn := giota.GetBestPoW() // Log powName.

	// Map chunks to giota.Transfer
	trs := make([]giota.Transfer, len(*chunks))
	for i, chunk := range *chunks {
		addr, err := giota.ToAddress(chunk.Address)
		if err != nil {
			return err
		}
		msg, err := giota.ToTrytes(chunk.Message)
		if err != nil {
			return err
		}

		trs[i] = giota.Transfer{
			Address: addr,
			Message: msg,
			Value:   value,
			Tag:     oysterTag,
		}
	}

	bd, err := giota.PrepareTransfers(api, seed, trs, nil, "", security)
	if err != nil {
		return err
	}

	return giota.SendTrytes(api, depth, []giota.Transaction(bd), mwm, powFn)
}
