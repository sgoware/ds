/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package list

import (
	"github.com/sgoware/ds/list/arraylist"
)

type Group[T any] struct{}

func (g *Group[T]) ArrayList() *arraylist.Group[T] {
	return &arraylist.Group[T]{}
}
