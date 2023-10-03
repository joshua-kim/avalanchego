// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"fmt"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

// Interface implements block-specific behavior
type Interface interface {
	snow.ContextInitializable

	ID() ids.ID
	Parent() ids.ID
	Height() uint64
	Bytes() []byte

	Visit(v Visitor) error
}

type Txs interface {
	Interface
	Txs() []*txs.Tx
}

type Banff interface {
	Interface
	Time() time.Time
}

func newData(block Interface, parentID ids.ID, height uint64) (data, error) {
	bytes, err := Codec.Marshal(Version, block)
	if err != nil {
		return data{}, fmt.Errorf("couldn't marshal block: %w", err)
	}

	id := hashing.ComputeHash256Array(bytes)

	return data{
		dataFields: dataFields{
			ID:     id,
			Parent: parentID,
			Height: height,
			Bytes:  bytes,
		},
	}, nil
}

type data struct {
	dataFields
}

func (d data) ID() ids.ID {
	return d.dataFields.ID
}

func (d data) Parent() ids.ID {
	return d.dataFields.Parent
}

func (d data) Height() uint64 {
	return d.dataFields.Height
}

func (d data) Bytes() []byte {
	return d.dataFields.Bytes
}

// non-exported so function names on the interface don't collide with field names
type dataFields struct {
	ID     ids.ID `serialize:"true"`
	Parent ids.ID `serialize:"true" json:"parentID"`
	Height uint64 `serialize:"true"`
	Bytes  []byte `serialize:"true"`
}

type banffData struct {
	data
	banffDataFields
}

func (b banffData) Time() time.Time {
	return b.banffDataFields.Time
}

type banffDataFields struct {
	Time time.Time `serialize:"true"`
}
