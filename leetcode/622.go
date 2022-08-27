package leetcode

type MyCircularQueue struct {
	list []int
	head int
	tail int
}

func Constructor(k int) MyCircularQueue {
	queue := MyCircularQueue{
		list: make([]int, k),
	}
	for i := 0; i < len(queue.list); i += 1 {
		queue.list[i] = -1
	}
	return queue
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if cap(this.list) == 1 {
		if this.list[0] != -1 {
			return false
		}
		this.list[0] = value
		return true
	}
	if (this.tail+1)%cap(this.list) == this.head {
		return false
	}

	if (this.tail == this.head && this.list[this.tail] != -1) || this.tail != this.head {
		// there is one item in the queue
		this.tail = this.incr(this.tail)
	}
	this.list[this.tail] = value
	return true
}

func (this *MyCircularQueue) incr(i int) int {
	i += 1
	if i >= cap(this.list) {
		i -= cap(this.list)
	}
	return i
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.tail == this.head {
		if this.list[this.head] != -1 {
			this.list[this.head] = -1
			return true
		}
		return false
	}
	this.list[this.head] = -1
	this.head = this.incr(this.head)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.head == this.tail {
		if this.list[this.head] == -1 {
			return -1
		}
		return this.list[this.head]
	}
	return this.list[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.head == this.tail && this.list[this.tail] == -1 {
		return -1
	}
	return this.list[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.tail == this.head && this.list[this.head] == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return this.incr(this.tail) == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
