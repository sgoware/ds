/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package enumerator

// enumerator can execute the given function for each element and returns a bool value for a certain logic.

// IndexEnumerator is an enumerator that provides functions for values which can be fetched by an index.
type IndexEnumerator interface {
	// Each call the given function once for each element.
	Each(func(index int, value any))

	// Any returns true if the function ever returns true for any element.
	Any(func(index int, value any) bool) bool

	// All returns true if the function returns true for all elements.
	All(func(index int, value any) bool) bool

	// Find returns (index, value) for which the function first returns true,
	// otherwise return -1, nil if no element matches the criteria.
	Find(func(index int, value any) bool) (int, any)
}

// KeyEnumerator is an enumerator that provides functions for values whose elements are k-v pairs.
type KeyEnumerator interface {
	// Each call the given function once for each element
	Each(func(key, value any))

	// Any returns true if the function ever returns true for any element.
	Any(func(key, value any) bool) bool

	// All returns true if the function returns true for all elements.
	All(func(key, value any) bool) bool

	// Find returns the first (key, value) for which the function is true,
	// otherwise return -1, nil if no element matches the criteria.
	Find(func(key, value any) bool) (any, any)
}
