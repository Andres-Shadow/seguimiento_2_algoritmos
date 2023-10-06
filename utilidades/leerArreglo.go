package utilidades

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func LeerArregloDesdeArchivo1(filename string) ([]int64, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n") // Dividir en líneas
	array := make([]int64, 0)                  // Inicializar el slice vacío

	for _, str := range lines {
		if str != "" { // Verificar si la línea no está vacía
			num, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, err
			}
			array = append(array, num) // Agregar el número al slice
		}
	}

	return array, nil
}
