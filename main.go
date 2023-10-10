package main

import (
	"fmt"
	"os"
	_ "os"
	"seg_2_algoritmos/algoritmos"
	_ "seg_2_algoritmos/algoritmos"
)

func main() {

	opcion := os.Args[1]
	var tam int
	fmt.Println("1. primer tamaño")
	fmt.Println("2. segundo tamaño")
	fmt.Println("3. tercer tamaño")
	fmt.Print("Seleccione el tamaño para el arreglo: ")
	fmt.Scan(&tam)
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
		//algoritmos.LlamarPigeonHoleSort()
		break
	case "8":
		//algoritmos.LlamarPigeonHoleSort()
		break
	case "9":
		//algoritmos.LlamarPigeonHoleSort()
		break
	case "10":
		//algoritmos.LlamarPigeonHoleSort()
		break
	}

}
