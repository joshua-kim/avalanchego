// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package heap

import (
	stdheap "container/heap"
	"github.com/ava-labs/avalanchego/utils"
)

var _ stdheap.Interface = (*queue[int])(nil)

// NewQueue returns an empty min-heap ordered by less. See OfQueue for more.
func NewQueue[T any](less func(a, b T) bool) Queue[T] {
	return OfQueue[T](less)
}

// OfQueue returns a pre-initialized heap. A max-heap can be created by negating less.
func OfQueue[T any](less func(a, b T) bool, ts ...T) Queue[T] {
	q := Queue[T]{
		queue: &queue[T]{
			entries: make([]T, 0),
			less:    less,
		},
	}

	for _, t := range ts {
		q.Push(t)
	}

	return q
}

type Queue[T any] struct {
	queue *queue[T]
}

func (q *Queue[T]) Len() int {
	return len(q.queue.entries)
}

func (q *Queue[T]) Push(t T) {
	stdheap.Push(q.queue, t)
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.Len() == 0 {
		return utils.Zero[T](), false
	}

	return stdheap.Pop(q.queue).(T), true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.Len() == 0 {
		return utils.Zero[T](), false
	}

	return q.queue.entries[0], true
}

func (q *Queue[T]) Fix(i int) {
	stdheap.Fix(q.queue, i)
}

type queue[T any] struct {
	entries []T
	less    func(a, b T) bool
}

func (q *queue[T]) Len() int {
	return len(q.entries)
}

func (q *queue[T]) Less(i, j int) bool {
	return q.less(q.entries[i], q.entries[j])
}

func (q *queue[T]) Swap(i, j int) {
	q.entries[i], q.entries[j] = q.entries[j], q.entries[i]
}

func (q *queue[T]) Push(e any) {
	q.entries = append(q.entries, e.(T))
}

func (q *queue[T]) Pop() any {
	end := len(q.entries) - 1

	popped := q.entries[end]
	q.entries[end] = utils.Zero[T]()
	q.entries = q.entries[:end]

	return popped
}
