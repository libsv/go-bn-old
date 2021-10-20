package models

import "github.com/libsv/go-bc"

type Info struct {
	Version                      uint32  `json:"version"`
	ProtocolVersion              uint32  `json:"protocolversion"`
	Wallet                       uint32  `json:"wallet"`
	Balance                      float64 `json:"balance"`
	Blocks                       uint64  `json:"blocks"`
	TimeOffset                   uint32  `json:"timeoffset"`
	Connections                  uint32  `json:"connections"`
	Proxy                        string  `json:"proxy"`
	Difficulty                   float64 `json:"difficulty"`
	Testnet                      bool    `json:"testnet"`
	Stn                          bool    `json:"stn"`
	KeypoolOldest                uint32  `json:"keypoololdest"`
	KeypoolSize                  uint32  `json:"keypoolsize"`
	PayTxFee                     float64 `json:"paytxfee"`
	RelayFee                     float64 `json:"relayfee"`
	Errors                       string  `json:"errors"`
	MaxBlockSize                 uint64  `json:"maxblocksize"`
	MaxMinedBlockSize            uint64  `json:"maxminedblocksize"`
	MaxStackMemoryUsageConsensus uint64  `json:"maxstackmemoryusageconsensus"`
}

type Block struct {
	*bc.BlockHeader
	Tx            []string
	Hash          string
	Confirmations uint32
	Size          uint32
	Height        uint32
	Version       uint32
	VersionHex    string
	NumTx         uint64
	Time          uint64
	MedianTime    uint64
}
