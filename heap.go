package cola_prioridad

const (
	errColaVacia        = "La cola esta vacia"
	tamInicial          = 10
	factorAchicamiento  = 4
	factorAgrandamiento = 2
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
		downHeap(arr, cmp, i, len(arr))
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
		downHeap(arr, cmp, 0, i)
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
	return len(heap.datos) == 0
}

func (heap *heap[T]) Encolar(elem T) {
	heap.redimensionarCapacidadEncolar()

	n := len(heap.datos)
	heap.datos = heap.datos[:n+1]
	heap.datos[n] = elem

	upHeap(heap.datos, heap.cmp, n)
}

func (heap *heap[T]) redimensionarCapacidadEncolar() {
	n := len(heap.datos)
	if n == cap(heap.datos) {
		nuevaCap := cap(heap.datos)*factorAgrandamiento + 1
		if nuevaCap < tamInicial {
			nuevaCap = tamInicial
		}
		nueva := make([]T, n, nuevaCap)
		copy(nueva, heap.datos)
		heap.datos = nueva
	}
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

	downHeap(heap.datos, heap.cmp, 0, len(heap.datos))
	heap.reducirCapacidad()
	return maximo_elemento
}

func (heap *heap[T]) reducirCapacidad() {
	if cap(heap.datos) > tamInicial && len(heap.datos) <= cap(heap.datos)/factorAchicamiento {
		nueva_capacidad := cap(heap.datos) / 2
		if nueva_capacidad < tamInicial {
			nueva_capacidad = tamInicial
		}
		nuevos_datos := make([]T, len(heap.datos), nueva_capacidad)
		copy(nuevos_datos, heap.datos)
		heap.datos = nuevos_datos
	}
}

func (heap *heap[T]) Cantidad() int {
	return len(heap.datos)
}
