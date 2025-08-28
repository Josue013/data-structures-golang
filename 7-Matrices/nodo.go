package main

// NodoInterno: representa un nodo en la matriz dispersa
type NodoInterno struct {
	X, Y      int          // Posiciones en la matriz
	Valor     rune         // Carácter almacenado
	Arriba    *NodoInterno // Puntero al nodo superior
	Abajo     *NodoInterno // Puntero al nodo inferior
	Izquierda *NodoInterno // Puntero al nodo izquierdo
	Derecha   *NodoInterno // Puntero al nodo derecho
}
