package utilidades

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GenerarArreglog(tam int) {

	var nombre string
	var cantidad int
	switch tam {
	case 1:
		nombre = "datos.txt"
		cantidad = 100000
		break
	case 2:
		nombre = "datos2.txt"
		cantidad = 200000
		break
	case 3:
		nombre = "datos3.txt"
		cantidad = 300000
	}
	// Semilla para generar números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Generar un arreglo de 10,000,000 de números aleatorios
	array := make([]int, cantidad)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1000000) + 1
	}
	// Guardar el arreglo en un archivo de texto
	err := guardarArregloEnArchivo(nombre, array)
	if err != nil {
		fmt.Println("Error al guardar el arreglo:", err)
		return
	}
	fmt.Println("Arreglo guardado en", nombre)

}

func guardarArregloEnArchivo(filename string, array []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, num := range array {
		_, err := fmt.Fprintln(file, num)
		if err != nil {
			return err
		}
	}

	return nil
}
