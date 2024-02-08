/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package utils

import (
	"github.com/sgoware/ds/utils/comparator"
)

type Group struct{}

func (g *Group) Comparator() *comparator.Group {
	return &comparator.Group{}
}
