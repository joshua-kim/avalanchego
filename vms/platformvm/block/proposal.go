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
	_ Interface = (*BanffProposal)(nil)
	_ Interface = (*ApricotProposal)(nil)
)

type BanffProposal struct {
	Tx *txs.Tx `serialize:"true" json:"tx"`
	// Transactions is currently unused. This is populated so that introducing
	// them in the future will not require a codec change.
	//
	// TODO: when Transactions is used, we must correctly verify and apply their
	//       changes.
	Transactions []*txs.Tx `serialize:"true" json:"-"`
}

func (b BanffProposal) initialize(bytes []byte) error {
	return nil
}

func (b BanffProposal) InitCtx(ctx *snow.Context) {
	for _, tx := range b.Transactions {
		tx.Unsigned.InitCtx(ctx)
	}
	b.Tx.Unsigned.InitCtx(ctx)
}

func (b BanffProposal) Visit(v Visitor) error {
	return v.BanffProposalBlock(b)
}

func NewBanffProposalBlock(timestamp time.Time, parentID ids.ID, height uint64, tx *txs.Tx) (Banff, error) {
	blk := Banff{
		Block: Block{
			Interface: &BanffProposal{
				Tx: tx,
			},
			Data: Data{
				ID:     ids.ID{},
				Parent: parentID,
				Height: height,
			},
		},
		Time: timestamp,
	}

	return blk, blk.initialize(blk.Bytes)
}

type ApricotProposal struct {
	Tx *txs.Tx `serialize:"true" json:"tx"`
}

func (b *ApricotProposal) initialize([]byte) error {
	return b.Tx.Initialize(txs.Codec)
}

func (b *ApricotProposal) InitCtx(ctx *snow.Context) {
	b.Tx.Unsigned.InitCtx(ctx)
}

func (b *ApricotProposal) Visit(v Visitor) error {
	return v.ApricotProposalBlock(b)
}

// NewApricotProposal is kept for testing purposes only.
// Following Banff activation and subsequent code cleanup, Apricot Proposal blocks
// should be only verified (upon bootstrap), never created anymore
func NewApricotProposal(
	parentID ids.ID,
	height uint64,
	tx *txs.Tx,
) (Block, error) {
	blk := Block{
		Interface: &ApricotProposal{
			Tx: tx,
		},
		Data: Data{
			Parent: parentID,
			Height: height,
		},
	}

	return blk, blk.initialize(blk.Bytes)
}
