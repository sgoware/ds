/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package iterator

// ForwardIteratorWithKey is a forward iterator whose elements are k-v pairs.
type ForwardIteratorWithKey[T comparable] interface {
	ForwardIterator[T]
	Key() T
}

// ReverseIteratorKey is a forward iterator whose elements are k-v pairs.
type ReverseIteratorKey[T comparable] interface {
	ReverseIterator[T]
	Key() T
}

// BidirectionalIteratorKey is a forward iterator whose elements are k-v pairs.
type BidirectionalIteratorKey[T comparable] interface {
	ForwardIterator[T]
	ReverseIterator[T]
	Key() T
}

// RandomIteratorKey is a forward iterator whose elements are k-v pairs.
type RandomIteratorKey[T comparable] interface {
	RandomAccessIterator[T]
	Key() T
}
