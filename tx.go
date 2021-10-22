package bn

import (
	"context"

	"github.com/libsv/go-bt/v2"
)

type TransactionClient interface {
	CreateRawTransaction(ctx context.Context, utxos []*bt.UTXO, outputs []*bt.Output) (interface{}, error)
	RawTransaction(ctx context.Context, txID string) (*bt.Tx, error)
}

func NewTransactionClient(oo ...optFunc) TransactionClient {
	return NewNodeClient(oo...)
}

func (c *client) RawTransaction(ctx context.Context, txID string) (*bt.Tx, error) {
	var resp bt.Tx
	return &resp, c.rpc.Do(ctx, "getrawtransaction", &resp, txID, true)
}

func (c *client) CreateRawTransaction(ctx context.Context, utxos []*bt.UTXO, outputs []*bt.Output) (interface{}, error) {
	return nil, nil
}
