package sort

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

func quick[T constraints.Ordered](slice []T, low int, high int) []T {
	if low < high {
		time.Sleep(time.Millisecond * 250)
		fmt.Println(slice)

		var p int
		slice, p = partition(slice, low, high)
		slice = quick(slice, low, p-1)
		slice = quick(slice, p+1, high)
	}

	return slice
}

func partition[T constraints.Ordered](arr []T, low, high int) ([]T, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {

		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func UseQuickSort[T constraints.Ordered](slice []T) {
	fmt.Println("ORIGINAL ", slice)
	result := quick(slice, 0, len(slice)-1)

	fmt.Println("RESULT ", result)

}
