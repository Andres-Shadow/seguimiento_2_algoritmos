package utilidades

import (
	"fmt"
)

func RecuperarArreglo() []int {
	fmt.Println("Inicio")

	leido, _ := LeerArregloDesdeArchivo1("./datos.txt")

	fmt.Println("Longitud del arreglo:", len(leido))
	for i := 0; i < len(leido); i++ {
		//fmt.Println(leido[i])
	}

	return leido
}
