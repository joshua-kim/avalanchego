// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"fmt"

	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
)

var _ block.Visitor = (*verifier)(nil)

// options supports build new option blocks
type options struct {
	// outputs populated by this struct's methods:
	commitBlock block.Interface
	abortBlock  block.Interface
}

func (*options) BanffAbort(*block.BanffAbort) error {
	return snowman.ErrNotOracle
}

func (*options) BanffCommitBlock(*block.BanffCommit) error {
	return snowman.ErrNotOracle
}

func (o *options) BanffProposalBlock(b *block.BanffProposal) error {
	timestamp := b.Timestamp()
	blkID := b.ID()
	nextHeight := b.Height() + 1

	var err error
	o.commitBlock, err = block.NewBanffCommit(timestamp, blkID, nextHeight)
	if err != nil {
		return fmt.Errorf(
			"failed to create commit block: %w",
			err,
		)
	}

	o.abortBlock, err = block.NewBanffAbort(timestamp, blkID, nextHeight)
	if err != nil {
		return fmt.Errorf(
			"failed to create abort block: %w",
			err,
		)
	}
	return nil
}

func (*options) BanffStandardBlock(*block.BanffStandard) error {
	return snowman.ErrNotOracle
}

func (*options) ApricotAbortBlock(*block.ApricotAbort) error {
	return snowman.ErrNotOracle
}

func (*options) ApricotCommitBlock(*block.ApricotCommit) error {
	return snowman.ErrNotOracle
}

func (o *options) ApricotProposalBlock(b *block.ApricotProposal) error {
	blkID := b.ID()
	nextHeight := b.Height() + 1

	var err error
	o.commitBlock, err = block.NewApricotCommitBlock(blkID, nextHeight)
	if err != nil {
		return fmt.Errorf(
			"failed to create commit block: %w",
			err,
		)
	}

	o.abortBlock, err = block.NewApricotAbort(blkID, nextHeight)
	if err != nil {
		return fmt.Errorf(
			"failed to create abort block: %w",
			err,
		)
	}
	return nil
}

func (*options) ApricotStandardBlock(*block.ApricotStandard) error {
	return snowman.ErrNotOracle
}

func (*options) ApricotAtomicBlock(*block.ApricotAtomic) error {
	return snowman.ErrNotOracle
}
