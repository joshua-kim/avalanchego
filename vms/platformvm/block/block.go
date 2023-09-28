// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"fmt"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
)

var _ Interface = (*Block)(nil)
var _ Interface = (*Banff)(nil)

// Interface implements block-specific behavior
type Interface interface {
	snow.ContextInitializable
	// note: initialize does not assume that block transactions
	// are initialized, and initializes them itself if they aren't.
	initialize(bytes []byte) error
	Visit(v Visitor) error
}

// Data common across all blocks
type Data struct {
	ID     ids.ID `serialize:"true"`
	Parent ids.ID `serialize:"true" json:"parentID"`
	Height uint64 `serialize:"true"`
}

// Block in the P-Chain
type Block struct {
	Interface
	Data  `serialize:"true"`
	Bytes []byte `serialize:"true"`
}

func (b *Block) initialize([]byte) error {
	bytes, err := Codec.Marshal(Version, b)
	if err != nil {
		return fmt.Errorf("couldn't marshal block: %w", err)
	}

	b.ID = hashing.ComputeHash256Array(bytes)
	b.Bytes = bytes

	return b.Interface.initialize(bytes)
}

type Banff struct {
	Block
	Time time.Time `serialize:"true"`
}

func (b *Banff) initialize([]byte) error {
	bytes, err := Codec.Marshal(Version, b)
	if err != nil {
		return fmt.Errorf("couldn't marshal block: %w", err)
	}

	return b.Interface.initialize(bytes)
}
