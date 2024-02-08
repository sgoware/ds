/*
 * Copyright (c) 2023 sgoware. All rights reserved.
 * Licensed under BSD 3-Clause "New" or "Revised" License that can be found in the LICENSE file.
 */

package arraylist

import (
	"fmt"
	"github.com/sgoware/ds/list/internal"
	"github.com/sgoware/ds/utils/comparator"
	"github.com/sgoware/ds/utils/sort"
	"strings"
)

type Group[T any] struct{}

const (
	name = "ArrayList"
)

// Assert internal.List interface implementation
var _ internal.List[int] = (*List[int])(nil)

// List stores elements using a slice
type List[T any] struct {
	elements []T
	size     int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New[T any](values ...T) *List[T] {
	l := &List[T]{}

	if len(values) > 0 {
		l.PushBack(values...)
	}

	return l
}

func (g *Group[T]) New(values ...T) *List[T] {
	return New(values...)
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Len() int {
	return l.size
}

func (l *List[T]) Clear() {
	l.size = 0
	l.elements = []T{}
}

func (l *List[T]) Values() []T {
	return l.elements
}

func (l *List[T]) String() string {
	str := name + "\n"

	var values []string

	for _, v := range l.elements[:l.size] {
		values = append(values, fmt.Sprintf("%v", v))
	}

	str += strings.Join(values, ", ")

	return str
}

func (l *List[T]) Insert(index int, values ...T) bool {
	if !l.withinRange(index) {
		return false
	}

	size := len(values)
	l.growBy(size)
	l.size += size

	copy(l.elements[index+size:], l.elements[index:l.size-size])
	copy(l.elements[index:], values)

	return true
}

func (l *List[T]) PushFront(values ...T) {
	l.growBy(len(values))

	size := len(values)

	copy(l.elements[size:], l.elements[:l.size])
	copy(l.elements[:], values[:])

	l.size += size
}

func (l *List[T]) PushBack(values ...T) {
	l.growBy(len(values))

	for _, value := range values {
		l.elements[l.size] = value
		l.size++
	}
}

func (l *List[T]) Remove(index int) bool {
	if !l.withinRange(index) {
		return false
	}

	clear(l.elements[index : index+1])
	copy(l.elements[index:], l.elements[index+1:l.size])
	l.size--

	l.shrink()

	return true
}

func (l *List[T]) PopFront() bool {
	return l.Remove(0)
}

func (l *List[T]) PopBack() bool {
	return l.Remove(l.size - 1)
}

func (l *List[T]) Set(index int, value T) bool {
	if !l.withinRange(index) {
		return false
	}

	l.elements[index] = value

	return true
}

func (l *List[T]) Swap(i, j int) bool {
	if l.withinRange(i) && l.withinRange(j) {
		l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
		return true
	}

	return false
}

// Get value by index
func (l *List[T]) Get(index int) (any, bool) {
	if !l.withinRange(index) {
		return nil, false
	}

	return l.elements[index], true
}

func (l *List[T]) Front() (any, bool) {
	return l.Get(0)
}

func (l *List[T]) Back() (any, bool) {
	return l.Get(l.size - 1)
}

func (l *List[T]) Sort(comparator comparator.Comparator[T]) {
	sort.Sort[T](l.elements[:l.size], comparator)
}

func (l *List[T]) withinRange(index int) bool {
	return index >= 0 && index < l.size
}

func (l *List[T]) resize(cap int) {
	newElements := make([]T, cap)
	copy(newElements, l.elements)
	l.elements = newElements
}

func (l *List[T]) growBy(n int) {
	curCap := cap(l.elements)

	if l.size+n >= curCap {
		newCapacity := int(growthFactor * float32(curCap+n))
		l.resize(newCapacity)
	}
}

func (l *List[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}

	curCap := cap(l.elements)

	if l.size <= int(float32(curCap)*shrinkFactor) {
		l.resize(l.size)
	}
}
