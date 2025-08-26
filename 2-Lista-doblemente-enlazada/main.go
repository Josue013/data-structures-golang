package main

import (
	"fmt"
	"os"
	"os/exec"
)

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
		fmt.Println("[]")
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

func (l *Lista_doble_enlazada) Graficar() {
	os.Mkdir("Reportes", os.ModePerm)

	dot := "digraph G{\n"
	dot += "	rankdir=LR;\n"
	dot += "	node [shape = record, height = 0.1]\n"

	// nodo inicial nil
	dot += "nil_inicio [label = \"nil\", shape = square];\n"

	// nodos de la lista
	actual := l.Cabeza
	for i := 0; i < l.Tamanio; i++ {
		dot += fmt.Sprintf("node%d [label = \"{<f0>%d | <f1>}\"];\n", i, actual.Dato)
		actual = actual.Siguiente
	}

	// Nodo final nil
	dot += "nil_final [label = \"nil\", shape = square];\n"

	if l.Cabeza != nil {
		dot += "node0 -> nil_inicio;\n"
		dot += "nil_inicio -> node0 [style=invis];\n"
	}

	// Conexiones dobles entre nodos
	actual = l.Cabeza
	for i := 0; i < l.Tamanio-1; i++ {
		dot += fmt.Sprintf("node%d -> node%d;\n", i, i+1)
		dot += fmt.Sprintf("node%d -> node%d;\n", i+1, i)
		actual = actual.Siguiente
	}

	// Conexión último nodo -> nil_final
	if l.Tamanio > 0 {
		dot += fmt.Sprintf("node%d -> nil_final;\n", l.Tamanio-1)
		dot += fmt.Sprintf("nil_final -> node%d [style=invis];\n", l.Tamanio-1)
	}
	dot += "}"

	f, err := os.Create("Reportes/lista_doble_enlazada.dot")
	if err != nil {
		fmt.Println("Error creando archivo:", err)
		return
	}
	defer f.Close()
	f.WriteString(dot)

	cmd := exec.Command("dot", "-Tpng", "Reportes/lista_doble_enlazada.dot", "-o", "Reportes/lista_doble_enlazada.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando Graphviz:", err)
		return
	}
	fmt.Println("Imagen generada exitosamente")
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
	listaDoblementeEnlazada.Graficar()

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

	// listaDoblementeEnlazada.Graficar()
}
