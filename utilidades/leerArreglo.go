package utilidades

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func LeerArregloDesdeArchivo1(filename string) ([]int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n") // Dividir en líneas
	array := make([]int, 0)                    // Inicializar el slice vacío

	for _, str := range lines {
		if str != "" { // Verificar si la línea no está vacía
			num, _ := strconv.ParseInt(str, 10, 64)
			array = append(array, int(num)) // Agregar el número al slice
		}
	}

	return array, nil
}
