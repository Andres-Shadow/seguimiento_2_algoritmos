package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

// Función para encontrar la brecha entre los elementos
func getNextGap(gap int) int {
	// Reducir la brecha por el factor de reducción
	gap = (gap * 10) / 13
	if gap < 1 {
		return 1
	}
	return gap
}

// Función para ordenar arr[] usando Comb Sort
func combSort(arr []int) {
	n := len(arr)

	// Inicializar la brecha
	gap := n

	// Inicializar swapped como true para asegurarse de que
	// el bucle se ejecute
	swapped := true

	// Continuar mientras la brecha sea mayor que 1 o la última
	// iteración causó un intercambio
	for gap != 1 || swapped {
		// Encontrar la siguiente brecha
		gap = getNextGap(gap)

		// Inicializar swapped como falso para que podamos
		// comprobar si se produjo un intercambio o no
		swapped = false

		// Comparar todos los elementos con la brecha actual
		for i := 0; i < n-gap; i++ {
			if arr[i] > arr[i+gap] {
				// Intercambiar arr[i] y arr[i+gap]
				temp := arr[i]
				arr[i] = arr[i+gap]
				arr[i+gap] = temp

				// Establecer swapped
				swapped = true
			}
		}
	}
}

func LlamarCombSort(tam int) {
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
	combSort(arr)
	utilidades.ImprimirArreglo(arr)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
