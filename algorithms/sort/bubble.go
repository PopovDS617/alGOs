package sort

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

func swap[T constraints.Ordered](array []T, a int, b int) {
	array[a], array[b] = array[b], array[a]
}

func bubble[T constraints.Ordered](array []T) []T {
	for i := 0; i < len(array)-1; i++ {

		for j := 0; j < len(array)-i-1; j++ {

			time.Sleep(time.Millisecond * 150)
			if array[j] > array[j+1] {
				swap(array, j, j+1)
			}
			fmt.Println("inner loop: ", array)
		}
		fmt.Println("outer loop: ", array)
	}
	return array
}

func bubbleWithOrder[T constraints.Ordered](array []T, ascending bool) []T {
	for i := 0; i < len(array)-1; i++ {

		for j := 0; j < len(array)-i-1; j++ {

			time.Sleep(time.Millisecond * 150)

			if ascending {
				if array[j] > array[j+1] {
					swap(array, j, j+1)
				}
			} else {
				if array[j] < array[j+1] {
					swap(array, j+1, j)
				}
			}

		}
		fmt.Println("inner loop: ", array)
	}
	fmt.Println("outer loop: ", array)

	return array
}

func UseBubbleSort(numarray []int) {

	// result := bubble(numarray)

	result := bubbleWithOrder(numarray, true)
	fmt.Println("ASCENDING RESULT ", result)

	result = bubbleWithOrder(numarray, false)
	fmt.Println("DESCENDING RESULT ", result)

}
