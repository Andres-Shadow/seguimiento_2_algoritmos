package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
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

func LlamarGnomeSort(tam int) {
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

	gnomeSort(arr, len(arr))

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
