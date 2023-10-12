package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
	"time"
)

type HeapSort struct{}

func (h *HeapSort) heapify(arr []int, N, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < N && arr[l] > arr[largest] {
		largest = l
	}

	if r < N && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		h.heapify(arr, N, largest)
	}
}

func (h *HeapSort) sort(arr []int) {
	N := len(arr)

	// Build heap (rearrange array)
	for i := N/2 - 1; i >= 0; i-- {
		h.heapify(arr, N, i)
	}

	// One by one extract an element from heap
	for i := N - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// Call heapify on the reduced heap
		h.heapify(arr, i, 0)
	}
}

func LlamarHeapSort(tam int) {
	startTime := time.Now()

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

	// Function call
	h := &HeapSort{}
	h.sort(arr)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
