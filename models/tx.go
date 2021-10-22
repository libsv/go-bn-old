package models

import (
	"encoding/json"

	"github.com/libsv/go-bt/v2"
)

type Output struct {
	BestBlock     string
	Confirmations uint32
	Coinbase      bool

	*bt.Output
}

func (o *Output) NodeJSON() interface{} {
	return o
}

func (o *Output) UnmarshalJSON(b []byte) error {
	oj := struct {
		BestBlock     string `json:"bestblock"`
		Confirmations uint32 `json:"confirmations"`
		Coinbase      bool   `json:"coinbase"`
	}{}

	if err := json.Unmarshal(b, &oj); err != nil {
		return err
	}

	var out bt.Output
	if err := json.Unmarshal(b, out.NodeJSON()); err != nil {
		return err
	}

	o.BestBlock = oj.BestBlock
	o.Confirmations = oj.Confirmations
	o.Coinbase = oj.Coinbase
	*o.Output = out

	return nil
}

type OutputSetInfo struct {
	Height         uint32  `json:"height"`
	BestBlock      string  `json:"bestblock"`
	Transactions   uint32  `json:"transactions"`
	OutputCount    uint32  `json:"txouts"`
	BogoSize       uint32  `json:"bogosize"`
	HashSerialised string  `json:"hash_serialized"`
	DiskSize       uint32  `json:"disk_size"`
	TotalAmount    float64 `json:"total_amount"`
}

type OptsOutput struct {
	IncludeMempool bool
}

func (o *OptsOutput) Args() []interface{} {
	return []interface{}{o.IncludeMempool}
}
