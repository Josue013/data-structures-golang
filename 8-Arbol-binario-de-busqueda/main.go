package main

func main() {
	// Crear un nuevo arbol binario de búsqueda
	arbol := NuevoArbolBST()

	arbol.Insertar(20) 
	arbol.Insertar(8) 
	arbol.Insertar(3) 
	arbol.Insertar(1) 
	arbol.Insertar(0) 
	arbol.Insertar(15)
	arbol.Insertar(30)
	arbol.Insertar(48)
	arbol.Insertar(26)
	arbol.Insertar(10)
	arbol.Insertar(7) 
	arbol.Insertar(5) 
	arbol.Insertar(60)
	arbol.Insertar(19)
	arbol.Insertar(11)
	arbol.Insertar(21)
	arbol.Insertar(3) 

	arbol.Eliminar(30)
	arbol.Eliminar(26)
	arbol.Eliminar(21) 
	arbol.Eliminar(15) 
	arbol.Eliminar(20) 

	// Mostrar los diferentes recorridos del árbol
	arbol.PreOrder()
	arbol.InOrder()
	arbol.PostOrder()

	// Generar la visualización gráfica del árbol
	arbol.Graficar()
}
