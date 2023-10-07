package main

import (
	"os"
	"seg_2_algoritmos/algoritmos"
)

func main() {
	opcion := os.Args[1]

	switch opcion {
	case "1":
		algoritmos.LlamarTimSort()
		break
	case "2":
		algoritmos.LlamarCombSort()
		break
	case "3":
		algoritmos.LlamarSelectionSort()
		break
	case "4":
		algoritmos.LlamarTreeSort()
		break
	case "5":
		algoritmos.LlamarPigeonHoleSort()
		break
	}

}
