// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ Interface = (*BanffAbort)(nil)
	_ Interface = (*ApricotAbort)(nil)
)

func NewBanffAbort(
	time time.Time,
	parentID ids.ID,
	height uint64,
) (Banff, error) {
	blk := Banff{
		Block: Block{
			Interface: &BanffAbort{},
			Data: Data{
				Parent: parentID,
				Height: height,
			},
		},
		Time: time,
	}

	return blk, initializeBanff(blk)
}

type BanffAbort struct{}

func (*BanffAbort) InitCtx(*snow.Context) {
	return
}

func (*BanffAbort) initialize([]byte) error {
	return nil
}

func (b *BanffAbort) Visit(v Visitor) error {
	return v.BanffAbort(b)
}

// NewApricotAbort is kept for testing purposes only.
// Following Banff activation and subsequent code cleanup, Apricot Abort blocks
// should be only verified (upon bootstrap), never created anymore
func NewApricotAbort(
	parentID ids.ID,
	height uint64,
) (Block, error) {
	blk := Block{
		Interface: &ApricotAbort{},
		Data: Data{
			Parent: parentID,
			Height: height,
		},
	}

	return blk, initialize(blk)
}

type ApricotAbort struct{}

func (b *ApricotAbort) initialize(bytes []byte) error {
	return nil
}

func (*ApricotAbort) InitCtx(*snow.Context) {}

func (*ApricotAbort) Txs() []*txs.Tx {
	return nil
}

func (b *ApricotAbort) Visit(v Visitor) error {
	return v.ApricotAbortBlock(b)
}
