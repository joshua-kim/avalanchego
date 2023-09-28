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

func NewApricotAtomic(
	parentID ids.ID,
	height uint64,
	tx txs.Tx,
) (*Block, error) {
	blk := &Block{
		Interface: &ApricotAtomic{
			Data: Data{
				ID:     ids.ID{},
				Parent: parentID,
				Bytes:  nil,
				Height: height,
				Txs:    []txs.Tx{tx},
			},
		},
	}
	return blk, initialize(blk)
}

// ApricotAtomic being accepted results in the atomic transaction contained
// in the block to be accepted and committed to the chain.
type ApricotAtomic struct {
	Data
}

func (b *ApricotAtomic) initialize(bytes []byte) error {
	b.Data.Bytes = bytes
	if err := b.Txs[0].Initialize(txs.Codec); err != nil {
		return fmt.Errorf("failed to initialize tx: %w", err)
	}
	return nil
}

func (b *ApricotAtomic) InitCtx(ctx *snow.Context) {
	b.Txs[0].Unsigned.InitCtx(ctx)
}

func (b *ApricotAtomic) Visit(v Visitor) error {
	return v.ApricotAtomicBlock(b)
}
