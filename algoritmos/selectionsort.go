package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

// Función para ordenar arr[] usando Selection Sort
func selectionSort(arr []int) {
	n := len(arr)

	// Mover el límite de la submatriz no ordenada uno por uno
	for i := 0; i < n-1; i++ {
		// Encontrar el elemento mínimo en la submatriz no ordenada
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		// Intercambiar el elemento mínimo encontrado con el primer elemento
		temp := arr[minIdx]
		arr[minIdx] = arr[i]
		arr[i] = temp
	}
}

func LlamarSelectionSort() {
	startTime := time.Now()
	arr := utilidades.RecuperarArreglo()
	selectionSort(arr)
	utilidades.ImprimirArreglo(arr)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)

}
