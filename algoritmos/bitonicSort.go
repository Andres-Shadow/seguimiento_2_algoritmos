package algoritmos

import (
	"fmt"
	"seg_2_algoritmos/utilidades"
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

func LlamarBitonicSort(tam int) {
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
	
	up := 1
	var ob BitonicSort
	ob.Sort(arr, len(arr), up)

	elapsedTime := time.Since(startTime)
	fmt.Println("Tiempo de ejecución:", elapsedTime)
}
