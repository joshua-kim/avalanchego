// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

// Data common across all blocks
type Data struct {
	ID     ids.ID   `serialize:"true"`
	Parent ids.ID   `serialize:"true"`
	Bytes  []byte   `serialize:"true"`
	Height uint64   `serialize:"true"`
	Txs    []txs.Tx `serialize:"true"`
}

// Block in the P-Chain
type Block struct {
	Interface
	Data `serialize:"true"`
}

// Interface implements block-specific behavior
type Interface interface {
	snow.ContextInitializable
	// Visit calls [visitor] with this block's concrete type
	Visit(visitor Visitor) error

	// note: initialize does not assume that block transactions
	// are initialized, and initializes them itself if they aren't.
	initialize(bytes []byte) error
}

type Banff struct {
	Block
	Timestamp time.Time
}

func initialize(blk Interface) error {
	// We serialize this block as a pointer so that it can be deserialized into
	// a Interface
	bytes, err := Codec.Marshal(Version, &blk)
	if err != nil {
		return fmt.Errorf("couldn't marshal block: %w", err)
	}
	return blk.initialize(bytes)
}
