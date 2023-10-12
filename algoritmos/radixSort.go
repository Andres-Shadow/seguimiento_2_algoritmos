package algoritmos

import (
	"fmt"
	"time"
)

func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func countSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func radixSort(arr []int) {
	max := getMax(arr)
	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}
}

func printAr(arr []int) {
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func LlamarRadixSort() {
	startTime := time.Now()

	//arr := utilidades.RecuperarArreglo()
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66, 22, 1, 55}
	radixSort(arr)
	printAr(arr)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
