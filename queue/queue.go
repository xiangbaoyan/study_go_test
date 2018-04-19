package queue

type Queue []int

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//Push Element into the queue
func (q *Queue) Push(a int) {
	*q = append(*q, a)
}

func (q *Queue) IsEmpty() bool {
	return 0 == len(*q)
}
