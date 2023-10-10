package utilidades

func RecuperarArreglo(nombre string) []int {
	leido, _ := LeerArregloDesdeArchivo1(nombre)
	return leido
}

func RecuperarArregloFlotante(nombre string) []float64 {
	leido, _ := LeerArregloDesdeArchivo2(nombre)
	return leido
}

func RecuperarArregloGrande(nombre string) []int64 {
	leido, _ := LeerArregloDesdeArchivo3(nombre)
	return leido
}
