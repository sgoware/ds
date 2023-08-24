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

type Group struct{}

var _ internal.Queue = (*Queue)(nil)

type Queue struct {
	list *arraylist.List
}

func New(values ...any) *Queue {
	return &Queue{list: arraylist.New(values...)}
}

func (g *Group) New(values ...any) *Queue {
	return New(values...)
}

func (q *Queue) Empty() bool {
	return q.list.Empty()
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) Clear() {
	q.list.Clear()
}

func (q *Queue) Values() []any {
	return q.list.Values()
}

func (q *Queue) String() string {
	str := "ArrayQueue\n"

	values := make([]string, 0, q.list.Len())

	for _, v := range q.list.Values() {
		values = append(values, fmt.Sprintf("%v", v))
	}

	str += strings.Join(values, ", ")

	return str
}

func (q *Queue) Push(values ...any) {
	q.list.PushBack(values...)
}

func (q *Queue) Pop() bool {
	return q.list.PopFront()
}

func (q *Queue) Top() (any, bool) {
	return q.list.Front()
}
