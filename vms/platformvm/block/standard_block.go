// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ Interface = (*BanffStandard)(nil)
	_ Interface = (*ApricotStandard)(nil)
)

type BanffStandard struct {
	Transactions []*txs.Tx `serialize:"true" json:"txs"`
}

func (BanffStandard) InitCtx(*snow.Context) {}

func (BanffStandard) initialize([]byte) error {
	return nil
}

func (b BanffStandard) Visit(v Visitor) error {
	return v.BanffStandardBlock(b)
}

func NewBanff(
	time time.Time,
	parentID ids.ID,
	height uint64,
	txs []*txs.Tx,
) (Banff, error) {
	blk := Banff{
		Block: Block{
			Interface: &BanffStandard{
				Transactions: txs,
			},
			Data: Data{
				Parent: parentID,
				Height: height,
			},
		},
		Time: time,
	}

	return blk, blk.initialize(blk.Bytes)
}

type ApricotStandard struct {
	Transactions []*txs.Tx `serialize:"true" json:"txs"`
}

func (b ApricotStandard) initialize([]byte) error {
	for _, tx := range b.Transactions {
		if err := tx.Initialize(txs.Codec); err != nil {
			return fmt.Errorf("failed to sign block: %w", err)
		}
	}
	return nil
}

func (b ApricotStandard) InitCtx(ctx *snow.Context) {
	for _, tx := range b.Transactions {
		tx.Unsigned.InitCtx(ctx)
	}
}

func (b ApricotStandard) Visit(v Visitor) error {
	return v.ApricotStandardBlock(b)
}

// NewApricotStandard is kept for testing purposes only.
// Following Banff activation and subsequent code cleanup, Apricot Standard blocks
// should be only verified (upon bootstrap), never created anymore
func NewApricotStandard(
	parentID ids.ID,
	height uint64,
	txs []*txs.Tx,
) (Block, error) {
	blk := Block{
		Interface: &ApricotStandard{
			Transactions: txs,
		},
		Data: Data{
			Parent: parentID,
			Height: height,
		},
	}

	return blk, blk.initialize(blk.Bytes)
}
