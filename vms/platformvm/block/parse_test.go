// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var preFundedKeys = secp256k1.TestKeys()

func TestStandardBlocks(t *testing.T) {
	// check Apricot standard block can be built and parsed
	require := require.New(t)
	blkTimestamp := time.Now()
	parentID := ids.ID{'p', 'a', 'r', 'e', 'n', 't', 'I', 'D'}
	height := uint64(2022)
	txs, err := testDecisionTxs()
	require.NoError(err)

	for _, cdc := range []codec.Manager{Codec, GenesisCodec} {
		// build block
		apricotStandardBlk, err := NewApricotStandard(parentID, height, txs)
		require.NoError(err)

		// parse block
		parsed, err := Parse(cdc, apricotStandardBlk.Bytes)
		require.NoError(err)
		parsedApricotStandardBlk, ok := parsed.(*ApricotStandard)
		require.True(ok)

		// compare content
		require.Equal(apricotStandardBlk.ID, parsedApricotStandardBlk.ID)
		require.Equal(apricotStandardBlk.Bytes, parsedApricotStandardBlk.Bytes)
		require.Equal(apricotStandardBlk.ParentID, parsedApricotStandardBlk.ParentID)
		require.Equal(apricotStandardBlk.Height, parsedApricotStandardBlk.Height)
		require.Equal(txs, parsedApricotStandardBlk.Txs)

		// check that banff standard block can be built and parsed
		banffStandardBlk, err := NewBanffStandard(blkTimestamp, parentID, height, txs)
		require.NoError(err)

		// parse block
		parsed, err = Parse(cdc, banffStandardBlk.Bytes)
		require.NoError(err)
		parsedBanffStandardBlk, ok := parsed.(*BanffStandard)
		require.True(ok)

		// compare content
		require.Equal(banffStandardBlk.ID, parsedBanffStandardBlk.ID)
		require.Equal(banffStandardBlk.Bytes, parsedBanffStandardBlk.Bytes)
		require.Equal(banffStandardBlk.ParentID, parsedBanffStandardBlk.ParentID)
		require.Equal(banffStandardBlk.Height, parsedBanffStandardBlk.Height)
		require.Equal(banffStandardBlk.Txs, parsedBanffStandardBlk.Txs)
		require.Equal(banffStandardBlk.Time, parsedBanffStandardBlk.Time)
	}
}

func TestProposalBlocks(t *testing.T) {
	// check Apricot proposal block can be built and parsed
	require := require.New(t)
	blkTimestamp := time.Now()
	parentID := ids.ID{'p', 'a', 'r', 'e', 'n', 't', 'I', 'D'}
	height := uint64(2022)
	tx, err := testProposalTx()
	require.NoError(err)

	for _, cdc := range []codec.Manager{Codec, GenesisCodec} {
		// build block
		apricotProposalBlk, err := NewApricotProposal(
			parentID,
			height,
			tx,
		)
		require.NoError(err)

		// parse block
		parsed, err := Parse(cdc, apricotProposalBlk.Bytes)
		require.NoError(err)
		parsedApricotProposalBlk, ok := parsed.(*ApricotProposal)
		require.True(ok)

		// compare content
		require.Equal(apricotProposalBlk.ID, parsedApricotProposalBlk.ID)
		require.Equal(apricotProposalBlk.Bytes, parsedApricotProposalBlk.Bytes)
		require.Equal(apricotProposalBlk.ParentID, parsedApricotProposalBlk.ParentID)
		require.Equal(apricotProposalBlk.Height, parsedApricotProposalBlk.Height)
		require.Equal(apricotProposalBlk.Txs, parsedApricotProposalBlk.Txs)

		banffProposalBlk, err := NewBanffProposalBlock(
			blkTimestamp,
			parentID,
			height,
			tx,
		)
		require.NoError(err)

		// parse block
		parsed, err = Parse(cdc, banffProposalBlk.Bytes)
		require.NoError(err)
		parsedBanffProposalBlk, ok := parsed.(*BanffProposal)
		require.True(ok)

		// compare content
		require.Equal(banffProposalBlk.ID, parsedBanffProposalBlk.ID)
		require.Equal(banffProposalBlk.Bytes, parsedBanffProposalBlk.Bytes)
		require.Equal(banffProposalBlk.ParentID, parsedBanffProposalBlk.ParentID)
		require.Equal(banffProposalBlk.Height, parsedBanffProposalBlk.Height)
		require.Equal(banffProposalBlk.Txs, parsedBanffProposalBlk.Txs)
		require.Equal(banffProposalBlk.Time, parsedBanffProposalBlk.Time)
	}
}

