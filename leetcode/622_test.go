package leetcode

import (
	"fmt"
	"testing"
)

func TestMyCircularQueue_EnQueue(t *testing.T) {
	q := Constructor(3)
	var succ bool
	succ = q.EnQueue(1)
	if !succ {
		t.Errorf("enqueue false before full")
	}
	if q.head != 0 {
		t.Errorf("wrong head after insert one element")
	}
	if q.tail != 0 {
		t.Errorf("wrong tail after insert one element")
	}
	succ = q.EnQueue(2)
	if !succ {
		t.Errorf("enqueue false before full")
	}
	if q.head != 0 {
		t.Errorf("wrong head after insert one element")
	}
	if q.tail != 1 {
		t.Errorf("wrong tail after insert one element")
	}
	succ = q.EnQueue(3)
	if !succ {
		t.Errorf("enqueue false before full")
	}
	if q.head != 0 {
		t.Errorf("wrong head after insert one element")
	}
	if q.tail != 2 {
		t.Errorf("wrong tail after insert one element")
	}
	succ = q.EnQueue(4)
	if succ {
		t.Errorf("enqueue true but full")
	}

}

func TestMyCircularQueue_DeQueue(t *testing.T) {
	q := Constructor(3)
	var succ bool
	succ = q.DeQueue()
	if succ {
		t.Errorf("wrong element dequeued")
	}

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 1 {
		t.Errorf("wrong head after dequeue one element")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 2 {
		t.Errorf("wrong head after dequeue one element")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 2 {
		t.Errorf("wrong head after dequeue one element")
	}
}

func TestMyCircularQueue_DeQueue2(t *testing.T) {
	q := Constructor(3)
	var succ bool
	succ = q.DeQueue()
	if succ {
		t.Errorf("wrong element dequeued")
	}

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 1 {
		t.Errorf("wrong head after dequeue one element")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 2 {
		t.Errorf("wrong head after dequeue one element")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 2 {
		t.Errorf("wrong head after dequeue one element")
	}

	succ = q.EnQueue(4)
	if q.head != 2 {
		t.Errorf("wrong head")
	}
	if q.tail != 2 {
		t.Errorf("wrong head")
	}

	succ = q.EnQueue(5)
	if q.head != 2 {
		t.Errorf("wrong head")
	}
	if q.tail != 0 {
		t.Errorf("wrong head")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 0 {
		t.Errorf("wrong head")
	}
	if q.tail != 0 {
		t.Errorf("wrong head")
	}
}

func TestMyCircularQueue_Comprehensive(t *testing.T) {
	q := Constructor(3)

	if !q.IsEmpty() {
		t.Errorf("is empty")
	}
	var succ bool

	succ = q.DeQueue()
	if succ {
		t.Errorf("wrong element dequeued")
	}

	if q.Rear() != -1 {
		t.Errorf("wrong rear")
	}

	q.EnQueue(1)

	if q.IsEmpty() {
		t.Errorf("not empty")
	}

	if q.Rear() != 1 {
		t.Errorf("wrong rear")
	}

	q.EnQueue(2)
	q.EnQueue(3)

	if q.Front() != 1 {
		t.Errorf("wrong front")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 1 {
		t.Errorf("wrong head after dequeue one element")
	}

	if q.Front() != 2 {
		t.Errorf("wrong front")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 2 {
		t.Errorf("wrong head after dequeue one element")
	}

	if q.Front() != 3 {
		t.Errorf("wrong front")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 2 {
		t.Errorf("wrong head after dequeue one element")
	}

	if !q.IsEmpty() {
		t.Errorf("is empty")
	}

	if q.IsFull() {
		t.Errorf("not full")
	}

	succ = q.EnQueue(4)
	if q.head != 2 {
		t.Errorf("wrong head")
	}
	if q.tail != 2 {
		t.Errorf("wrong head")
	}

	succ = q.EnQueue(5)
	if q.head != 2 {
		t.Errorf("wrong head")
	}
	if q.tail != 0 {
		t.Errorf("wrong head")
	}

	succ = q.DeQueue()
	if !succ {
		t.Errorf("dequeue fail")
	}
	if q.head != 0 {
		t.Errorf("wrong head")
	}
	if q.tail != 0 {
		t.Errorf("wrong head")
	}
}

func TestMyCircularQueue_Integrated(t *testing.T) {
	q := Constructor(8)
	var res bool
	res = q.EnQueue(3)
	fmt.Println(res)
	res = q.EnQueue(9)
	fmt.Println(res)

	res = q.EnQueue(5)
	fmt.Println(res)

	res = q.EnQueue(0)
	fmt.Println(res)

	q.DeQueue()
	q.DeQueue()
	if q.Rear() != 5 {
		t.Errorf("wrong rear")
	}
}

func TestMyCircularQueue_Integrated2(t *testing.T) {
	q := Constructor(1)
	var res bool
	res = q.EnQueue(1)
	fmt.Println(res)
	res = q.EnQueue(2)
	fmt.Println(res)

}
