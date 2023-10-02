// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var _ Interface = (*ApricotAtomic)(nil)

func NewApricotAtomic(parentID ids.ID, height uint64, tx *txs.Tx) (*ApricotAtomic, error) {
	blk := &ApricotAtomic{
		Tx: tx,
	}

	if err := blk.Tx.Initialize(txs.Codec); err != nil {
		return nil, err
	}

	data, err := newData(blk, parentID, height)
	blk.data = data

	return blk, err
}

// ApricotAtomic being accepted results in the atomic transaction contained
// in the block to be accepted and committed to the chain.
type ApricotAtomic struct {
	data
	Tx *txs.Tx `serialize:"true" json:"tx"`
}

func (b *ApricotAtomic) InitCtx(ctx *snow.Context) {
	b.Tx.Unsigned.InitCtx(ctx)
}

func (b *ApricotAtomic) Visit(v Visitor) error {
	return v.ApricotAtomicBlock(b)
}
