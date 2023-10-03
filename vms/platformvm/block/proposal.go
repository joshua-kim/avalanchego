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
	_ Txs = (*BanffProposal)(nil)
	_ Txs = (*ApricotProposal)(nil)
)

type proposalFields struct {
	Tx *txs.Tx `serialize:"true" json:"tx"`
}

type proposal struct {
	proposalFields
}

func (p proposal) Txs() []*txs.Tx {
	return []*txs.Tx{p.proposalFields.Tx}
}

func (p proposal) InitCtx(ctx *snow.Context) {
	p.proposalFields.Tx.Unsigned.InitCtx(ctx)
}

type BanffProposal struct {
	banffData
	proposal
	// Txs is currently unused. This is populated so that introducing
	// them in the future will not require a codec change.
	//
	// TODO: when Txs is used, we must correctly verify and apply their
	//       changes.
	Transactions []*txs.Tx `serialize:"true" json:"-"`
}

func (b *BanffProposal) InitCtx(ctx *snow.Context) {
	for _, tx := range b.Transactions {
		tx.Unsigned.InitCtx(ctx)
	}

	b.proposal.InitCtx(ctx)
}

func (b *BanffProposal) Visit(v Visitor) error {
	return v.BanffProposalBlock(b)
}

func NewBanffProposalBlock(time time.Time, parentID ids.ID, height uint64, tx *txs.Tx) (*BanffProposal, error) {
	blk := &BanffProposal{
		banffData: banffData{
			banffDataFields: banffDataFields{
				Time: time,
			},
		},
		proposal: proposal{
			proposalFields: proposalFields{
				Tx: tx,
			},
		},
	}

	data, err := newData(blk, parentID, height)
	blk.data = data

	return blk, err
}

func NewApricotProposal(
	parentID ids.ID,
	height uint64,
	tx *txs.Tx,
) (*ApricotProposal, error) {
	blk := &ApricotProposal{
		data: data{
			dataFields: dataFields{
				Parent: parentID,
				Height: height,
			},
		},
		proposal: proposal{
			proposalFields: proposalFields{
				Tx: tx,
			},
		},
	}

	if err := blk.Tx.Initialize(txs.Codec); err != nil {
		return nil, err
	}

	data, err := newData(blk, parentID, height)
	blk.data = data
	return blk, err
}

type ApricotProposal struct {
	data
	proposal
}

func (b *ApricotProposal) Visit(v Visitor) error {
	return v.ApricotProposalBlock(b)
}
