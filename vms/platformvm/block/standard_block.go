// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	pvmTxs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ Banff = (*BanffStandard)(nil)
	_ Txs   = (*BanffStandard)(nil)

	_ Txs = (*ApricotStandard)(nil)
)

func NewBanffStandard(
	time time.Time,
	parentID ids.ID,
	height uint64,
	txs []*pvmTxs.Tx,
) (*BanffStandard, error) {
	blk := &BanffStandard{
		banffData: banffData{
			banffDataFields: banffDataFields{
				Time: time,
			},
		},
		transactions: transactions{
			transactionsFields: transactionsFields{
				Txs: txs,
			},
		},
	}

	data, err := newData(blk, parentID, height)
	blk.data = data
	return blk, err
}

type BanffStandard struct {
	banffData
	transactions
}

func (*BanffStandard) InitCtx(*snow.Context) {}

func (b *BanffStandard) Visit(v Visitor) error {
	return v.BanffStandardBlock(b)
}

func NewApricotStandard(
	parentID ids.ID,
	height uint64,
	txs []*pvmTxs.Tx,
) (*ApricotStandard, error) {
	blk := &ApricotStandard{
		transactions: transactions{
			transactionsFields: transactionsFields{
				Txs: txs,
			},
		},
	}

	for _, tx := range blk.Txs() {
		if err := tx.Initialize(pvmTxs.Codec); err != nil {
			return nil, fmt.Errorf("failed to sign block: %w", err)
		}
	}

	data, err := newData(blk, parentID, height)
	blk.data = data
	return blk, err
}

type ApricotStandard struct {
	data
	transactions
}

func (b *ApricotStandard) InitCtx(ctx *snow.Context) {
	for _, tx := range b.Txs() {
		tx.Unsigned.InitCtx(ctx)
	}
}

func (b *ApricotStandard) Visit(v Visitor) error {
	return v.ApricotStandardBlock(b)
}

type transactionsFields struct {
	Txs []*pvmTxs.Tx `serialize:"true" json:"txs"`
}

type transactions struct {
	transactionsFields
}

func (t transactions) Txs() []*pvmTxs.Tx {
	return t.transactionsFields.Txs
}
