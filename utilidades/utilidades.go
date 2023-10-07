package utilidades

import "fmt"

func ImprimirArreglo(arr []int) {
	fmt.Println("Arreglo ordenado:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}
