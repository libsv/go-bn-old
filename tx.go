package bn

import (
	"context"

	"github.com/libsv/go-bn/models"
	"github.com/libsv/go-bt/v2"
)

type TransactionClient interface {
	CreateRawTransaction(ctx context.Context, utxos bt.UTXOs, params models.ParamsCreateRawTransaction) (*bt.Tx, error)
	FundRawTransaction(ctx context.Context, tx *bt.Tx, opts *models.OptsFundRawTransaction) (*bt.Tx, error)
	RawTransaction(ctx context.Context, txID string) (*bt.Tx, error)
	SignRawTransaction(ctx context.Context, tx *bt.Tx, opts *models.OptsSignRawTransaction) (*bt.Tx, error)
	SendRawTransaction(ctx context.Context, tx *bt.Tx, opts *models.OptsSendRawTransaction) (*bt.Tx, error)
}

func NewTransactionClient(oo ...optFunc) TransactionClient {
	return NewNodeClient(oo...)
}

func (c *client) CreateRawTransaction(ctx context.Context, utxos bt.UTXOs, params models.ParamsCreateRawTransaction) (*bt.Tx, error) {
	params.SetIsMainnet(c.isMainnet)
	var resp string
	if err := c.rpc.Do(ctx, "createrawtransaction", &resp, c.argsFor(&params, utxos.NodeJSON())...); err != nil {
		return nil, err
	}
	return bt.NewTxFromString(resp)
}

func (c *client) FundRawTransaction(ctx context.Context, tx *bt.Tx, opts *models.OptsFundRawTransaction) (*bt.Tx, error) {
	var resp models.FundTransaction
	return resp.Tx, c.rpc.Do(ctx, "fundrawtransaction", &resp, c.argsFor(opts, tx.String())...)
}

func (c *client) RawTransaction(ctx context.Context, txID string) (*bt.Tx, error) {
	var resp bt.Tx
	return &resp, c.rpc.Do(ctx, "getrawtransaction", &resp, txID, true)
}

func (c *client) SendRawTransaction(ctx context.Context, tx *bt.Tx, opts *models.OptsSendRawTransaction) (*bt.Tx, error) {
	var resp models.SendRawTransaction
	return resp.Tx, c.rpc.Do(ctx, "sendrawtransaction", &resp, c.argsFor(opts, tx.String())...)
}
