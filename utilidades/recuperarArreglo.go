package utilidades

func RecuperarArreglo() []int {
	leido, _ := LeerArregloDesdeArchivo1("./datos.txt")
	return leido
}


func RecuperarArregloFlotante() []float64 {
	leido, _ := LeerArregloDesdeArchivo2("./datos.txt")
	return leido
}
