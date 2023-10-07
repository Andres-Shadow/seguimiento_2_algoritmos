package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

func bucketSort(arr []float64) {
	n := len(arr)
	if n <= 0 {
		return
	}

	// 1) Crear n buckets vacíos
	buckets := make([][]float64, n)
	for i := 0; i < n; i++ {
		buckets[i] = []float64{}
	}

	// 2) Colocar los elementos del arreglo en diferentes buckets
	for i := 0; i < n; i++ {
		idx := int(arr[i] * float64(n))
		buckets[idx] = append(buckets[idx], arr[i])
	}

	// 3) Ordenar los buckets individuales
	for i := 0; i < n; i++ {
		bucket := buckets[i]
		for j := 1; j < len(bucket); j++ {
			key := bucket[j]
			k := j - 1
			for k >= 0 && bucket[k] > key {
				bucket[k+1] = bucket[k]
				k--
			}
			bucket[k+1] = key
		}
	}

	// 4) Concatenar todos los buckets en el arreglo original
	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[index] = buckets[i][j]
			index++
		}
	}
}

func LlamarBucketSort() {
	startTime := time.Now()
	arr := utilidades.RecuperarArregloFlotante()
	bucketSort(arr)
	utilidades.ImprimirArregloFlotante(arr)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
