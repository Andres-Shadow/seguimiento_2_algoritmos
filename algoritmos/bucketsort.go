package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

func bucketSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Encontrar el máximo y mínimo valores en el arreglo para determinar el rango.
	minVal, maxVal := arr[0], arr[0]
	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}

	// Crear buckets (cubos) basados en el rango de valores.
	bucketSize := (maxVal-minVal)/len(arr) + 1
	numBuckets := (maxVal-minVal)/bucketSize + 1
	buckets := make([][]int, numBuckets)

	// Colocar elementos en los buckets correspondientes.
	for _, val := range arr {
		index := (val - minVal) / bucketSize
		buckets[index] = append(buckets[index], val)
	}

	// Ordenar cada bucket individualmente (puedes usar otro algoritmo de ordenamiento aquí).
	for i := 0; i < numBuckets; i++ {
		insertion(buckets[i])
	}

	// Concatenar los buckets ordenados para obtener el arreglo ordenado.
	index := 0
	for i := 0; i < numBuckets; i++ {
		for _, val := range buckets[i] {
			arr[index] = val
			index++
		}
	}
}

func insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func LlamarBucketSort(tam int) {
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

	bucketSort(arr)
	//utilidades.ImprimirArreglo(arr)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
