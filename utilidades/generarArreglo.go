package utilidades

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GenerarArreglog() {
	// Semilla para generar números aleatorios
	rand.Seed(time.Now().UnixNano())

	// Generar un arreglo de 10,000,000 de números aleatorios
	array := make([]int64, 3000)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Int63n(99999999999000) + 1000000000
	}
	// Guardar el arreglo en un archivo de texto
	filename := "datos3.txt"
	err := guardarArregloEnArchivo(filename, array)
	if err != nil {
		fmt.Println("Error al guardar el arreglo:", err)
		return
	}
	fmt.Println("Arreglo guardado en", filename)

}

func guardarArregloEnArchivo(filename string, array []int64) error {
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
