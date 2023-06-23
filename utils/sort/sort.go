/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package sort

import (
	"github.com/sgoware/ds/utils/comparator"
	"sort"
)

type Group struct{}

func Sort(values []interface{}, comparator comparator.Comparator) {
	sort.Sort(sortable{values, comparator})
}
