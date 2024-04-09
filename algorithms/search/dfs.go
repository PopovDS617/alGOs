package search

import (
	"log"
	"reflect"
)

type Node struct {
	Data any
	Prev *Node
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) add(value any) {

	newNode := &Node{Data: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode

}

func (l *List) remove(value any) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == value {
		l.Head = l.Head.Next
		return
	}

	curr := l.Head
	for curr.Next != nil && curr.Next.Data != value {
		curr = curr.Next
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func makeList(input []any) *List {
	list := &List{}

	for _, v := range input {
		list.add(v)
	}

	return list
}

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
	Address    string
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
		Address:    "Test avenue 1",
		Head:       Employee{FirstName: "John", LastName: "Johnson", Age: 32, Position: "Head of the New York office"},
		HR:         Employee{FirstName: "Max", LastName: "Paxton", Age: 33, Position: "HR of the New York office"},
		Accountant: Employee{FirstName: "Richard", LastName: "Johansson", Age: 34, Position: "Accountant of the New York office"},
	},
	LosAngeles: Office{
		Address:    "Test avenue 2",
		Head:       Employee{FirstName: "Sarah", LastName: "Brodson", Age: 22, Position: "Head of the Los Angeles office"},
		HR:         Employee{FirstName: "Louise", LastName: "Cooper", Age: 23, Position: "HR of the Los Angeles office"},
		Accountant: Employee{FirstName: "Anna", LastName: "Gregson", Age: 24, Position: "Accountant of the Los Angeles office"}},
	Washington: Office{
		Address:    "Test avenue 3",
		Head:       Employee{FirstName: "Patrick", LastName: "Hewett", Age: 42, Position: "Head of the Washington office"},
		HR:         Employee{FirstName: "Richard", LastName: "Langley", Age: 43, Position: "HR of the Washington office"},
		Accountant: Employee{FirstName: "Donovan", LastName: "Svergsson", Age: 44, Position: "Accountant of the Washington office"}},
}

func UseDFS(key string, value string, structure any) bool {
	stack := NewStack()
	stack.push(structure)

	var v reflect.Value
	var t reflect.Type

	for stack.size() > 0 {
		currObj := stack.pop()

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
				stack.push(obj)
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
	// fmt.Printf("DFS ==> not found, used stack %d times\n\n", count)
}
func UseDFSRecursively(key string, value string, currObj any) bool {

	v := reflect.ValueOf(currObj)
	t := reflect.TypeOf(currObj)
	// fmt.Printf("%+v - %+v\n\n", t.Name(), v)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct {
			obj := v.Field(i).Interface()
			res := UseDFSRecursively(key, value, obj)

			if res {
				return true
			}
		}

		currT := t.Field(i).Name
		currV := v.Field(i).Interface()

		if currT == key && currV == value {
			// fmt.Println("BFS ===> Value is found")
			// fmt.Printf("BFS ===> Here is full info %+v\n\n", v)
			return true
		}

	}

	return false

}

func UseDFSWithLL(key string, value string, structure any) bool {
	arrayed := []any{structure}
	list := makeList(arrayed)

	currNode := list.Head

	var v reflect.Value
	var t reflect.Type

	for currNode.Data != nil {

		v = reflect.ValueOf(currNode.Data)
		t = reflect.TypeOf(currNode.Data)

		if t.Kind() != reflect.Struct {
			log.Fatalf("not a struct, but %v", v.Kind())
			return false
		}

		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).Kind() == reflect.Struct {

				obj := v.Field(i).Interface()
				list.add(obj)
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
		if currNode.Next == nil {
			return false
		} else {
			currNode = currNode.Next
		}
	}

	return false
	// fmt.Printf("DFS ==> not found, used stack %d times\n\n", count)
}
