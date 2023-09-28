// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

var (
	_ Interface = (*BanffCommit)(nil)
	_ Interface = (*ApricotCommit)(nil)
)

type BanffCommit struct{}

func (BanffCommit) initialize([]byte) error {
	return nil
}

func (BanffCommit) InitCtx(*snow.Context) {
	return
}

func (b BanffCommit) Visit(v Visitor) error {
	return v.BanffCommitBlock(b)
}

func NewBanffCommit(
	timestamp time.Time,
	parentID ids.ID,
	height uint64,
) (Banff, error) {
	blk := Banff{
		Block: Block{
			Interface: &BanffCommit{},
			Data: Data{
				Parent: parentID,
				Height: height,
			},
		},
		Time: timestamp,
	}

	return blk, blk.initialize(blk.Bytes)
}

type ApricotCommit struct{}

func (ApricotCommit) initialize([]byte) error {
	return nil
}

func (ApricotCommit) InitCtx(*snow.Context) {}

func (b ApricotCommit) Visit(v Visitor) error {
	return v.ApricotCommitBlock(b)
}

func NewApricotCommitBlock(
	parentID ids.ID,
	height uint64,
) (Block, error) {
	blk := Block{
		Interface: &ApricotCommit{},
		Data: Data{
			Parent: parentID,
			Height: height,
		},
	}

	return blk, blk.initialize(blk.Bytes)
}
