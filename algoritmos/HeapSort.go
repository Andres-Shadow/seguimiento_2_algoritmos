package algoritmos

import (
	"fmt"
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

func printArray(arr []int) {
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func LlamarHeapSort() {
	startTime := time.Now()

	//arr := utilidades.RecuperarArreglo()
	arr := []int{12, 11, 13, 5, 6, 7}
	//N := len(arr)

	// Function call
	h := &HeapSort{}
	h.sort(arr)

	fmt.Println("Array ordenado :")
	printArray(arr)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
