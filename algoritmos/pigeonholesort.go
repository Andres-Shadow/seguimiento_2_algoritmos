package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

func pigeonholeSort(arr []int64) {
	// Encontrar el valor mínimo y máximo en el arreglo
	minVal, maxVal := arr[0], arr[0]
	for _, num := range arr {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	println("paso el primer for")
	// Crear un mapa para los cajones
	cajones := make(map[int64]int)

	// Colocar los elementos en los cajones
	for _, num := range arr {
		cajones[num]++
	}

	println("paso el segundo for")

	// Extraer los elementos ordenados de los cajones
	idx := 0
	for val := minVal; val <= maxVal; val++ {
		count, found := cajones[val]
		if found {
			for count > 0 {
				arr[idx] = val
				idx++
				count--
			}
		}
	}
	println("paso el tercer for")
}

func LlamarPigeonHoleSort(tam int) {
	startTime := time.Now()
	var arr []int64
	switch tam {
	case 1:
		arr = utilidades.RecuperarArregloGrande("datos.txt")
		break
	case 2:
		arr = utilidades.RecuperarArregloGrande("datos2.txt")
		break
	case 3:
		arr = utilidades.RecuperarArregloGrande("datos3.txt")
		break
	}

	// Llamar al algoritmo de ordenamiento pigeonhole
	pigeonholeSort(arr)
	utilidades.ImprimirArreglo2(arr)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
