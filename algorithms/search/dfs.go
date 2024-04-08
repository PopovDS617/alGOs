package search

import (
	"fmt"
	"log"
	"reflect"
)

type stack struct {
	data []any
}

func NewStack() stack {
	return stack{}
}

func (s *stack) size() int {
	return len(s.data)
}

func (s *stack) push(newData any) {
	s.data = append(s.data, newData)
}

func (s *stack) pop() any {

	last := s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return last

}

type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Position  string
}

type Office struct {
	Head       Employee
	HR         Employee
	Accountant Employee
}

type Company struct {
	NewYork    Office
	LosAngeles Office
	Washington Office
}

var ExampleCompany = Company{
	NewYork: Office{
		Head:       Employee{FirstName: "John", LastName: "Johnson", Age: 32, Position: "Head of the New York office"},
		HR:         Employee{FirstName: "Max", LastName: "Paxton", Age: 33, Position: "HR of the New York office"},
		Accountant: Employee{FirstName: "Richard", LastName: "Johansson", Age: 34, Position: "Accountant of the New York office"},
	},
	LosAngeles: Office{
		Head:       Employee{FirstName: "Sarah", LastName: "Brodson", Age: 22, Position: "Head of the Los Angeles office"},
		HR:         Employee{FirstName: "Louise", LastName: "Cooper", Age: 23, Position: "HR of the Los Angeles office"},
		Accountant: Employee{FirstName: "Anna", LastName: "Gregson", Age: 24, Position: "Accountant of the Los Angeles office"}},
	Washington: Office{
		Head:       Employee{FirstName: "Patrick", LastName: "Hewett", Age: 42, Position: "Head of the Washington office"},
		HR:         Employee{FirstName: "Richard", LastName: "Langley", Age: 43, Position: "HR of the Washington office"},
		Accountant: Employee{FirstName: "Donovan", LastName: "Svergsson", Age: 44, Position: "Accountant of the Washington office"}},
}

func UseDFS() {

	count := 0

	stack := NewStack()

	stack.push(ExampleCompany)

	for stack.size() > 0 {
		count++

		currObj := stack.pop()

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
				stack.push(obj)
				continue
			}

			currT := t.Field(i).Name
			currV := v.Field(i).Interface()

			if currT == "LastName" && currV == "Paxton" {
				fmt.Printf("DFS ==> Mr. Paxton is found, used stack %d times\n", count)
				fmt.Printf("DFS ==> Here is full info on him %+v\n\n", v)
				return
			}

		}

	}

	fmt.Printf("DFS ==> not found, used stack %d times\n\n", count)

}
