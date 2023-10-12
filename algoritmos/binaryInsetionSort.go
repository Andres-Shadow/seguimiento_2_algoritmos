package algoritmos

import (
	"fmt"
	"time"
)

func binarySearch(arr []int, low, high, x int) int {
	if high <= low {
		if x > arr[low] {
			return low + 1
		}
		return low
	}

	mid := (low + high) / 2

	if x == arr[mid] {
		return mid + 1
	} else if x > arr[mid] {
		return binarySearch(arr, mid+1, high, x)
	}
	return binarySearch(arr, low, mid-1, x)
}

func binaryInsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := binarySearch(arr, 0, i-1, x)
		copy(arr[j+1:i+1], arr[j:i])
		arr[j] = x
	}
}

func LlamarBinaryInsertionSort() {

	startTime := time.Now()

	//arr := utilidades.RecuperarArreglo()
	arr := []int{37, 23, 0, 17, 12, 72, 31, 46, 100, 88, 54, 32}
	binaryInsertionSort(arr)

	fmt.Println("Sorted array:")
	fmt.Println(arr)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
