// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ BlueberryBlock = &BlueberryStandardBlock{}
	_ Block          = &ApricotStandardBlock{}
)

type BlueberryStandardBlock struct {
	Time                 uint64 `serialize:"true" json:"time"`
	ApricotStandardBlock `serialize:"true"`
}

func (b *BlueberryStandardBlock) Timestamp() time.Time  { return time.Unix(int64(b.Time), 0) }
func (b *BlueberryStandardBlock) Visit(v Visitor) error { return v.BlueberryStandardBlock(b) }

func NewBlueberryStandardBlock(
	timestamp time.Time,
	parentID ids.ID,
	height uint64,
	txes []*txs.Tx,
) (*BlueberryStandardBlock, error) {
	blk := &BlueberryStandardBlock{
		Time: uint64(timestamp.Unix()),
		ApricotStandardBlock: ApricotStandardBlock{
			CommonBlock: CommonBlock{
				PrntID: parentID,
				Hght:   height,
			},
			Transactions: txes,
		},
	}
	return blk, initialize(blk)
}

type ApricotStandardBlock struct {
	CommonBlock  `serialize:"true"`
	Transactions []*txs.Tx `serialize:"true" json:"txs"`
}

func (b *ApricotStandardBlock) initialize(bytes []byte) error {
	b.CommonBlock.initialize(bytes)
	for _, tx := range b.Transactions {
		if err := tx.Sign(txs.Codec, nil); err != nil {
			return fmt.Errorf("failed to sign block: %w", err)
		}
	}
	return nil
}

func (b *ApricotStandardBlock) Txs() []*txs.Tx        { return b.Transactions }
func (b *ApricotStandardBlock) Visit(v Visitor) error { return v.ApricotStandardBlock(b) }

func NewApricotStandardBlock(
	parentID ids.ID,
	height uint64,
	txes []*txs.Tx,
) (*ApricotStandardBlock, error) {
	blk := &ApricotStandardBlock{
		CommonBlock: CommonBlock{
			PrntID: parentID,
			Hght:   height,
		},
		Transactions: txes,
	}
	return blk, initialize(blk)
}