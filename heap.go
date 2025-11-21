package cola_prioridad

const (
	errColaVacia        = "La cola esta vacia"
	tamInicial          = 10
	factorAchicamiento  = 4
	factorAgrandamiento = 2
)

type heap[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func intercambio[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		downHeap(arr, cmp, i, len(arr))
	}
}

func CrearHeapArr[T any](arr []T, cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, len(arr))
	copy(datos, arr)
	h := &heap[T]{datos: datos, cant: len(arr), cmp: cmp}
	heapify(h.datos, h.cmp)
	return h
}

func HeapSort[T any](arr []T, cmp func(T, T) int) {
	heapify(arr, cmp)
	for i := len(arr) - 1; i > 0; i-- {
		intercambio(arr, 0, i)
		downHeap(arr, cmp, 0, i)
	}
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, 0, tamInicial), cant: 0, cmp: cmp}
}

func upHeap[T any](arr []T, cmp func(T, T) int, posHijo int) {
	if posHijo == 0 {
		return
	}
	posPadre := (posHijo - 1) / 2
	if cmp(arr[posHijo], arr[posPadre]) > 0 {
		intercambio(arr, posHijo, posPadre)
		upHeap(arr, cmp, posPadre)
	}
}

func downHeap[T any](arr []T, cmp func(T, T) int, posPadre int, limite int) {
	cant := limite
	posHijoIzq := 2*posPadre + 1
	posHijoDer := 2*posPadre + 2
	mayor := posPadre

	if posHijoIzq < cant && cmp(arr[posHijoIzq], arr[mayor]) > 0 {
		mayor = posHijoIzq
	}
	if posHijoDer < cant && cmp(arr[posHijoDer], arr[mayor]) > 0 {
		mayor = posHijoDer
	}

	if mayor != posPadre {
		intercambio(arr, posPadre, mayor)
		downHeap(arr, cmp, mayor, limite)
	}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *heap[T]) redimensionarHeap() {
	if heap.cant == cap(heap.datos) {
		nuevaCap := cap(heap.datos) * factorAgrandamiento
		if nuevaCap < tamInicial {
			nuevaCap = tamInicial
		}
		nueva := make([]T, heap.cant, nuevaCap)
		copy(nueva, heap.datos)
		heap.datos = nueva
		return
	}

	if cap(heap.datos) > tamInicial && heap.cant <= cap(heap.datos)/factorAchicamiento {
		nuevaCap := cap(heap.datos) / factorAchicamiento
		if nuevaCap < tamInicial {
			nuevaCap = tamInicial
		}
		nueva := make([]T, heap.cant, nuevaCap)
		copy(nueva, heap.datos)
		heap.datos = nueva
		return
	}
}

func (heap *heap[T]) Encolar(elem T) {
	heap.redimensionarHeap()

	if heap.cant < len(heap.datos) {
		heap.datos[heap.cant] = elem
	} else {
		heap.datos = append(heap.datos, elem)
	}
	upHeap(heap.datos, heap.cmp, heap.cant)
	heap.cant++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(errColaVacia)
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(errColaVacia)
	}
	maximo := heap.datos[0]
	ultimo := heap.cant - 1
	heap.datos[0] = heap.datos[ultimo]

	heap.cant--
	heap.datos = heap.datos[:heap.cant]

	downHeap(heap.datos, heap.cmp, 0, heap.cant)
	heap.redimensionarHeap()
	return maximo
}

func (heap *heap[T]) Cantidad() int {
	return heap.cant
}