func TestCommitBlock(t *testing.T) {
	// check Apricot commit block can be built and parsed
	require := require.New(t)
	blkTimestamp := time.Now()
	parentID := ids.ID{'p', 'a', 'r', 'e', 'n', 't', 'I', 'D'}
	height := uint64(2022)

	for _, cdc := range []codec.Manager{Codec, GenesisCodec} {
		// build block
		apricotCommitBlk, err := NewApricotCommitBlock(parentID, height)
		require.NoError(err)

		// parse block
		parsed, err := Parse(cdc, apricotCommitBlk.Bytes)
		require.NoError(err)
		parsedApricotCommitBlk, ok := parsed.(*ApricotCommit)
		require.True(ok)

		// compare content
		require.Equal(apricotCommitBlk.ID, parsedApricotCommitBlk.ID)
		require.Equal(apricotCommitBlk.Bytes, parsedApricotCommitBlk.Bytes)
		require.Equal(apricotCommitBlk.ParentID, parsedApricotCommitBlk.ParentID)
		require.Equal(apricotCommitBlk.Height, parsedApricotCommitBlk.Height)

		// check that banff commit block can be built and parsed
		banffCommitBlk, err := NewBanffCommit(blkTimestamp, parentID, height)
		require.NoError(err)

		// parse block
		parsed, err = Parse(cdc, banffCommitBlk.Bytes)
		require.NoError(err)
		parsedBanffCommitBlk, ok := parsed.(*BanffCommit)
		require.True(ok)

		// compare content
		require.Equal(banffCommitBlk.ID, parsedBanffCommitBlk.ID)
		require.Equal(banffCommitBlk.Bytes, parsedBanffCommitBlk.Bytes)
		require.Equal(banffCommitBlk.ParentID, parsedBanffCommitBlk.ParentID)
		require.Equal(banffCommitBlk.Height, parsedBanffCommitBlk.Height)
		require.Equal(banffCommitBlk.Time, parsedBanffCommitBlk.Time)
	}
}

func TestAbortBlock(t *testing.T) {
	// check Apricot abort block can be built and parsed
	require := require.New(t)
	blkTimestamp := time.Now()
	parentID := ids.ID{'p', 'a', 'r', 'e', 'n', 't', 'I', 'D'}
	height := uint64(2022)

	for _, cdc := range []codec.Manager{Codec, GenesisCodec} {
		// build block
		apricotAbortBlk, err := NewApricotAbort(parentID, height)
		require.NoError(err)

		// parse block
		parsed, err := Parse(cdc, apricotAbortBlk.Bytes)
		require.NoError(err)
		parsedApricotAbortBlk, ok := parsed.(*ApricotAbort)
		require.True(ok)

		// compare content
		require.Equal(apricotAbortBlk.ID, parsedApricotAbortBlk.ID)
		require.Equal(apricotAbortBlk.Bytes, parsedApricotAbortBlk.Bytes)
		require.Equal(apricotAbortBlk.ParentID, parsedApricotAbortBlk.ParentID)
		require.Equal(apricotAbortBlk.Height, parsedApricotAbortBlk.Height)

		// check that banff abort block can be built and parsed
		banffAbortBlk, err := NewBanffAbort(blkTimestamp, parentID, height)
		require.NoError(err)

		// parse block
		parsed, err = Parse(cdc, banffAbortBlk.Bytes)
		require.NoError(err)
		parsedBanffAbort, ok := parsed.(*BanffAbort)
		require.True(ok)

		// compare content
		require.Equal(banffAbortBlk.ID, parsedBanffAbort.ID)
		require.Equal(banffAbortBlk.Bytes, parsedBanffAbort.Bytes)
		require.Equal(banffAbortBlk.ParentID, parsedBanffAbort.ParentID)
		require.Equal(banffAbortBlk.Height, parsedBanffAbort.Height)

		// timestamp check for banff blocks only
		require.IsType(&BanffAbort{}, parsed)
		parsedBanffAbortBlk := parsed.(*BanffAbort)
		require.Equal(banffAbortBlk.Time, parsedBanffAbortBlk.Time)
	}
}

func TestAtomicBlock(t *testing.T) {
	// check atomic block can be built and parsed
	require := require.New(t)
	parentID := ids.ID{'p', 'a', 'r', 'e', 'n', 't', 'I', 'D'}
	height := uint64(2022)
	tx, err := testAtomicTx()
	require.NoError(err)

	for _, cdc := range []codec.Manager{Codec, GenesisCodec} {
		// build block
		atomicBlk, err := NewApricotAtomic(
			parentID,
			height,
			tx,
		)
		require.NoError(err)

		// parse block
		parsed, err := Parse(cdc, atomicBlk.Bytes)
		require.NoError(err)
		parsedAtomicBlk, ok := parsed.(*ApricotAtomic)
		require.True(ok)

		// compare content
		require.Equal(atomicBlk.ID, parsedAtomicBlk.ID)
		require.Equal(atomicBlk.Bytes, parsedAtomicBlk.Bytes)
		require.Equal(atomicBlk.ParentID, parsedAtomicBlk.ParentID)
		require.Equal(atomicBlk.Height, parsedAtomicBlk.Height)
		require.Equal(atomicBlk.Tx, parsedAtomicBlk.Tx)
	}
}

