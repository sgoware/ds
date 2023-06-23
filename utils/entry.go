/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package utils

import (
	"github.com/sgoware/ds/utils/comparator"
	"github.com/sgoware/ds/utils/sort"
)

type Group struct{}

var (
	insComparator = comparator.Group{}
	insSort       = sort.Group{}
)

func (g *Group) Comparator() *comparator.Group {
	return &insComparator
}

func (g *Group) Sort() *sort.Group {
	return &insSort
}
