package search

import (
	"log"
	"reflect"
)

type queue struct {
	data []any
}

func NewQueue() queue {
	return queue{}
}

func (s *queue) size() int {
	return len(s.data)
}

func (s *queue) enqueue(newData any) {
	s.data = append(s.data, newData)
}

func (s *queue) dequeue() any {

	head := s.data[0]

	s.data = s.data[1:]

	return head

}

func UseBFS(key string, value string, structure any) bool {

	queue := NewQueue()

	queue.enqueue(structure)

	var v reflect.Value
	var t reflect.Type

	for queue.size() > 0 {

		currObj := queue.dequeue()

		v = reflect.ValueOf(currObj)
		t = reflect.TypeOf(currObj)

		// fmt.Printf("%+v - %+v\n\n", t.Name(), v)
		if t.Kind() != reflect.Struct {
			log.Fatalf("not a struct, but %v", v.Kind())
			return false
		}

		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).Kind() == reflect.Struct {
				obj := v.Field(i).Interface()
				queue.enqueue(obj)
				continue
			}

			currT := t.Field(i).Name
			currV := v.Field(i).Interface()

			if currT == key && currV == value {
				// fmt.Printf("BFS ===> Value is found, used queue %d times\n", count)
				// fmt.Printf("BFS ===> Here is full info %+v\n\n", v)
				return true
			}

		}

	}

	return false
	// fmt.Printf("BFS ===> not found, used queue %d times\n\n", count)

}
