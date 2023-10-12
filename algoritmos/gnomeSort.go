package algoritmos

import (
	"fmt"
	"time"
)

func gnomeSort(arr []int, n int) {
	index := 0

	for index < n {
		if index == 0 {
			index++
		}
		if arr[index] >= arr[index-1] {
			index++
		} else {
			// Intercambiar elementos
			arr[index], arr[index-1] = arr[index-1], arr[index]
			index--
		}
	}
}

func LlamarGnomeSort() {
	startTime := time.Now()

	//arr := utilidades.RecuperarArreglo()
	arr := []int{34, 2, 10, 0, 55, 7}

	gnomeSort(arr, len(arr))

	fmt.Print("array ordenado: ")
	fmt.Println(arr)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
