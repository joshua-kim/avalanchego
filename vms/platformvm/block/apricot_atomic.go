// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"fmt"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var _ Interface = (*ApricotAtomic)(nil)

func NewApricotAtomic(parentID ids.ID, height uint64, tx *txs.Tx) (Block, error) {
	blk := Block{
		Interface: &ApricotAtomic{
			Tx: tx,
		},
		Data: Data{
			ID:     ids.ID{},
			Parent: parentID,
			Height: height,
		},
		Bytes: nil,
	}

	return blk, initialize(blk)
}

// ApricotAtomic being accepted results in the atomic transaction contained
// in the block to be accepted and committed to the chain.
type ApricotAtomic struct {
	Tx *txs.Tx `serialize:"true" json:"tx"`
}

func (b *ApricotAtomic) initialize(bytes []byte) error {
	if err := b.Tx.Initialize(txs.Codec); err != nil {
		return fmt.Errorf("failed to initialize tx: %w", err)
	}
	return nil
}

func (b *ApricotAtomic) InitCtx(ctx *snow.Context) {
	b.Tx.Unsigned.InitCtx(ctx)
}

func (b *ApricotAtomic) Visit(v Visitor) error {
	return v.ApricotAtomicBlock(b)
}
