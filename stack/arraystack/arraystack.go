/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package arraystack

import (
	"fmt"
	"github.com/sgoware/ds/list/arraylist"
	"strings"
)

type Group[T any] struct{}

const (
	name = "ArrayStack"
)

//var _ internal.Stack = (*Stack)(nil)

type Stack[T any] struct {
	list *arraylist.List[T]
}

func New[T any](values ...T) *Stack[T] {
	return &Stack[T]{list: arraylist.New(values...)}
}

func (g *Group[T]) New(values ...T) *Stack[T] {
	return New(values...)
}

func (q *Stack[T]) Empty() bool {
	return q.list.Empty()
}

func (q *Stack[T]) Len() int {
	return q.list.Len()
}

func (q *Stack[T]) Clear() {
	q.list.Clear()
}

func (q *Stack[T]) Values() []T {
	return q.list.Values()
}

func (q *Stack[T]) String() string {
	str := name + "\n"

	values := make([]string, 0, q.list.Len())

	for _, v := range q.list.Values() {
		values = append(values, fmt.Sprintf("%v", v))
	}

	str += strings.Join(values, ", ")

	return str
}

func (q *Stack[T]) Push(values ...T) {
	q.list.PushBack(values...)
}

func (q *Stack[T]) Pop() bool {
	return q.list.PopBack()
}

func (q *Stack[T]) Top() (T, bool) {
	return q.list.Back()
}
