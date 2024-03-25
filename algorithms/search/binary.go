package search

import (
	"fmt"
	"slices"
	"time"

	"golang.org/x/exp/constraints"
)

func binaryRecursion(numarray []int, left int, right int, target int) int {

	if left > right {
		return -1
	}

	mid := (left + right) / 2

	time.Sleep(time.Millisecond * 500)

	fmt.Println("array: ", numarray[left:right+1])

	if numarray[mid] == target {
		return mid
	}

	if numarray[mid] > target {

		return binaryRecursion(numarray, left, mid-1, target)
	}

	return binaryRecursion(numarray, mid+1, right, target)

}

func binaryLoop[T constraints.Ordered](slice []T, target T) int {

	left := 0
	right := len(slice)

	for left < right {
		time.Sleep(time.Millisecond * 500)

		fmt.Println("array: ", slice[left:right])

		mid := left + (right-left)/2
		midValue := slice[mid]

		if midValue == target {
			return mid
		} else if midValue > target {
			right = mid

		} else if midValue < target {
			left = mid + 1
		}

	}
	return -1
}

func binaryLoopWithOrderChoice[T constraints.Ordered](slice []T, target T, ascending bool) int {

	left := 0
	right := len(slice)

	for left < right {
		time.Sleep(time.Millisecond * 500)

		fmt.Println("array: ", slice[left:right])

		mid := left + (right-left)/2
		midValue := slice[mid]

		if midValue == target {
			return mid
		} else if midValue > target {
			if ascending {

				right = mid
			} else {
				left = mid
			}

		} else if midValue < target {
			if ascending {

				left = mid + 1
			} else {
				right = mid + 1
			}
		}

	}
	return -1
}

func UseBinarySearch(numarray []int, target int) int {

	left := 0
	right := len(numarray) - 1

	result := binaryRecursion(numarray, left, right, target)

	fmt.Println("result of binary search via recursion: ", result)

	time.Sleep(time.Second * 1)

	result = binaryLoop(numarray, target)

	fmt.Println("result of binary search via loop: ", result)

	slices.Reverse(numarray)

	result = binaryLoopWithOrderChoice(numarray, target, false)

	fmt.Println("result of binary search via loop when array is in descending order: ", result)

	return result
}
