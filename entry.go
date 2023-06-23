/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package ds

import (
	"github.com/sgoware/ds/list"
	"github.com/sgoware/ds/queue"
	"github.com/sgoware/ds/stack"
	"github.com/sgoware/ds/utils"
)

var (
	insUtils = utils.Group{}
	insList  = list.Group{}
	insQueue = queue.Group{}
	insStack = stack.Group{}
)

func Utils() *utils.Group {
	return &insUtils
}

func List() *list.Group {
	return &insList
}

func Queue() *queue.Group {
	return &insQueue
}

func Stack() *stack.Group {
	return &insStack
}
