// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

type Visitor interface {
	BanffAbort(Data) error
	BanffCommitBlock(Data) error
	BanffProposalBlock(BanffProposal) error
	BanffStandardBlock(BanffStandard) error

	ApricotAbortBlock(Data) error
	ApricotCommitBlock(Data) error
	ApricotProposalBlock(Data) error
	ApricotStandardBlock(ApricotStandard) error
	ApricotAtomicBlock(ApricotAtomic) error
}
