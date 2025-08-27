package main

// Nodo: representa un nodo del árbol binario de búsqueda
type Nodo struct {
	Valor     int    // Valor almacenado en el nodo
	Izquierdo *Nodo  // Apuntador al hijo izquierdo
	Derecho   *Nodo  // Apuntador al hijo derecho
}

// NuevoNodo: crea un nuevo nodo con el valor especificado
func NuevoNodo(valor int) *Nodo {
	return &Nodo{Valor: valor}
}
