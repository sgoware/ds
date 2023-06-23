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

type Group struct{}

// Assert internal.List interface implementation
var _ internal.List = (*List)(nil)

// List stores elements using a slice
type List struct {
	elements []any
	size     int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New(values ...any) *List {
	l := &List{}

	if len(values) > 0 {
		l.PushBack(values...)
	}

	return l
}

func (g *Group) New(values ...any) *List {
	return New(values...)
}

func (l *List) Empty() bool {
	return l.size == 0
}

func (l *List) Len() int {
	return l.size
}

func (l *List) Clear() {
	l.size = 0
	l.elements = []any{}
}

func (l *List) Values() []any {
	return l.elements
}

func (l *List) String() string {
	str := "ArrayList\n"

	var values []string

	for _, v := range l.elements[:l.size] {
		values = append(values, fmt.Sprintf("%v", v))
	}

	str += strings.Join(values, ", ")

	return str
}

func (l *List) Insert(index int, values ...any) bool {
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

func (l *List) PushFront(values ...any) {
	l.growBy(len(values))

	size := len(values)

	copy(l.elements[size:], l.elements[:l.size])
	copy(l.elements[:], values[:])

	l.size += size
}

func (l *List) PushBack(values ...any) {
	l.growBy(len(values))

	for _, value := range values {
		l.elements[l.size] = value
		l.size++
	}
}

func (l *List) Remove(index int) bool {
	if !l.withinRange(index) {
		return false
	}

	l.elements[index] = nil
	copy(l.elements[index:], l.elements[index+1:l.size])
	l.size--

	l.shrink()

	return true
}

func (l *List) PopFront() bool {
	return l.Remove(0)
}

func (l *List) PopBack() bool {
	return l.Remove(l.size - 1)
}

func (l *List) Set(index int, value any) bool {
	if !l.withinRange(index) {
		return false
	}

	l.elements[index] = value

	return true
}

func (l *List) Swap(i, j int) bool {
	if l.withinRange(i) && l.withinRange(j) {
		l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
		return true
	}

	return false
}

// Get value by index
func (l *List) Get(index int) (any, bool) {
	if !l.withinRange(index) {
		return nil, false
	}

	return l.elements[index], true
}

func (l *List) Front() (any, bool) {
	return l.Get(0)
}

func (l *List) Back() (any, bool) {
	return l.Get(l.size - 1)
}

func (l *List) Sort(comparator comparator.Comparator) {
	sort.Sort(l.elements[:l.size], comparator)
}

func (l *List) withinRange(index int) bool {
	return index >= 0 && index < l.size
}

func (l *List) resize(cap int) {
	newElements := make([]any, cap)
	copy(newElements, l.elements)
	l.elements = newElements
}

func (l *List) growBy(n int) {
	curCap := cap(l.elements)

	if l.size+n >= curCap {
		newCapacity := int(growthFactor * float32(curCap+n))
		l.resize(newCapacity)
	}
}

func (l *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}

	curCap := cap(l.elements)

	if l.size <= int(float32(curCap)*shrinkFactor) {
		l.resize(l.size)
	}
}
