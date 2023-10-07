package utilidades

func RecuperarArreglo() []int {
	leido, _ := LeerArregloDesdeArchivo1("./datos.txt")
	return leido
}
