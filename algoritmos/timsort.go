package algoritmos

import (
	"fmt"
	"math"
	"seg_2_algoritmos/utilidades"
	"time"
)

var MIN_MERGE = 32

func minRunLength(n int) int {
	if n < 0 {
		panic("AssertionError: n >= 0")
	}

	r := 0
	for n >= MIN_MERGE {
		r |= (n & 1)
		n >>= 1
	}
	return n + r
}

func insertionSort(arr []int, left int, right int) {
	for i := left + 1; i <= right; i++ {
		temp := arr[i]
		j := i - 1
		for j >= left && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func merge(arr []int, l int, m int, r int) {
	len1 := m - l + 1
	len2 := r - m
	left := make([]int, len1)
	right := make([]int, len2)

	for x := 0; x < len1; x++ {
		left[x] = arr[l+x]
	}
	for x := 0; x < len2; x++ {
		right[x] = arr[m+1+x]
	}

	i := 0
	j := 0
	k := l

	for i < len1 && j < len2 {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len1 {
		arr[k] = left[i]
		k++
		i++
	}

	for j < len2 {
		arr[k] = right[j]
		k++
		j++
	}
}

func timSort(arr []int, n int) {
	minRun := minRunLength(MIN_MERGE)

	for i := 0; i < n; i += minRun {
		insertionSort(arr, i, int(math.Min(float64(i+MIN_MERGE-1), float64(n-1))))
	}

	for size := minRun; size < n; size = 2 * size {
		for left := 0; left < n; left += 2 * size {
			mid := left + size - 1
			right := int(math.Min(float64(left+2*size-1), float64(n-1)))

			if mid < right {
				merge(arr, left, mid, right)
			}
		}
	}
}

func LlamarTimSort(tam int) {

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

	n := len(arr)

	timSort(arr, n)

	utilidades.ImprimirArreglo(arr)
	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
