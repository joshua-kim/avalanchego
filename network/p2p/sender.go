// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p2p

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	_ AppSender = (*Sender)(nil)
	_ AppSender = (*FakeSender)(nil)
)

type AppRequestSender interface {
	SendAppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, bytes []byte) error
}

type AppResponseSender interface {
	SendAppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, bytes []byte) error
}

type AppGossipSender interface {
	SendAppGossip(ctx context.Context, bytes []byte) error
	SendAppGossipSpecific(ctx context.Context, nodeID ids.NodeID, bytes []byte) error
}

type CrossChainAppRequestSender interface {
	SendCrossChainAppRequest(ctx context.Context, chainID ids.ID, requestID uint32, bytes []byte) error
}

type CrossChainAppResponseSender interface {
	SendCrossChainAppResponse(ctx context.Context, chainID ids.ID, requestID uint32, bytes []byte) error
}

type AppSender interface {
	AppRequestSender
	AppResponseSender
	AppGossipSender
	CrossChainAppRequestSender
	CrossChainAppResponseSender
}

func NewSender(sender common.AppSender) *Sender {
	return &Sender{
		sender: sender,
	}
}

type Sender struct {
	sender common.AppSender
}

func (s Sender) SendAppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, bytes []byte) error {
	return s.sender.SendAppRequest(ctx, set.Of(nodeID), requestID, bytes)
}

func (s Sender) SendAppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, bytes []byte) error {
	return s.sender.SendAppResponse(ctx, nodeID, requestID, bytes)
}

func (s Sender) SendAppGossip(ctx context.Context, bytes []byte) error {
	return s.sender.SendAppGossip(ctx, bytes)
}

func (s Sender) SendAppGossipSpecific(ctx context.Context, nodeID ids.NodeID, bytes []byte) error {
	return s.sender.SendAppGossipSpecific(ctx, set.Of(nodeID), bytes)
}

func (s Sender) SendCrossChainAppRequest(ctx context.Context, chainID ids.ID, requestID uint32, bytes []byte) error {
	return s.sender.SendCrossChainAppRequest(ctx, chainID, requestID, bytes)
}

func (s Sender) SendCrossChainAppResponse(ctx context.Context, chainID ids.ID, requestID uint32, bytes []byte) error {
	return s.sender.SendCrossChainAppResponse(ctx, chainID, requestID, bytes)
}

type FakeSender struct {
	SentAppRequest, SentAppResponse,
	SentAppGossip, SentAppGossipSpecific,
	SentCrossChainAppRequest, SentCrossChainAppResponse chan []byte
}

func (f FakeSender) SendAppRequest(_ context.Context, _ ids.NodeID, _ uint32, bytes []byte) error {
	if f.SentAppRequest == nil {
		return nil
	}

	f.SentAppRequest <- bytes
	return nil
}

func (f FakeSender) SendAppResponse(_ context.Context, _ ids.NodeID, _ uint32, bytes []byte) error {
	if f.SentAppResponse == nil {
		return nil
	}

	f.SentAppResponse <- bytes
	return nil
}

func (f FakeSender) SendAppGossip(_ context.Context, bytes []byte) error {
	if f.SentAppGossip == nil {
		return nil
	}

	f.SentAppGossip <- bytes
	return nil
}

func (f FakeSender) SendAppGossipSpecific(_ context.Context, _ ids.NodeID, bytes []byte) error {
	if f.SentAppGossipSpecific == nil {
		return nil
	}

	f.SentAppGossipSpecific <- bytes
	return nil
}

func (f FakeSender) SendCrossChainAppRequest(_ context.Context, _ ids.ID, _ uint32, bytes []byte) error {
	if f.SentCrossChainAppRequest == nil {
		return nil
	}

	f.SentCrossChainAppRequest <- bytes
	return nil
}

func (f FakeSender) SendCrossChainAppResponse(_ context.Context, _ ids.ID, _ uint32, bytes []byte) error {
	if f.SentCrossChainAppResponse == nil {
		return nil
	}

	f.SentCrossChainAppResponse <- bytes
	return nil
}

type MockSender struct {
	SendAppRequestF            func(context.Context, ids.NodeID, uint32, []byte) error
	SendAppResponseF           func(context.Context, ids.NodeID, uint32, []byte) error
	SendAppGossipF             func(context.Context, []byte) error
	SendAppGossipSpecificF     func(context.Context, set.Set[ids.NodeID], []byte) error
	SendCrossChainAppRequestF  func(context.Context, ids.ID, uint32, []byte) error
	SendCrossChainAppResponseF func(context.Context, ids.ID, uint32, []byte) error
}

func (m MockSender) SendAppRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, bytes []byte) error {
	if m.SendAppRequestF == nil {
		return nil
	}

	return m.SendAppRequestF(ctx, nodeID, requestID, bytes)
}

func (m MockSender) SendAppResponse(ctx context.Context, nodeID ids.NodeID, requestID uint32, bytes []byte) error {
	if m.SendAppResponseF == nil {
		return nil
	}

	return m.SendAppResponseF(ctx, nodeID, requestID, bytes)
}

func (m MockSender) SendAppGossip(ctx context.Context, bytes []byte) error {
	if m.SendAppGossipF == nil {
		return nil
	}

	return m.SendAppGossipF(ctx, bytes)
}

func (m MockSender) SendAppGossipSpecific(ctx context.Context, nodeIDs set.Set[ids.NodeID], bytes []byte) error {
	if m.SendAppGossipSpecificF == nil {
		return nil
	}

	return m.SendAppGossipSpecificF(ctx, nodeIDs, bytes)
}

func (m MockSender) SendCrossChainAppRequest(ctx context.Context, chainID ids.ID, requestID uint32, bytes []byte) error {
	if m.SendCrossChainAppRequestF == nil {
		return nil
	}

	return m.SendCrossChainAppRequestF(ctx, chainID, requestID, bytes)
}

func (m MockSender) SendCrossChainAppResponse(ctx context.Context, chainID ids.ID, requestID uint32, bytes []byte) error {
	if m.SendCrossChainAppResponseF == nil {
		return nil
	}

	return m.SendCrossChainAppResponseF(ctx, chainID, requestID, bytes)
}
