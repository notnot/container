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
	chunkSize   = 32 // benchmarked optimum on a 64-bit machine
	chunkCenter = chunkSize / 2
)

//// Deque /////////////////////////////////////////////////////////////////////

// Deque is a double ended queue that can handle items of any type.
type Deque struct {
	chunks list.List
	fC     _Chunk // front chunk (shortcut)
	bC     _Chunk // back chunk (shortcut)
	fI     int    // front item index
	bI     int    // back item index
	size   int
}

// New returns a pointer to an empty deque.
func New() *Deque {
	deque := Deque{}
	deque.reset()
	chunk := make(_Chunk, chunkSize)
	deque.fC = chunk
	deque.bC = chunk
	deque.chunks.PushBack(chunk)
	return &deque
}

// PushFront adds an item to the front of the deque.
func (d *Deque) PushFront(item int) {
	if d.fI == 0 { // 'front' chunk full?
		// add a new chunk at the front
		d.fC = make(_Chunk, chunkSize)
		d.chunks.PushFront(d.fC)
		d.fI = chunkSize
	}
	d.fI--
	d.fC[d.fI] = item
	d.size++
}

// PushBack adds an item to the back of the deque.
func (d *Deque) PushBack(item int) {
	if d.bI == chunkSize-1 { // 'back' chunk full?
		// add a new chunk at the back
		d.bC = make(_Chunk, chunkSize)
		d.chunks.PushBack(d.bC)
		d.bI = -1
	}
	d.bI++
	d.bC[d.bI] = item
	d.size++
}

// PopFront removes and returns the item from the front of the deque.
// Returns the zero value when the deque is empty.
func (d *Deque) PopFront() int {
	if d.size <= 0 {
		return 0
	}
	item := d.fC[d.fI]
	d.fI++
	d.size--

	if d.fI == chunkSize { // 'front' chunk empty?
		if d.size == 0 { // deque is empty, reset it
			d.reset()
		} else {
			d.chunks.Remove(d.chunks.Front())
			d.fI = 0
			d.fC = d.chunks.Front().Value.(_Chunk)
		}
	}

	return item
}

// PopBack removes and returns the item from the back of the deque.
// Returns the zero value when the deque is empty.
func (d *Deque) PopBack() int {
	if d.size <= 0 {
		return 0
	}
	item := d.bC[d.bI]
	d.bI--
	d.size--

	if d.bI == -1 { // 'back' chunk empty?
		if d.size == 0 { // deque is empty, reset it
			d.reset()
		} else {
			d.chunks.Remove(d.chunks.Back())
			d.bI = chunkSize - 1
			d.bC = d.chunks.Back().Value.(_Chunk)
		}
	}

	return item
}

// FrontItem returns the item at the front of the deque.
// Returns the zero value when the deque is empty.
func (d *Deque) FrontItem() int {
	if d.size <= 0 {
		return 0
	} else {
		return d.fC[d.fI]
	}
}

// BackItem returns the item at the back of the deque.
// Returns the zero value when the deque is empty.
func (d *Deque) BackItem() int {
	if d.size <= 0 {
		return 0
	} else {
		return d.bC[d.bI]
	}
}

// Front returns an iterator positioned at the front of the deque, or nil if
// the deque is empty.
func (d *Deque) Front() *Iterator {
	if d.size == 0 {
		return nil
	}
	fNode := d.chunks.Front()
	return &Iterator{
		Value: d.fC[d.fI],
		deque: d,
		node:  fNode,
		chunk: fNode.Value.(_Chunk),
		i:     d.fI,
		pos:   0,
	}
}

// Back returns an iterator positioned at the back of the deque, or nil if
// the deque is empty.
func (d *Deque) Back() *Iterator {
	if d.size == 0 {
		return nil
	}
	bNode := d.chunks.Back()
	return &Iterator{
		Value: d.bC[d.bI],
		deque: d,
		node:  bNode,
		chunk: bNode.Value.(_Chunk),
		i:     d.bI,
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
	d.fI = chunkCenter + 1
	d.bI = chunkCenter
}

//// Iterator //////////////////////////////////////////////////////////////////

// Iterator points to a deque item and can be used to iterate through the deque.
type Iterator struct {
	Value int

	deque *Deque
	node  *list.Element // current chunk node
	chunk _Chunk        // current chunk (shortcut)
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
		it.chunk = it.node.Value.(_Chunk)
		it.i = 0
	}
	it.Value = it.chunk[it.i]
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
		it.chunk = it.node.Value.(_Chunk)
		it.i = chunkSize - 1
	}
	it.Value = it.chunk[it.i]
	return it
}

//// _Chunk ////////////////////////////////////////////////////////////////////

type _Chunk []int
