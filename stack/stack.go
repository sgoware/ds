/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package stack

import (
	"github.com/sgoware/ds/stack/arraystack"
)

type Group struct{}

var insArrayStack = arraystack.Group{}

func (g *Group) ArrayStack() *arraystack.Group {
	return &insArrayStack
}
