package main

type Nodo struct {
	Dato      int
	Siguiente *Nodo
	Anterior  *Nodo
}

type Lista_doble_enlazada struct {
	Cabeza  *Nodo
	Cola    *Nodo
	Tamanio int
}

func (l *Lista_doble_enlazada) Insertar(dato int) {
	nuevo := &Nodo{Dato: dato}

	if l.Cabeza == nil && l.Cola == nil {
		l.Cabeza = nuevo
		l.Cola = nuevo
		l.Tamanio++
	} else {
		l.Cola.Siguiente = nuevo
		nuevo.Anterior = l.Cola
		l.Cola = nuevo
		l.Tamanio++
	}
}

func (l *Lista_doble_enlazada) Buscar(dato int) bool {
	actual := l.Cabeza

	for actual != nil {
		if actual.Dato == dato {
			return true
		}
		actual = actual.Siguiente
	}
	return false
}

func (l *Lista_doble_enlazada) Eliminar(dato int) {
	if l.Cabeza == nil {
		return
	}

	actual := l.Cabeza
	previo := l.Cabeza

	if l.Cabeza.Dato == dato {
		if l.Cabeza.Siguiente != nil {
			l.Cabeza = l.Cabeza.Siguiente
			l.Cabeza.Anterior = nil
		} else {
			l.Cabeza = nil
			l.Cola = nil
		}
		l.Tamanio--
	} else if l.Cola.Dato == dato {
		if l.Cola.Anterior != nil {
			l.Cola = l.Cola.Anterior
			l.Cola.Siguiente = nil
		} else {
			l.Cola = nil
			l.Cabeza = nil
		}
		l.Tamanio--
	} else {
		for actual != nil && actual.Dato != dato {
			previo = actual
			actual = actual.Siguiente
		}
		if actual == nil {
			return
		}

		previo.Siguiente = actual.Siguiente
		actual.Siguiente.Anterior = previo
		l.Tamanio--
	}
}

func (l *Lista_doble_enlazada) Imprimir() {
	if l.Cabeza == nil {
		println("[]")
		return
	}
	actual := l.Cabeza
	print("[")
	for actual != l.Cola {
		print(actual.Dato, ", ")
		actual = actual.Siguiente
	}
	print(actual.Dato, "]\n")
}

func main() {
	listaDoblementeEnlazada := Lista_doble_enlazada{}
	listaDoblementeEnlazada.Insertar(1)
	listaDoblementeEnlazada.Insertar(2)
	listaDoblementeEnlazada.Insertar(3)
	listaDoblementeEnlazada.Insertar(4)
	listaDoblementeEnlazada.Insertar(5)
	listaDoblementeEnlazada.Insertar(6)
	listaDoblementeEnlazada.Insertar(7)
	listaDoblementeEnlazada.Insertar(8)

	listaDoblementeEnlazada.Imprimir()

	println("Buscar: ", listaDoblementeEnlazada.Buscar(7))

	listaDoblementeEnlazada.Eliminar(1)
	listaDoblementeEnlazada.Imprimir()

	listaDoblementeEnlazada.Eliminar(8)
	listaDoblementeEnlazada.Imprimir()

	listaDoblementeEnlazada.Eliminar(4)
	listaDoblementeEnlazada.Imprimir()

	listaDoblementeEnlazada.Eliminar(5)
	listaDoblementeEnlazada.Imprimir()
}
