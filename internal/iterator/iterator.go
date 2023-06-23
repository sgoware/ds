/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package iterator

// ForwardIterator is an iterator that can only move forward through a sequence of elements.
type ForwardIterator interface {
	// Value returns current elements' value.
	Value() any

	// First moves the iterator to the first element.
	// returns false if container does not have any element.
	First() bool

	// Next moves the iterator to the next element.
	// returns false if container does not have a next element.
	// if Next() is called for the first time, it will point the iterator to the first element if it exists.
	Next() bool

	// Reset resets the iterator to its initial state.
	Reset()
}

// ReverseIterator is an iterator that moves backwards through a sequence of elements.
type ReverseIterator interface {
	// Value returns current elements' value.
	Value() any

	// Last moves the iterator to the last element.
	// returns false if container does not have any elements.
	Last() bool

	// Prev move the iterator to the previous element.
	// returns false if container does not have a previous element.
	// if Prev() is called for the first time, it will point the iterator to the last element if it exists.
	Prev() bool

	// Reset resets the iterator to its initial state.
	Reset()
}

// BidirectionalIterator is an iterator that can move forward or backward one element at a time.
type BidirectionalIterator interface {
	ForwardIterator
	ReverseIterator
}

// RandomAccessIterator is an iterator that can move forward or backward any number of elements at a time
// and should support O(1) time complexity for random access to elements.
type RandomAccessIterator interface {
	BidirectionalIterator
	To(index int) bool
}
