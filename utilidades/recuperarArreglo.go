package utilidades

import (
	"fmt"
	"time"
)

func RecuperarArreglo() {
	fmt.Println("Inicio")

	startTime := time.Now() // Registra el tiempo de inicio

	leido, err := LeerArregloDesdeArchivo1("./datos.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	elapsedTime := time.Since(startTime) // Calcula el tiempo transcurrido

	fmt.Println("Longitud del arreglo:", len(leido))
	for i := 0; i < len(leido); i++ {
		//fmt.Println(leido[i])
	}

	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
