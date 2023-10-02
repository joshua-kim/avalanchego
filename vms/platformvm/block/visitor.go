// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

type Visitor interface {
	BanffAbort(*BanffAbort) error
	BanffCommitBlock(*BanffCommit) error
	BanffProposalBlock(*BanffProposal) error
	BanffStandardBlock(*BanffStandard) error

	ApricotAbortBlock(*ApricotAbort) error
	ApricotCommitBlock(*ApricotCommit) error
	ApricotProposalBlock(*ApricotProposal) error
	ApricotStandardBlock(*ApricotStandard) error
	ApricotAtomicBlock(*ApricotAtomic) error
}
