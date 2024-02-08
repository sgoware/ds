/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package sort

import (
	"github.com/sgoware/ds/utils/comparator"
	"sort"
)

func Sort[T any](values []T, comparator comparator.Comparator[T]) {
	sort.Sort(sortable[T]{values, comparator})
}
