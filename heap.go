package heap

type Heap[T any] struct {
	datos []T
	cmp   func(T, T) int 
}

func CrearHeap[T any](cmp func(T, T) int) *Heap[T] {
	return &Heap[T]{datos: []T{}, cmp: cmp}
}

func (heap *Heap[T]) EstaVacia() bool{
	return  len(heap.datos) == 0
}

func (heap *Heap[T]) Encolar(elem T) {
	heap.datos = append(heap.datos, elem)
	upHeap(heap.datos, heap.cmp, len(heap.datos)-1)	
}

func (heap *Heap[T]) VerMax() T{
	if heap.EstaVacia(){
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *Heap[T]) Desencolar() T{
	if heap.EstaVacia(){
		panic("La cola esta vacia")
	}
	maximo_elemento := heap.VerMax()
	ultimo_elemento := len(heap.datos) - 1
	heap.datos[0] = heap.datos[ultimo_elemento]
	heap.datos = heap.datos[:ultimo_elemento]
	
	if !heap.EstaVacia(){
		downHeap(heap.datos, heap.cmp, 0)
	}
	return maximo_elemento
}

func (heap *Heap[T]) Cantidad() int{
	return len(heap.datos)
}
