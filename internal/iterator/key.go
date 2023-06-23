/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package iterator

// ForwardIteratorWithKey is a forward iterator whose elements are k-v pairs.
type ForwardIteratorWithKey interface {
	ForwardIterator
	Key() any
}

// ReverseIteratorKey is a forward iterator whose elements are k-v pairs.
type ReverseIteratorKey interface {
	ReverseIterator
	Key() any
}

// BidirectionalIteratorKey is a forward iterator whose elements are k-v pairs.
type BidirectionalIteratorKey interface {
	ForwardIterator
	ReverseIterator
	Key() any
}

// RandomIteratorKey is a forward iterator whose elements are k-v pairs.
type RandomIteratorKey interface {
	RandomAccessIterator
	Key() any
}
