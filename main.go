package main

import (
	"fmt"
	"os"
	"seg_2_algoritmos/algoritmos"
	"seg_2_algoritmos/utilidades"
	"strconv"
)

func main() {

	opcion := os.Args[1]
	var tam int

	if len(os.Args) < 3 {
		fmt.Println("1. primer tamaño")
		fmt.Println("2. segundo tamaño")
		fmt.Println("3. tercer tamaño")
		fmt.Print("Seleccione el tamaño para el arreglo: ")
		fmt.Scan(&tam)
	} else {
		tam, _ = strconv.Atoi(os.Args[2])
	}

	switch opcion {
	case "1":
		algoritmos.LlamarTimSort(tam)
		break
	case "2":
		algoritmos.LlamarCombSort(tam)
		break
	case "3":
		algoritmos.LlamarSelectionSort(tam)
		break
	case "4":
		algoritmos.LlamarTreeSort(tam)
		break
	case "5":
		algoritmos.LlamarPigeonHoleSort(tam)
		break
	case "6":
		algoritmos.LlamarBucketSort(tam)
		break
	case "7":
		algoritmos.LlamarQuickSort(tam)
		break
	case "8":
		algoritmos.LlamarHeapSort(tam)
		break
	case "9":
		algoritmos.LlamarBitonicSort(tam)
		break
	case "10":
		algoritmos.LlamarGnomeSort(tam)
		break
	case "11":
		algoritmos.LlamarBinaryInsertionSort(tam)
		break
	case "12":
		algoritmos.LlamarRadixSort(tam)
		break
	case "x":
		utilidades.GenerarArreglog(tam)
		break
	}

}
