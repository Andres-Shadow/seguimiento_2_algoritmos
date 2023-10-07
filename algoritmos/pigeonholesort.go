package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

func pigeonholeSort(arr []int) {
	min := arr[0]
	max := arr[0]
	n := len(arr)

	// Encontrar el valor mínimo y máximo en el arreglo
	for a := 0; a < n; a++ {
		if arr[a] > max {
			max = arr[a]
		}
		if arr[a] < min {
			min = arr[a]
		}
	}

	// Crear un mapa para almacenar la frecuencia de los elementos
	frequencyMap := make(map[int]int)

	// Calcular las frecuencias
	for _, num := range arr {
		frequencyMap[num]++
	}

	// Reconstruir el arreglo ordenado
	index := 0
	for j := min; j <= max; j++ {
		for frequencyMap[j] > 0 {
			arr[index] = j
			index++
			frequencyMap[j]--
		}
	}
}

func LlamarPigeonHoleSort() {
	startTime := time.Now()
	arr := utilidades.RecuperarArreglo()
	pigeonholeSort(arr)
	utilidades.ImprimirArreglo(arr)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
