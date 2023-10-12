package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
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

func LlamarRadixSort(tam int) {
	startTime := time.Now()

	var arr []int
	switch tam {
	case 1:
		arr = utilidades.RecuperarArreglo("datos.txt")
		break
	case 2:
		arr = utilidades.RecuperarArreglo("datos2.txt")
		break
	case 3:
		arr = utilidades.RecuperarArreglo("datos3.txt")
		break
	}

	radixSort(arr)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
