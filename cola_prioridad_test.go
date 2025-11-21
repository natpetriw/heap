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

func TestHeapVacioRecienCreado(t *testing.T) {
	t.Log("Heap recién creado debe comenzar vacío y lanzar panics al acceder")
	h := TDAHeap.CrearHeap(cmpEnteros)
	require.True(t, h.EstaVacia())
	require.Equal(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
}

func TestEncolarUnElemento(t *testing.T) {
	t.Log("Encolar un único elemento y verificar su comportamiento")
	h := TDAHeap.CrearHeap(cmpEnteros)
	h.Encolar(99)
	require.False(t, h.EstaVacia())
	require.Equal(t, 1, h.Cantidad())
	require.EqualValues(t, 99, h.VerMax())
	require.EqualValues(t, 99, h.Desencolar())
	require.True(t, h.EstaVacia())
}

func TestEncolarVariosElementos(t *testing.T) {
	t.Log("Encola varios elementos y verifica que el máximo sea correcto")
	h := TDAHeap.CrearHeap(cmpEnteros)
	for _, v := range []int{12, 4, 18, 9, 20, 7} {
		h.Encolar(v)
	}
	require.Equal(t, 6, h.Cantidad())
	require.EqualValues(t, 20, h.VerMax())
	require.EqualValues(t, 20, h.Desencolar())
	require.EqualValues(t, 18, h.VerMax())
}

func TestDesencolarEnOrdenDePrioridad(t *testing.T) {
	t.Log("Verifica que los elementos se desencolen en orden descendente según prioridad")
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

func TestCrearHeapDesdeArreglo(t *testing.T) {
	t.Log("Construye un heap a partir de un arreglo inicial")
	valores := []int{6, 3, 8, 1, 9, 2}
	h := TDAHeap.CrearHeapArr(valores, cmpEnteros)
	require.Equal(t, len(valores), h.Cantidad())
	require.EqualValues(t, 9, h.VerMax())
}

func TestHeapComparacionStringsPorLongitud(t *testing.T) {
	t.Log("Heap de strings ordenado según el largo de cada palabra")
	h := TDAHeap.CrearHeap(cmpStringsPorLargo)
	for _, s := range []string{"uva", "manzana", "pera", "kiwi", "sandía"} {
		h.Encolar(s)
	}
	require.EqualValues(t, "manzana", h.VerMax())
	h.Desencolar()
	require.EqualValues(t, "sandía", h.VerMax())
}

func TestHeapConComparadorDeProductos(t *testing.T) {
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

func TestHeapConElementosIguales(t *testing.T) {
	t.Log("Inserta elementos repetidos y verifica que el heap mantenga consistencia")
	h := TDAHeap.CrearHeap(cmpEnteros)
	for i := 0; i < 4; i++ {
		h.Encolar(42)
	}
	for i := 0; i < 4; i++ {
		require.EqualValues(t, 42, h.Desencolar())
	}
	require.True(t, h.EstaVacia())
}

func TestHeapVolumen(t *testing.T) {
	t.Log("Prueba de gran volumen para validar eficiencia y correcto funcionamiento")
	const N = 100000
	h := TDAHeap.CrearHeap(cmpEnteros)
	for i := 0; i < N; i++ {
		h.Encolar(i)
	}
	require.Equal(t, N, h.Cantidad())
	require.EqualValues(t, N-1, h.VerMax())

	for i := N - 1; i >= 0; i-- {
		maximo := h.Desencolar()
		require.EqualValues(t, i, maximo)
	}
	require.True(t, h.EstaVacia())
}

func TestHeapConDatosAleatorios(t *testing.T) {
	t.Log("Crea un heap con datos mezclados y verifica la propiedad de heap")
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

func TestHeapSortOrdenaCorrectamente(t *testing.T) {
	t.Log("Verifica que el HeapSort ordene correctamente un arreglo")
	arr := []int{8, 2, 5, 1, 9, 4}
	TDAHeap.HeapSort(arr, cmpEnteros)
	require.EqualValues(t, []int{1, 2, 4, 5, 8, 9}, arr)
}

func TestPanicsNoModificanElHeap(t *testing.T) {
	t.Log("Las operaciones que lanzan panic no deben alterar el heap")
	h := TDAHeap.CrearHeap(cmpEnteros)
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
	require.Equal(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
	require.True(t, h.EstaVacia())
}
