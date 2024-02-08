/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package iterator

// ForwardIteratorWithIndex is a forward iterator whose values can be fetched by an index.
type ForwardIteratorWithIndex[T any] interface {
	ForwardIterator[T]
	Index() int
}

// ReverseIteratorWithIndex is a reverse iterator whose values can be fetched by an index.
type ReverseIteratorWithIndex[T any] interface {
	ReverseIterator[T]
	Index() int
}

// BidirectionalIteratorWithIndex is a bidirectional iterator whose values can be fetched by an index.
type BidirectionalIteratorWithIndex[T any] interface {
	ForwardIterator[T]
	ReverseIterator[T]
	Index() int
}

// RandomAccessIteratorWithIndex is a random access iterator whose values can be fetched by an index.
type RandomAccessIteratorWithIndex[T any] interface {
	RandomAccessIterator[T]
	Index() int
}
