package cola_prioridad

const (
	errColaVacia = "La cola esta vacia"
)

type heap[T any] struct {
	datos []T
	cmp   func(T, T) int
}

func intercambio[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		downHeap(arr, cmp, i)
	}
}

func CrearHeapArr[T any](arr []T, cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, len(arr))
	copy(datos, arr)
	heap := &heap[T]{datos: datos, cmp: cmp}
	heapify(heap.datos, heap.cmp)
	return heap
}

func HeapSort[T any](arr []T, cmp func(T, T) int) {
	heapify(arr, cmp)
	for i := len(arr) - 1; i > 0; i-- {
		intercambio(arr, 0, i)
		downHeap(arr[:i], cmp, 0)
	}
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: []T{}, cmp: cmp}
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

func downHeap[T any](arr []T, cmp func(T, T) int, posPadre int) {
	cant := len(arr)
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
		downHeap(arr, cmp, mayor)
	}
}

func (heap *heap[T]) EstaVacia() bool {
	return len(heap.datos) == 0
}

func (heap *heap[T]) Encolar(elem T) {
	heap.datos = append(heap.datos, elem)
	upHeap(heap.datos, heap.cmp, len(heap.datos)-1)
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
	maximo_elemento := heap.VerMax()
	ultimo_elemento := len(heap.datos) - 1
	heap.datos[0] = heap.datos[ultimo_elemento]
	heap.datos = heap.datos[:ultimo_elemento]

	if !heap.EstaVacia() {
		downHeap(heap.datos, heap.cmp, 0)
	}

	heap.achicarSlice()
	return maximo_elemento
}

func (heap *heap[T]) achicarSlice() {
	if len(heap.datos) > 0 && len(heap.datos) <= cap(heap.datos)/4 {
		datosReducidos := make([]T, len(heap.datos))
		copy(datosReducidos, heap.datos)
		heap.datos = datosReducidos
	}
}

func (heap *heap[T]) Cantidad() int {
	return len(heap.datos)
}
