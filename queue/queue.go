/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package queue

import (
	"github.com/sgoware/ds/queue/arrayqueue"
)

type Group[T any] struct{}

func (g *Group[T]) ArrayQueue() *arrayqueue.Group[T] {
	return &arrayqueue.Group[T]{}
}
