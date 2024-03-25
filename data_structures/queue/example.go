package queue

import (
	"bytes"
	"fmt"
)

type queue struct {
	data []string
}

func (q *queue) enqueue(element string) {
	q.data = append(q.data, element)
	fmt.Println("Enqueued:", element)

}

func (q *queue) dequeue() string {
	if len(q.data) == 1 {
		element := q.data[0]
		q.data = []string{}

		return element
	} else if len(q.data) != 0 {
		element := q.data[0]
		q.data = q.data[1:]

		return element
	}

	return ""
}

func (q *queue) print() {
	var b bytes.Buffer

	for _, v := range q.data {
		b.WriteString(v + " ")
	}

	fmt.Println("===> Queue [ ", b.String(), "]")
}

func newQueue() *queue {
	return &queue{
		data: []string{},
	}
}

func UseQueue() {

	var text = []string{"hello", "world", "from", "the", "computer"}
	q := newQueue()

	for _, v := range text {
		q.enqueue(v)
	}

	q.dequeue()
	q.dequeue()

	q.print()

}
