/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package arrayqueue

import (
	"fmt"
	"github.com/sgoware/ds/list/arraylist"
	"github.com/sgoware/ds/queue/internal"
	"strings"
)

type Group[T any] struct{}

const (
	name = "ArrayQueue"
)

var _ internal.Queue[int] = (*Queue[int])(nil)

type Queue[T any] struct {
	list *arraylist.List[T]
}

func New[T any](values ...T) *Queue[T] {
	return &Queue[T]{list: arraylist.New(values...)}
}

func (g *Group[T]) New(values ...T) *Queue[T] {
	return New(values...)
}

func (q *Queue[T]) Empty() bool {
	return q.list.Empty()
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Clear() {
	q.list.Clear()
}

func (q *Queue[T]) Values() []T {
	return q.list.Values()
}

func (q *Queue[T]) String() string {
	str := name + "\n"

	values := make([]string, 0, q.list.Len())

	for _, v := range q.list.Values() {
		values = append(values, fmt.Sprintf("%v", v))
	}

	str += strings.Join(values, ", ")

	return str
}

func (q *Queue[T]) Push(values ...T) {
	q.list.PushBack(values...)
}

func (q *Queue[T]) Pop() bool {
	return q.list.PopFront()
}

func (q *Queue[T]) Top() (T, bool) {
	return q.list.Front()
}
