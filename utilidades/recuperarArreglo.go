package utilidades

func RecuperarArreglo() []int {
	leido, _ := LeerArregloDesdeArchivo1("D:\\Goland projects\\seguimiento_2_algoritmos\\datos.txt")
	return leido
}


func RecuperarArregloFlotante() []float64 {
	leido, _ := LeerArregloDesdeArchivo2("./datos.txt")
	return leido
}
