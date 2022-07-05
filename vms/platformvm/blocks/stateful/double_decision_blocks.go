// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stateful

import (
	"fmt"
)

// doubleDecisionBlock contains the accept for a pair of blocks
type doubleDecisionBlock struct {
	decisionBlock
}

func (ddb *doubleDecisionBlock) acceptParent() error {
	blkID := ddb.baseBlk.ID()
	ddb.txExecutorBackend.Ctx.Log.Verbo("Accepting block with ID %s", blkID)

	parentIntf, err := ddb.parentBlock()
	if err != nil {
		return err
	}

	parent, ok := parentIntf.(*ProposalBlock)
	if !ok {
		ddb.txExecutorBackend.Ctx.Log.Error("double decision block should only follow a proposal block")
		return fmt.Errorf("expected Proposal block but got %T", parentIntf)
	}

	parent.commonBlock.accept()
	parent.verifier.AddStatelessBlock(parent.ProposalBlock, parent.Status())
	if err := parent.verifier.MarkAccepted(parent.ProposalBlock); err != nil {
		return fmt.Errorf("failed to accept proposal block %s: %w",
			parent.ID(),
			err,
		)
	}

	return nil
}

func (ddb *doubleDecisionBlock) updateState() error {
	parentIntf, err := ddb.parentBlock()
	if err != nil {
		return err
	}

	parent, ok := parentIntf.(*ProposalBlock)
	if !ok {
		ddb.txExecutorBackend.Ctx.Log.Error("double decision block should only follow a proposal block")
		return fmt.Errorf("expected Proposal block but got %T", parentIntf)
	}

	// Update the state of the chain in the database
	ddb.onAcceptState.Apply(ddb.verifier.GetState())
	if err := ddb.verifier.Commit(); err != nil {
		return fmt.Errorf("failed to commit vm's state: %w", err)
	}

	for _, child := range ddb.children {
		child.setBaseState()
	}
	if ddb.onAcceptFunc != nil {
		ddb.onAcceptFunc()
	}

	// remove this block and its parent from memory
	parent.free()
	ddb.free()
	return nil
}