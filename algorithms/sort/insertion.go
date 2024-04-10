package sort

import (
	"fmt"
	"time"
)

func UseInsertionSort(arr []int, asc bool) {

	len := len(arr)

	for i := 1; i < len; i++ {
		// time.Sleep(time.Second * 1)
		// fmt.Println("outer", arr)

		for j := 0; j < i; j++ {
			fmt.Println("i", i, "j", j)

			fmt.Println("j at ", arr[j], "i at:", arr[i])
			if asc {
				if arr[j] > arr[i] {
					arr[j], arr[i] = arr[i], arr[j]
					fmt.Println("swapping ", arr[j], "and", arr[i])
				}

			} else {
				if arr[j] < arr[i] {
					arr[j], arr[i] = arr[i], arr[j]
				}
			}
			time.Sleep(time.Second * 1)
			fmt.Println("inner", arr)
			fmt.Println("inner end")
		}
	}

	fmt.Println(arr)
}

// 1024, 345, 44, 2, 25, 44, 1, 3, 101, 94
