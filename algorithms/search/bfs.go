package search

import (
	"fmt"
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

type employee struct {
	FirstName string
	LastName  string
	Age       int
	Position  string
}

type office struct {
	Head       employee
	HR         employee
	Accountant employee
}

type company struct {
	NewYork    office
	LosAngeles office
	Washington office
}

var exampleCompany = company{
	NewYork: office{
		Head:       employee{FirstName: "John", LastName: "Johnson", Age: 32, Position: "Head of the New York office"},
		HR:         employee{FirstName: "Max", LastName: "Paxton", Age: 33, Position: "HR of the New York office"},
		Accountant: employee{FirstName: "Richard", LastName: "Johansson", Age: 34, Position: "Accountant of the New York office"},
	},
	LosAngeles: office{
		Head:       employee{FirstName: "Sarah", LastName: "Brodson", Age: 22, Position: "Head of the Los Angeles office"},
		HR:         employee{FirstName: "Louise", LastName: "Cooper", Age: 23, Position: "HR of the Los Angeles office"},
		Accountant: employee{FirstName: "Anna", LastName: "Gregson", Age: 24, Position: "Accountant of the Los Angeles office"}},
	Washington: office{
		Head:       employee{FirstName: "Patrick", LastName: "Hewett", Age: 42, Position: "Head of the Washington office"},
		HR:         employee{FirstName: "Richard", LastName: "Langley", Age: 43, Position: "HR of the Washington office"},
		Accountant: employee{FirstName: "Donovan", LastName: "Svergsson", Age: 44, Position: "Accountant of the Washington office"}},
}

func UseBFS() {

	count := 0

	queue := NewQueue()

	queue.enqueue(exampleCompany)

	for queue.size() > 0 {
		count++

		currObj := queue.dequeue()

		v := reflect.ValueOf(currObj)
		t := reflect.TypeOf(currObj)

		// fmt.Printf("%+v - %+v\n\n", t.Name(), v)
		if t.Kind() != reflect.Struct {
			log.Fatalf("not a struct, but %v", v.Kind())
			return
		}

		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).Kind() == reflect.Struct {
				obj := v.Field(i).Interface()
				queue.enqueue(obj)
				continue
			}

			currT := t.Field(i).Name
			currV := v.Field(i).Interface()

			if currT == "LastName" && currV == "Langley" {
				fmt.Printf("BFS ===> Mr. Langley is found, used queue %d times\n", count)
				fmt.Printf("BFS ===> Here is full info on him %+v\n\n", v)
				return
			}

		}

	}

	fmt.Printf("BFS ===> not found, used queue %d times\n\n", count)

}
