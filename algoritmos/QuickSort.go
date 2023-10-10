package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

// Esta función se utiliza para intercambiar elementos en el arreglo durante el proceso de particionamiento.
func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// utiliza como pivote el high y pone los mas pequeños a la izquierda y los mas grandes a la derecha
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i + 1
}

func quickSort(arr []int, low, high int) {
	//si aun hay elementos sin ordenar
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func LlamarQuickSort(tam int) {

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
	N := len(arr)

	quickSort(arr, 0, N-1)

	fmt.Println("Sorted array:")
	utilidades.ImprimirArreglo(arr)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
