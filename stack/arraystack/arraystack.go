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

type Group struct{}

//var _ internal.Stack = (*Stack)(nil)

type Stack struct {
	list *arraylist.List
}

func New(values ...any) *Stack {
	return &Stack{list: arraylist.New(values...)}
}

func (g *Group) New(values ...any) *Stack {
	return New(values...)
}

func (q *Stack) Empty() bool {
	return q.list.Empty()
}

func (q *Stack) Len() int {
	return q.list.Len()
}

func (q *Stack) Clear() {
	q.list.Clear()
}

func (q *Stack) Values() []any {
	return q.list.Values()
}

func (q *Stack) String() string {
	str := "ArrayQueue\n"

	values := make([]string, 0, q.list.Len())

	for _, v := range q.list.Values() {
		values = append(values, fmt.Sprintf("%v", v))
	}

	str += strings.Join(values, ", ")

	return str
}

func (q *Stack) Push(values ...any) {
	q.list.PushBack(values...)
}

func (q *Stack) Pop() bool {
	return q.list.PopBack()
}

func (q *Stack) Top() (any, bool) {
	return q.list.Back()
}
