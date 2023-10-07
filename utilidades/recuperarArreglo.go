package utilidades

func RecuperarArreglo() []int {
	leido, _ := LeerArregloDesdeArchivo1("D:\\Goland projects\\seguimiento_2_algoritmos\\datos.txt")
	return leido
}
