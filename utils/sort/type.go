/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package sort

import (
	"github.com/sgoware/ds/utils/comparator"
)

type sortable struct {
	values     []interface{}
	comparator comparator.Comparator
}

func (s sortable) Len() int {
	return len(s.values)
}
func (s sortable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s sortable) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}
