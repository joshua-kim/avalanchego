// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import "github.com/ava-labs/avalanchego/codec"

func Parse(c codec.Manager, b []byte) (Interface, error) {
	var blk Interface
	if _, err := c.Unmarshal(b, &blk); err != nil {
		return nil, err
	}
	return blk, blk.initialize(b)
}
