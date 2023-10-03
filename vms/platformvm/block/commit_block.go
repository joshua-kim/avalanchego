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

func NewBanffCommit(timestamp time.Time, parentID ids.ID, height uint64) (*BanffCommit, error) {
	blk := &BanffCommit{
		banffData: banffData{
			banffDataFields: banffDataFields{
				Time: timestamp,
			},
		},
	}

	data, err := newData(blk, parentID, height)
	blk.data = data
	return blk, err
}

type BanffCommit struct {
	banffData
}

func (*BanffCommit) InitCtx(*snow.Context) {
	return
}

func (b *BanffCommit) Visit(v Visitor) error {
	return v.BanffCommitBlock(b)
}

func NewApricotCommitBlock(parentID ids.ID, height uint64) (*ApricotCommit, error) {
	blk := &ApricotCommit{}

	data, err := newData(blk, parentID, height)
	blk.data = data
	return blk, err
}

type ApricotCommit struct {
	data
}

func (*ApricotCommit) InitCtx(*snow.Context) {}

func (b *ApricotCommit) Visit(v Visitor) error {
	return v.ApricotCommitBlock(b)
}