func testAtomicTx() (*txs.Tx, error) {
	utx := &txs.ImportTx{
		BaseTx: txs.BaseTx{BaseTx: avax.BaseTx{
			NetworkID:    10,
			BlockchainID: ids.ID{'c', 'h', 'a', 'i', 'n', 'I', 'D'},
			Outs: []*avax.TransferableOutput{{
				Asset: avax.Asset{ID: ids.ID{'a', 's', 's', 'e', 'r', 't'}},
				Out: &secp256k1fx.TransferOutput{
					Amt: uint64(1234),
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs:     []ids.ShortID{preFundedKeys[0].PublicKey().Address()},
					},
				},
			}},
			Ins: []*avax.TransferableInput{{
				UTXOID: avax.UTXOID{
					TxID:        ids.ID{'t', 'x', 'I', 'D'},
					OutputIndex: 2,
				},
				Asset: avax.Asset{ID: ids.ID{'a', 's', 's', 'e', 'r', 't'}},
				In: &secp256k1fx.TransferInput{
					Amt:   uint64(5678),
					Input: secp256k1fx.Input{SigIndices: []uint32{0}},
				},
			}},
			Memo: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		}},
		SourceChain: ids.ID{'c', 'h', 'a', 'i', 'n'},
		ImportedInputs: []*avax.TransferableInput{{
			UTXOID: avax.UTXOID{
				TxID:        ids.Empty.Prefix(1),
				OutputIndex: 1,
			},
			Asset: avax.Asset{ID: ids.ID{'a', 's', 's', 'e', 'r', 't'}},
			In: &secp256k1fx.TransferInput{
				Amt:   50000,
				Input: secp256k1fx.Input{SigIndices: []uint32{0}},
			},
		}},
	}
	signers := [][]*secp256k1.PrivateKey{{preFundedKeys[0]}}
	return txs.NewSigned(utx, txs.Codec, signers)
}

func testDecisionTxs() ([]*txs.Tx, error) {
	countTxs := 2
	decisionTxs := make([]*txs.Tx, 0, countTxs)
	for i := 0; i < countTxs; i++ {
		// Create the tx
		utx := &txs.CreateChainTx{
			BaseTx: txs.BaseTx{BaseTx: avax.BaseTx{
				NetworkID:    10,
				BlockchainID: ids.ID{'c', 'h', 'a', 'i', 'n', 'I', 'D'},
				Outs: []*avax.TransferableOutput{{
					Asset: avax.Asset{ID: ids.ID{'a', 's', 's', 'e', 'r', 't'}},
					Out: &secp256k1fx.TransferOutput{
						Amt: uint64(1234),
						OutputOwners: secp256k1fx.OutputOwners{
							Threshold: 1,
							Addrs:     []ids.ShortID{preFundedKeys[0].PublicKey().Address()},
						},
					},
				}},
				Ins: []*avax.TransferableInput{{
					UTXOID: avax.UTXOID{
						TxID:        ids.ID{'t', 'x', 'I', 'D'},
						OutputIndex: 2,
					},
					Asset: avax.Asset{ID: ids.ID{'a', 's', 's', 'e', 'r', 't'}},
					In: &secp256k1fx.TransferInput{
						Amt:   uint64(5678),
						Input: secp256k1fx.Input{SigIndices: []uint32{0}},
					},
				}},
				Memo: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			}},
			SubnetID:    ids.ID{'s', 'u', 'b', 'n', 'e', 't', 'I', 'D'},
			ChainName:   "a chain",
			VMID:        ids.GenerateTestID(),
			FxIDs:       []ids.ID{ids.GenerateTestID()},
			GenesisData: []byte{'g', 'e', 'n', 'D', 'a', 't', 'a'},
			SubnetAuth:  &secp256k1fx.Input{SigIndices: []uint32{1}},
		}

		signers := [][]*secp256k1.PrivateKey{{preFundedKeys[0]}}
		tx, err := txs.NewSigned(utx, txs.Codec, signers)
		if err != nil {
			return nil, err
		}
		decisionTxs = append(decisionTxs, tx)
	}
	return decisionTxs, nil
}

func testProposalTx() (*txs.Tx, error) {
	utx := &txs.RewardValidatorTx{
		TxID: ids.ID{'r', 'e', 'w', 'a', 'r', 'd', 'I', 'D'},
	}

	signers := [][]*secp256k1.PrivateKey{{preFundedKeys[0]}}
	return txs.NewSigned(utx, txs.Codec, signers)
}
