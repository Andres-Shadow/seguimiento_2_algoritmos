package algoritmos

import (
	"fmt"
	"time"
)

type BitonicSort struct{}

func (b *BitonicSort) compAndSwap(a []int, i, j, dir int) {
	if (a[i] > a[j] && dir == 1) || (a[i] < a[j] && dir == 0) {
		// Intercambiar elementos
		a[i], a[j] = a[j], a[i]
	}
}

func (b *BitonicSort) bitonicMerge(a []int, low, cnt, dir int) {
	if cnt > 1 {
		k := cnt / 2
		for i := low; i < low+k; i++ {
			b.compAndSwap(a, i, i+k, dir)
		}
		b.bitonicMerge(a, low, k, dir)
		b.bitonicMerge(a, low+k, k, dir)
	}
}

func (b *BitonicSort) bitonicSort(a []int, low, cnt, dir int) {
	if cnt > 1 {
		k := cnt / 2
		b.bitonicSort(a, low, k, 1)
		b.bitonicSort(a, low+k, k, 0)
		b.bitonicMerge(a, low, cnt, dir)
	}
}

func (b *BitonicSort) Sort(a []int, N, up int) {
	b.bitonicSort(a, 0, N, up)
}

func printArra(arr []int) {
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func LlamarBitonicSort() {
	startTime := time.Now()

	arr := []int{3, 7, 4, 8, 6, 2, 1, 5}
	up := 1
	var ob BitonicSort
	ob.Sort(arr, len(arr), up)

	fmt.Println("\nArreglo ordenado")
	printArra(arr)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
