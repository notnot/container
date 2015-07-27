// deque_test.go. jpad 2015

package deque_int_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/notnot/container/deque_int"
)

//// tests /////////////////////////////////////////////////////////////////////

func TestEmpty_PopFront(t *testing.T) {
	deque := deque_int.New()

	for i := 0; i < 3; i++ {
		front := deque.PopFront()
		if front != 0 {
			t.Errorf("got: %v, want: 0", front)
		}
	}
}

func TestEmpty_PopBack(t *testing.T) {
	deque := deque_int.New()

	for i := 0; i < 3; i++ {
		back := deque.PopBack()
		if back != 0 {
			t.Errorf("got: %v, want: 0", back)
		}
	}
}

func TestEmpty_iterator(t *testing.T) {
	deque := deque_int.New()

	front := deque.Front()
	if front != nil {
		t.Errorf("got: %v, want: <nil>", front)
	}

	back := deque.Back()
	if back != nil {
		t.Errorf("got: %v, want: <nil>", back)
	}
}

func TestEmpty_item(t *testing.T) {
	deque := deque_int.New()

	front := deque.FrontItem()
	if front != 0 {
		t.Errorf("got: %d, want: 0", front)
	}

	back := deque.BackItem()
	if back != 0 {
		t.Errorf("got: %d, want: 0", back)
	}
}

func TestPushPeek(t *testing.T) {
	deque := deque_int.New()

	deque.PushFront(0)
	if item := deque.FrontItem(); item != 0 {
		t.Errorf("got: %d, want: 0", item)
	}

	deque.PushBack(1)
	if item := deque.BackItem(); item != 1 {
		t.Errorf("got: %d, want: 1", item)
	}
}

func TestPushPop(t *testing.T) {
	deque := deque_int.New()

	deque.PushFront(0)
	deque.PushBack(1)

	front := deque.PopFront()
	if front != 0 {
		t.Errorf("got: %d, want: 0", front)
	}
	back := deque.PopBack()
	if back != 1 {
		t.Errorf("got: %d, want: 1", back)
	}
}

func TestPushPopRandom(t *testing.T) {
	const N = 1000
	deque := deque_int.New()

	// randomly push items to the front or to the back
	for i := 0; i < N; i++ {
		switch rand.Intn(2) {
		case 0:
			deque.PushFront(0)
		case 1:
			deque.PushBack(1)
		}
	}

	// randomly pop items from the front or from the back
	for i := 0; i < N; i++ {
		if deque.Size() <= 0 {
			t.Errorf("deque empty!")
		}
		switch rand.Intn(2) {
		case 0:
			_ = deque.PopFront()
		case 1:
			_ = deque.PopBack()
		}
	}
}

func TestSize(t *testing.T) {
	const N = 10
	deque := deque_int.New()

	if deque.Size() != 0 {
		t.Errorf("got: %d, want: 0", deque.Size())
	}

	for i := 0; i < N; i++ {
		deque.PushFront(i)
	}
	if deque.Size() != N {
		t.Errorf("got: %d, want: %d", deque.Size(), N)
	}
	for deque.Size() > 0 {
		deque.PopFront()
	}
	if deque.Size() != 0 {
		t.Errorf("got: %d, want: 0", deque.Size())
	}

	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	if deque.Size() != N {
		t.Errorf("got: %d, want: %d", deque.Size(), N)
	}
	for deque.Size() > 0 {
		deque.PopBack()
	}
	if deque.Size() != 0 {
		t.Errorf("got: %d, want: 0", deque.Size())
	}
}

func TestClear(t *testing.T) {
	const N = 10
	deque := deque_int.New()

	for i := 0; i < N; i++ {
		deque.PushFront(i)
	}
	deque.Clear()
	if deque.Size() != 0 {
		t.Errorf("got: %d, want: 0", deque.Size())
	}
}

func TestIterate(t *testing.T) {
	const N = 1000
	deque := deque_int.New()
	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}

	// iterate from front to back
	i := 0
	for e := deque.Front(); e != nil; e = e.Next() {
		if e.Value != i {
			t.Errorf("got: %d, want: %d", e.Value, i)
		}
		i++
	}

	// iterate from back to front
	i = N - 1
	for e := deque.Back(); e != nil; e = e.Prev() {
		if e.Value != i {
			t.Errorf("got: %d, want: %d", e.Value, i)
		}
		i--
	}
}

//// benchmarks ////////////////////////////////////////////////////////////////

func BenchmarkPushPopFront_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque_int.New()
		for i := 0; i < 10; i++ {
			deque.PushFront(i)
		}

		sum := 0
		for i := 0; i < 10; i++ {
			sum += deque.PopFront()
		}
	}
}

func BenchmarkPushPopFront_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque_int.New()
		for i := 0; i < 100; i++ {
			deque.PushFront(i)
		}

		sum := 0
		for i := 0; i < 100; i++ {
			sum += deque.PopFront()
		}
	}
}

func BenchmarkPushPopFront_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque_int.New()
		for i := 0; i < 1000; i++ {
			deque.PushFront(i)
		}

		sum := 0
		for i := 0; i < 1000; i++ {
			sum += deque.PopFront()
		}
	}
}

func BenchmarkPushPopBack_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque_int.New()
		for i := 0; i < 10; i++ {
			deque.PushBack(i)
		}

		sum := 0
		for i := 0; i < 10; i++ {
			sum += deque.PopBack()
		}
	}
}

func BenchmarkPushPopBack_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque_int.New()
		for i := 0; i < 100; i++ {
			deque.PushBack(i)
		}

		sum := 0
		for i := 0; i < 100; i++ {
			sum += deque.PopBack()
		}
	}
}

func BenchmarkPushPopBack_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque_int.New()
		for i := 0; i < 1000; i++ {
			deque.PushBack(i)
		}

		sum := 0
		for i := 0; i < 1000; i++ {
			sum += deque.PopBack()
		}
	}
}

func BenchmarkFrontItem(b *testing.B) {
	const N = 16
	deque := deque_int.New()
	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = deque.FrontItem()
	}
}

func BenchmarkBackItem(b *testing.B) {
	const N = 16
	deque := deque_int.New()
	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = deque.BackItem()
	}
}

func BenchmarkIterate_forward(b *testing.B) {
	const N = 1024
	deque := deque_int.New()
	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for it := deque.Front(); it != nil; it = it.Next() {
			_ = it.Value
		}
	}
}

func BenchmarkIterate_backward(b *testing.B) {
	const N = 1024
	deque := deque_int.New()
	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for it := deque.Back(); it != nil; it = it.Prev() {
			_ = it.Value
		}
	}
}

//// examples //////////////////////////////////////////////////////////////////

func ExampleIterator() {
	const N = 10
	deque := deque_int.New()
	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}

	// iterate from front to back
	for it := deque.Front(); it != nil; it = it.Next() {
		fmt.Printf("%v", it.Value)
	}
	fmt.Println()

	// iterate from back to front
	for it := deque.Back(); it != nil; it = it.Prev() {
		fmt.Printf("%v", it.Value)
	}
	fmt.Println()

	// Output:
	// 0123456789
	// 9876543210
}
