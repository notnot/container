// deque_test.go. jpad 2015

package deque_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/notnot/container/deque"
)

//// tests /////////////////////////////////////////////////////////////////////

func TestEmpty_PopFront(t *testing.T) {
	deque := deque.New()

	for i := 0; i < 3; i++ {
		front := deque.PopFront()
		if front != nil {
			t.Errorf("got: %v, want: <nil>")
		}
	}
}

func TestEmpty_PopBack(t *testing.T) {
	deque := deque.New()

	for i := 0; i < 3; i++ {
		back := deque.PopBack()
		if back != nil {
			t.Errorf("got: %v, want: <nil>")
		}
	}
}

func TestEmpty_iterator(t *testing.T) {
	deque := deque.New()

	front := deque.Front()
	if front.Value != nil {
		t.Errorf("got: %v, want: <nil>", front.Value)
	}

	back := deque.Back()
	if back.Value != nil {
		t.Errorf("got: %v, want: <nil>", back.Value)
	}
}

func TestEmpty_item(t *testing.T) {
	deque := deque.New()

	frontItem := deque.FrontItem()
	if frontItem != nil {
		t.Errorf("got: %v, want: <nil>", frontItem)
	}

	backItem := deque.BackItem()
	if backItem != nil {
		t.Errorf("got: %v, want: <nil>", backItem)
	}
}

func TestPushPeek(t *testing.T) {
	deque := deque.New()

	deque.PushFront("a")
	if deque.FrontItem() != "a" {
		t.Errorf("got: %v, want: a")
	}

	deque.PushBack("z")
	if deque.BackItem() != "z" {
		t.Errorf("got: %v, want: z")
	}
}

func TestPushPop(t *testing.T) {
	deque := deque.New()

	deque.PushFront("a")
	deque.PushBack("z")

	front := deque.PopFront()
	if front != "a" {
		t.Errorf("got: %v, want: a")
	}
	back := deque.PopBack()
	if back != "z" {
		t.Errorf("got: %v, want: z")
	}
}

func TestPushPopRandom(t *testing.T) {
	const N = 1000
	deque := deque.New()

	// randomly push items to the front or to the back
	for i := 0; i < N; i++ {
		switch rand.Intn(2) {
		case 0:
			deque.PushFront("f")
		case 1:
			deque.PushBack("b")
		}
	}

	// randomly pop items from the front or from the back
	for i := 0; i < N; i++ {
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
	deque := deque.New()

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
	deque := deque.New()

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
	deque := deque.New()
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
	fmt.Println()

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
		deque := deque.New()
		for i := 0; i < 10; i++ {
			deque.PushFront(i)
		}
		for i := 0; i < 10; i++ {
			_ = deque.PopFront()
		}
	}
}

func BenchmarkPushPopFront_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque.New()
		for i := 0; i < 100; i++ {
			deque.PushFront(i)
		}
		for i := 0; i < 100; i++ {
			_ = deque.PopFront()
		}
	}
}

func BenchmarkPushPopFront_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque.New()
		for i := 0; i < 1000; i++ {
			deque.PushFront(i)
		}
		for i := 0; i < 1000; i++ {
			_ = deque.PopFront()
		}
	}
}

func BenchmarkPushPopBack_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque.New()
		for i := 0; i < 10; i++ {
			deque.PushBack(i)
		}
		for i := 0; i < 10; i++ {
			_ = deque.PopBack()
		}
	}
}

func BenchmarkPushPopBack_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque.New()
		for i := 0; i < 100; i++ {
			deque.PushBack(i)
		}
		for i := 0; i < 100; i++ {
			_ = deque.PopBack()
		}
	}
}

func BenchmarkPushPopBack_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deque := deque.New()
		for i := 0; i < 1000; i++ {
			deque.PushBack(i)
		}
		for i := 0; i < 1000; i++ {
			_ = deque.PopBack()
		}
	}
}

func BenchmarkIterate_forward(b *testing.B) {
	const N = 1024
	deque := deque.New()
	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// iterate from front to back
		for it := deque.Front(); it != nil; it = it.Next() {
			_ = it.Value
		}
	}
}

func BenchmarkIterate_backward(b *testing.B) {
	const N = 1024
	deque := deque.New()
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
	deque := deque.New()
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
