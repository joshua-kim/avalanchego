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
	_ Banff     = (*BanffAbortBlock)(nil)
	_ Interface = (*ApricotAbortBlock)(nil)
)

type BanffAbortBlock struct {
	Time              uint64 `serialize:"true" json:"time"`
	ApricotAbortBlock `serialize:"true"`
}

func (b *BanffAbortBlock) Timestamp() time.Time {
	return time.Unix(int64(b.Time), 0)
}

func (b *BanffAbortBlock) Visit(v Visitor) error {
	return v.BanffAbortBlock(b)
}

func NewBanffAbortBlock(
	timestamp time.Time,
	parentID ids.ID,
	height uint64,
) (*BanffAbortBlock, error) {
	blk := &BanffAbortBlock{
		Time: uint64(timestamp.Unix()),
		ApricotAbortBlock: ApricotAbortBlock{
			Block: CommonBlock{
				PrntID: parentID,
				Hght:   height,
			},
		},
	}
	return blk, initialize(blk)
}

type ApricotAbortBlock struct {
	Block `serialize:"true"`
}

func (b *ApricotAbortBlock) initialize(bytes []byte) error {
	b.Block.initialize(bytes)
	return nil
}

func (*ApricotAbortBlock) InitCtx(*snow.Context) {}

func (*ApricotAbortBlock) Txs() []*txs.Tx {
	return nil
}

func (b *ApricotAbortBlock) Visit(v Visitor) error {
	return v.ApricotAbortBlock(b)
}

// NewApricotAbortBlock is kept for testing purposes only.
// Following Banff activation and subsequent code cleanup, Apricot Abort blocks
// should be only verified (upon bootstrap), never created anymore
func NewApricotAbortBlock(
	parentID ids.ID,
	height uint64,
) (*ApricotAbortBlock, error) {
	blk := &ApricotAbortBlock{
		Block: CommonBlock{
			PrntID: parentID,
			Hght:   height,
		},
	}
	return blk, initialize(blk)
}
