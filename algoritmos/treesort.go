package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

// Estructura para representar un nodo en el árbol binario de búsqueda
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Función para insertar un nuevo nodo en el árbol binario de búsqueda de manera recursiva
func insertRec(root *Node, key int) *Node {
	if root == nil {
		return &Node{key, nil, nil}
	}

	if key < root.key {
		root.left = insertRec(root.left, key)
	} else if key > root.key {
		root.right = insertRec(root.right, key)
	}

	return root
}

// Función para realizar el recorrido en orden del árbol binario de búsqueda
func inorderRec(root *Node) {
	if root != nil {
		inorderRec(root.left)
		fmt.Println(root.key, " ")
		inorderRec(root.right)
	}
}

// Función para crear el árbol binario de búsqueda a partir de un arreglo
func treeins(arr []int) *Node {
	var root *Node
	for i := 0; i < len(arr); i++ {
		root = insertRec(root, arr[i])
	}
	return root
}

func LlamarTreeSort(tam int) {
	startTime := time.Now()
	var root *Node
	var arr []int
	switch tam {
	case 1:
		arr = utilidades.RecuperarArreglo("datos.txt")
		break
	case 2:
		arr = utilidades.RecuperarArreglo("datos2.txt")
		break
	case 3:
		arr = utilidades.RecuperarArreglo("datos3.txt")
		break
	}
	root = treeins(arr)
	inorderRec(root)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
