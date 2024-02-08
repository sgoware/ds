/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package internal

import (
	"github.com/sgoware/ds/internal/container"
	"github.com/sgoware/ds/utils/comparator"
)

type List[T any] interface {
	container.Container[T]

	// Insert given values after the element by the index
	Insert(index int, values ...T) bool
	PushFront(value ...T)

	// Remove value by index
	Remove(index int) bool
	PopFront() bool

	// Set value to the element by the index
	Set(index int, value T) bool
	// Swap two element by the indies
	Swap(index1, index2 int) bool

	// Sort list by the given comparator
	Sort(comparator comparator.Comparator[T])
}
