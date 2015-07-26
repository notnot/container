// deque.go, jpad 2015

/*
Package deque_int implements an efficient double ended queue to store integers.
An iterator is provided with which forward and backward iteration through the
deque is possible.
*/
package deque_int

import (
	"container/list"
)

const (
	chunkSize   = 32 // benchmark optimum on a 64-bit machine
	chunkCenter = chunkSize / 2
)

//// Deque /////////////////////////////////////////////////////////////////////

// Deque is a double ended queue that can handle items of any type.
type Deque struct {
	chunks list.List
	frontI int // index of item at the front
	backI  int // index of item at the back
	size   int
}

// New returns a pointer to an empty deque.
func New() *Deque {
	deque := Deque{}
	deque.reset()
	deque.chunks.PushBack(make(_Chunk, chunkSize))
	return &deque
}

// PushFront adds an item to the front of the deque.
func (d *Deque) PushFront(item int) {
	// get 'front' chunk
	var chunk _Chunk
	if d.frontI == 0 { // 'front' chunk full?
		// add a new chunk at the front
		chunk = make(_Chunk, chunkSize)
		d.chunks.PushFront(chunk)
		d.frontI = chunkSize
	} else {
		chunk = d.chunks.Front().Value.(_Chunk)
	}

	d.frontI--
	chunk[d.frontI] = item
	d.size++
}

// PushBack adds an item to the back of the deque.
func (d *Deque) PushBack(item int) {
	// get 'back' chunk
	var chunk _Chunk
	if d.backI == chunkSize-1 { // 'back' chunk full?
		// add a new chunk at the back
		chunk = make(_Chunk, chunkSize)
		d.chunks.PushBack(chunk)
		d.backI = -1
	} else {
		chunk = d.chunks.Back().Value.(_Chunk)
	}

	d.backI++
	chunk[d.backI] = item
	d.size++
}

// PopFront removes and returns the item from the front of the deque.
// Returns false when the deque is empty.
func (d *Deque) PopFront() (int, bool) {
	if d.size <= 0 {
		return 0, false
	}
	node := d.chunks.Front()
	chunk := node.Value.(_Chunk)
	item := chunk[d.frontI]
	d.frontI++
	d.size--

	if d.frontI == chunkSize { // 'front' chunk empty?
		if d.size == 0 { // deque is empty, reset it
			d.reset()
		} else {
			d.chunks.Remove(node)
			d.frontI = 0
		}
	}

	return item, true
}

// PopBack removes and returns the item from the back of the deque.
// Returns false when the deque is empty.
func (d *Deque) PopBack() (int, bool) {
	if d.size <= 0 {
		return 0, false
	}
	node := d.chunks.Back()
	chunk := node.Value.(_Chunk)
	item := chunk[d.backI]
	d.backI--
	d.size--

	if d.backI == -1 { // 'back' chunk empty?
		if d.size == 0 { // deque is empty, reset it
			d.reset()
		} else {
			d.chunks.Remove(node)
			d.backI = chunkSize - 1
		}
	}

	return item, true
}

// FrontItem returns the item at the front of the deque.
func (d *Deque) FrontItem() (int, bool) {
	if d.size <= 0 {
		return 0, false
	} else {
		return d.chunks.Front().Value.(_Chunk)[d.frontI], true
	}
}

// BackItem returns the item at the back of the deque.
func (d *Deque) BackItem() (int, bool) {
	if d.size <= 0 {
		return 0, false
	} else {
		return d.chunks.Back().Value.(_Chunk)[d.backI], true
	}
}

// Front returns an iterator positioned at the front of the deque, or nil if
// the deque is empty.
func (d *Deque) Front() *Iterator {
	if d.size == 0 {
		return nil
	}
	front := d.chunks.Front()
	return &Iterator{
		Value: front.Value.(_Chunk)[d.frontI],
		deque: d,
		node:  front,
		i:     d.frontI,
		pos:   0,
	}
}

// Back returns an iterator positioned at the back of the deque, or nil if the
// deque is empty.
func (d *Deque) Back() *Iterator {
	if d.size == 0 {
		return nil
	}
	back := d.chunks.Back()
	return &Iterator{
		Value: back.Value.(_Chunk)[d.backI],
		deque: d,
		node:  back,
		i:     d.backI,
		pos:   d.size - 1,
	}
}

// Size returns the number of items in the deque.
func (d *Deque) Size() int {
	return d.size
}

// Clear removes all items from the deque.
func (d *Deque) Clear() {
	*d = *New()
}

func (d *Deque) reset() {
	d.frontI = chunkCenter + 1
	d.backI = chunkCenter
}

//// Iterator //////////////////////////////////////////////////////////////////

// Iterator points to a deque item and can be used to iterate through the deque.
type Iterator struct {
	Value int

	deque *Deque
	node  *list.Element // current chunk
	i     int           // current item index
	pos   int           // iteration position
}

// Next returns an iterator that points to the next deque element, or nil if
// there is no next element.
func (it *Iterator) Next() *Iterator {
	it.pos++
	if it.pos >= it.deque.size { // no more items
		return nil
	}
	it.i++
	if it.i >= chunkSize { // next chunk?
		it.node = it.node.Next()
		it.i = 0
	}
	it.Value = it.node.Value.(_Chunk)[it.i]
	return it
}

// Prev returns an iterator that points to the previous deque element, or nil
// if there is no previous element.
func (it *Iterator) Prev() *Iterator {
	it.pos--
	if it.pos < 0 { // no more items
		return nil
	}
	it.i--
	if it.i < 0 { // previous chunk?
		it.node = it.node.Prev()
		it.i = chunkSize - 1
	}
	it.Value = it.node.Value.(_Chunk)[it.i]
	return it
}

//// _Chunk ////////////////////////////////////////////////////////////////////

type _Chunk []int
