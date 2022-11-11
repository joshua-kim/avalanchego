// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracedvm

import (
	"context"

	"go.opentelemetry.io/otel/attribute"

	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
)

var _ snowstorm.Tx = (*tracedTx)(nil)

type tracedTx struct {
	snowstorm.Tx

	vm *vertexVM
}

func (t *tracedTx) Verify(ctx context.Context) error {
	ctx, span := t.vm.tracer.Start(ctx, "tracedTx.Verify", oteltrace.WithAttributes(
		attribute.Stringer("txID", t.ID()),
	))
	defer span.End()

	return t.Tx.Verify(ctx)
}

func (t *tracedTx) Accept(ctx context.Context) error {
	ctx, span := t.vm.tracer.Start(ctx, "tracedTx.Accept", oteltrace.WithAttributes(
		attribute.Stringer("txID", t.ID()),
	))
	defer span.End()

	return t.Tx.Accept(ctx)
}

func (t *tracedTx) Reject(ctx context.Context) error {
	ctx, span := t.vm.tracer.Start(ctx, "tracedTx.Accept", oteltrace.WithAttributes(
		attribute.Stringer("txID", t.ID()),
	))
	defer span.End()

	return t.Tx.Reject(ctx)
}
