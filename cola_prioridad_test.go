package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	TDAHeap "tdas/cola_prioridad"
	"testing"
)

type Producto struct {
	nombre string
	precio float64
}

func cmpEnteros(a, b int) int {
	return a - b
}

func cmpStringsPorLargo(a, b string) int {
	return len(a) - len(b)
}

func cmpProductosPorPrecio(a, b Producto) int {
	if a.precio > b.precio {
		return 1
	}
	if a.precio < b.precio {
		return -1
	}
	return 0
}

func TestHeapRecienCreado(t *testing.T) {
	t.Log("Heap recién creado debe estar vacío y lanzar panics al acceder")
	h := TDAHeap.CrearHeap[int](cmpEnteros)
	require.True(t, h.EstaVacia())
	require.Equal(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
}

func TestInsertarUnElemento(t *testing.T) {
	t.Log("Insertar un único elemento y verificar comportamiento")
	h := TDAHeap.CrearHeap(cmpEnteros)
	h.Encolar(99)
	require.False(t, h.EstaVacia())
	require.Equal(t, 1, h.Cantidad())
	require.EqualValues(t, 99, h.VerMax())
	require.EqualValues(t, 99, h.Desencolar())
	require.True(t, h.EstaVacia())
}

func TestInsertarVariosElementos(t *testing.T) {
	t.Log("Inserta varios elementos y verifica prioridad máxima")
	h := TDAHeap.CrearHeap(cmpEnteros)
	for _, v := range []int{12, 4, 18, 9, 20, 7} {
		h.Encolar(v)
	}
	require.Equal(t, 6, h.Cantidad())
	require.EqualValues(t, 20, h.VerMax())
	require.EqualValues(t, 20, h.Desencolar())
	require.EqualValues(t, 18, h.VerMax())
}

func TestOrdenDesencoladoDesc(t *testing.T) {
	t.Log("Desencola y verifica que salgan en orden descendente")
	h := TDAHeap.CrearHeap(cmpEnteros)
	elementos := []int{2, 14, 8, 11}
	for _, e := range elementos {
		h.Encolar(e)
	}
	for _, esperado := range []int{14, 11, 8, 2} {
		require.EqualValues(t, esperado, h.Desencolar())
	}
	require.True(t, h.EstaVacia())
}

func TestHeapDesdeArreglo(t *testing.T) {
	t.Log("Construye un heap desde un arreglo inicial")
	valores := []int{6, 3, 8, 1, 9, 2}
	h := TDAHeap.CrearHeapArr(valores, cmpEnteros)
	require.Equal(t, len(valores), h.Cantidad())
	require.EqualValues(t, 9, h.VerMax())
}

func TestHeapStringsPorLongitud(t *testing.T) {
	t.Log("Heap de strings según el largo de cada palabra")
	h := TDAHeap.CrearHeap(cmpStringsPorLargo)
	for _, s := range []string{"uva", "manzana", "pera", "kiwi", "sandía"} {
		h.Encolar(s)
	}
	require.EqualValues(t, "manzana", h.VerMax())
	h.Desencolar()
	require.EqualValues(t, "sandía", h.VerMax())
}

func TestHeapConProductos(t *testing.T) {
	t.Log("Heap de productos ordenado por precio")
	h := TDAHeap.CrearHeap(cmpProductosPorPrecio)
	h.Encolar(Producto{"Auriculares", 75.50})
	h.Encolar(Producto{"Notebook", 980.00})
	h.Encolar(Producto{"Mouse", 35.00})
	h.Encolar(Producto{"Monitor", 250.00})
	h.Encolar(Producto{"Teclado", 120.00})

	require.EqualValues(t, "Notebook", h.VerMax().nombre)
	require.EqualValues(t, 980.00, h.VerMax().precio)

	h.Desencolar()
	require.EqualValues(t, "Monitor", h.VerMax().nombre)
	require.EqualValues(t, 250.00, h.VerMax().precio)
}

func TestHeapElementosIguales(t *testing.T) {
	t.Log("Inserta elementos repetidos y verifica consistencia")
	h := TDAHeap.CrearHeap(cmpEnteros)
	for i := 0; i < 4; i++ {
		h.Encolar(42)
	}
	for i := 0; i < 4; i++ {
		require.EqualValues(t, 42, h.Desencolar())
	}
	require.True(t, h.EstaVacia())
}

func TestHeapGrande(t *testing.T) {
	t.Log("Heap con gran volumen de datos")
	h := TDAHeap.CrearHeap(cmpEnteros)
	for i := 500; i >= 1; i-- {
		h.Encolar(i)
	}
	require.Equal(t, 500, h.Cantidad())
	require.EqualValues(t, 500, h.VerMax())
	require.EqualValues(t, 500, h.Desencolar())
	require.EqualValues(t, 499, h.VerMax())
}

func TestOperacionesAlternadas(t *testing.T) {
	t.Log("Intercala inserciones, eliminaciones y chequeos de máximo")
	h := TDAHeap.CrearHeap(cmpEnteros)
	h.Encolar(10)
	h.Encolar(25)
	h.Encolar(3)
	require.EqualValues(t, 25, h.VerMax())
	h.Desencolar()
	h.Encolar(18)
	h.Encolar(40)
	require.EqualValues(t, 40, h.VerMax())
	require.EqualValues(t, 40, h.Desencolar())
	require.EqualValues(t, 18, h.VerMax())
	h.Encolar(50)
	require.EqualValues(t, 50, h.VerMax())
}

func TestHeapAleatorio(t *testing.T) {
	t.Log("Crea un heap con valores mezclados y verifica propiedad de heap")
	h := TDAHeap.CrearHeap(cmpEnteros)
	datos := []int{15, 42, 3, 27, 19, 8, 60, 12}
	for _, v := range datos {
		h.Encolar(v)
	}
	require.EqualValues(t, 60, h.VerMax())
	require.Equal(t, len(datos), h.Cantidad())
	h.Desencolar()
	require.GreaterOrEqual(t, h.VerMax(), 42)
}

func TestHeapSortBasico(t *testing.T) {
	t.Log("Prueba básica de ordenamiento con HeapSort")
	arr := []int{8, 2, 5, 1, 9, 4}
	TDAHeap.HeapSort(arr, cmpEnteros)
	require.EqualValues(t, []int{1, 2, 4, 5, 8, 9}, arr)
}

func TestPanicsNoAlteranEstado(t *testing.T) {
	t.Log("Las operaciones con panic no deben modificar el heap")
	h := TDAHeap.CrearHeap[int](cmpEnteros)
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
	require.Equal(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
	require.True(t, h.EstaVacia())
}
