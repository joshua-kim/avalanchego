// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

var (
	_ Interface = (*BanffAbort)(nil)
	_ Interface = (*ApricotAbort)(nil)
)

func NewBanffAbort(time time.Time, parentID ids.ID, height uint64) (*BanffAbort, error) {
	blk := &BanffAbort{
		banffData: banffData{
			Time: time,
		},
	}

	data, err := newData(blk, parentID, height)
	blk.data = data

	return blk, err
}

type BanffAbort struct {
	banffData
}

func (*BanffAbort) InitCtx(*snow.Context) {
	return
}

func (b *BanffAbort) Visit(v Visitor) error {
	return v.BanffAbort(b)
}

// NewApricotAbort is kept for testing purposes only.
// Following Banff activation and subsequent code cleanup, Apricot Abort blocks
// should be only verified (upon bootstrap), never created anymore
func NewApricotAbort(
	parentID ids.ID,
	height uint64,
) (*ApricotAbort, error) {
	blk := &ApricotAbort{}

	data, err := newData(blk, parentID, height)
	blk.data = data
	return blk, err
}

type ApricotAbort struct {
	data
}

func (*ApricotAbort) InitCtx(*snow.Context) {}

func (b *ApricotAbort) Visit(v Visitor) error {
	return v.ApricotAbortBlock(b)
}
