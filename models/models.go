package models

import (
	"errors"
	"fmt"

	validator "github.com/theflyingcodr/govalidator"
)

type blockVerbosity string

const (
	VerbosityRawBlock                blockVerbosity = "RAW_BLOCK"
	VerbosityDecodeHeader            blockVerbosity = "DECODE_HEADER"
	VerbosityDecodeTransactions      blockVerbosity = "DECODE_TRANSACTIONS"
	VerbosityDecodeHeaderAndCoinbase blockVerbosity = "DECODE_HEADER_AND_COINBASE"
)

type Request struct {
	ID      string        `json:"id"`
	JSONRpc string        `json:"jsonRpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
}

type Response struct {
	Result interface{} `json:"result"`
	Error  *Error      `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type BlockOptions struct {
	Verbosity blockVerbosity
}

func (b *BlockOptions) Validate() error {
	return validator.New().Validate("verbosity", func() error {
		switch b.Verbosity {
		case "", VerbosityRawBlock, VerbosityDecodeHeader,
			VerbosityDecodeTransactions, VerbosityDecodeHeaderAndCoinbase:
			return nil
		}

		return errors.New("invalid value")
	}).Err()
}
