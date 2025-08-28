package main

// NodoEncabezado: representa un encabezado de fila o columna
type NodoEncabezado struct {
	ID        int
	Siguiente *NodoEncabezado
	Anterior  *NodoEncabezado
	Acceso    *NodoInterno
}

// ListaEncabezado: lista doblemente enlazada para manejar los encabezados
type ListaEncabezado struct {
	Primero *NodoEncabezado
	Ultimo  *NodoEncabezado
}

// Insertar: inserta un nuevo encabezado en la lista de manera ordenada
func (l *ListaEncabezado) Insertar(nuevo *NodoEncabezado) {
	if l.Primero == nil {
		l.Primero = nuevo
		l.Ultimo = nuevo
		return
	}
	// Inserción ordenada
	actual := l.Primero
	for actual != nil && nuevo.ID > actual.ID {
		actual = actual.Siguiente
	}
	if actual == l.Primero {
		nuevo.Siguiente = l.Primero
		l.Primero.Anterior = nuevo
		l.Primero = nuevo
	} else if actual == nil {
		l.Ultimo.Siguiente = nuevo
		nuevo.Anterior = l.Ultimo
		l.Ultimo = nuevo
	} else {
		nuevo.Anterior = actual.Anterior
		nuevo.Siguiente = actual
		actual.Anterior.Siguiente = nuevo
		actual.Anterior = nuevo
	}
}

// Buscar: busca un encabezado por su ID
func (l *ListaEncabezado) Buscar(id int) *NodoEncabezado {
	actual := l.Primero
	for actual != nil {
		if actual.ID == id {
			return actual
		}
		actual = actual.Siguiente
	}
	return nil
}
