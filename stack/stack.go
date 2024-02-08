/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package stack

import (
	"github.com/sgoware/ds/stack/arraystack"
)

type Group[T any] struct{}

func (g *Group[T]) ArrayStack() *arraystack.Group[T] {
	return &arraystack.Group[T]{}
}
