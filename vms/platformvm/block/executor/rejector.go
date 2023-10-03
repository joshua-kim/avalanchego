// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/vms/platformvm/block"
)

var _ block.Visitor = (*rejector)(nil)

// rejector handles the logic for rejecting a block.
// All errors returned by this struct are fatal and should result in the chain
// being shutdown.
type rejector struct {
	*backend
	addTxsToMempool bool
}

func (r *rejector) BanffAbort(b *block.BanffAbort) error {
	return r.rejectBlock(b, "banff abort")
}

func (r *rejector) BanffCommitBlock(b block.Banff) error {
	return r.rejectBlock(b.Data, "banff commit")
}

func (r *rejector) BanffProposalBlock(b block.Banff) error {
	return r.rejectBlock(b.Data, "banff proposal")
}

func (r *rejector) BanffStandardBlock(b block.Banff) error {
	return r.rejectBlock(b.Data, "banff standard")
}

func (r *rejector) ApricotAbortBlock(b block.data) error {
	return r.rejectBlock(b, "apricot abort")
}

func (r *rejector) ApricotCommitBlock(b block.data) error {
	return r.rejectBlock(b, "apricot commit")
}

func (r *rejector) ApricotProposalBlock(b *block.ApricotProposal) error {
	return r.rejectBlock(b, "apricot proposal")
}

func (r *rejector) ApricotStandardBlock(b *block.ApricotStandard) error {
	return r.rejectBlock(b, "apricot standard")
}

func (r *rejector) ApricotAtomicBlock(b *block.ApricotAtomic) error {
	return r.rejectBlock(b, "apricot atomic")
}

func (r *rejector) rejectBlock(b block.Txs, blockType string) error {
	blkID := b.ID()
	defer r.free(blkID)

	r.ctx.Log.Verbo(
		"rejecting block",
		zap.String("blockType", blockType),
		zap.Stringer("blkID", blkID),
		zap.Uint64("height", b.Height()),
		zap.Stringer("parentID", b.Parent()),
	)

	if !r.addTxsToMempool {
		return nil
	}

	for _, tx := range b.Txs() {
		if err := r.Mempool.Add(tx); err != nil {
			r.ctx.Log.Debug(
				"failed to reissue tx",
				zap.Stringer("txID", tx.ID()),
				zap.Stringer("blkID", blkID),
				zap.Error(err),
			)
		}
	}

	return nil
}
