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

func Utils() *utils.Group {
	return &utils.Group{}
}

func List[T any]() *list.Group[T] {
	return &list.Group[T]{}
}

func Queue[T any]() *queue.Group[T] {
	return &queue.Group[T]{}
}

func Stack[T any]() *stack.Group[T] {
	return &stack.Group[T]{}
}
