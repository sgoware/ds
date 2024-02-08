/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package enumerator

// enumerator can execute the given function for each element and returns a bool value for a certain logic.

// IndexEnumerator is an enumerator that provides functions for values which can be fetched by an index.
type IndexEnumerator[T any] interface {
	// Each call the given function once for each element.
	Each(func(index int, value T))

	// Any returns true if the function ever returns true for any element.
	Any(func(index int, value T) bool) bool

	// All returns true if the function returns true for all elements.
	All(func(index int, value T) bool) bool

	// Find returns (index, value) for which the function first returns true,
	// otherwise return -1, nil if no element matches the criteria.
	Find(func(index int, value T) bool) (int, T)
}

// KeyEnumerator is an enumerator that provides functions for values whose elements are k-v pairs.
type KeyEnumerator[K comparable, V any] interface {
	// Each call the given function once for each element
	Each(func(key K, value V))

	// Any returns true if the function ever returns true for any element.
	Any(func(key K, value V) bool) bool

	// All returns true if the function returns true for all elements.
	All(func(key K, value V) bool) bool

	// Find returns the first (key, value) for which the function is true,
	// otherwise return -1, nil if no element matches the criteria.
	Find(func(key K, value V) bool) (K, V)
}
