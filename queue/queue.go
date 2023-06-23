/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package queue

import (
	"github.com/sgoware/ds/queue/arrayqueue"
)

type Group struct{}

var insArrayQueue = arrayqueue.Group{}

func (g *Group) ArrayQueue() *arrayqueue.Group {
	return &insArrayQueue
}
