/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package container

type Container[T any] interface {
	// Empty return true if container does not contain any elements
	Empty() bool
	// Len return number of elements within the container
	Len() int
	// Clear remove all elements from the container
	Clear()
	// Values return all elements of the container
	Values() []T
	// String return a string representation of the container
	String() string
}
