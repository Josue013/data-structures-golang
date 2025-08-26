package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Nodo struct {
	Siguiente *Nodo
	Dato      int
}

type Lista_simple struct {
	Cabeza  *Nodo
	Tamanio int
}

func (l *Lista_simple) Insertar(dato int) {
	nuevo := &Nodo{Dato: dato}

	if l.Cabeza == nil {
		l.Cabeza = nuevo
		nuevo.Siguiente = nil
		l.Tamanio++
	} else {
		actual := l.Cabeza
		for actual.Siguiente != nil {
			actual = actual.Siguiente
		}
		actual.Siguiente = nuevo
		nuevo.Siguiente = nil
		l.Tamanio++
	}
}

func (l *Lista_simple) Buscar(dato int) bool {

	actual := l.Cabeza

	for actual != nil {
		if actual.Dato == dato {
			return true
		}
		actual = actual.Siguiente
	}
	return false
}

func (l *Lista_simple) Eliminar(dato int) {
	actual := l.Cabeza
	previo := l.Cabeza

	if l.Cabeza == nil {
		return
	}

	if l.Cabeza.Dato == dato {
		l.Cabeza = l.Cabeza.Siguiente
		l.Tamanio--
	} else {
		for actual.Siguiente != nil {
			if actual.Dato == dato {
				previo.Siguiente = actual.Siguiente
				l.Tamanio--
				return
			}
			previo = actual
			actual = actual.Siguiente
		}

		if actual.Dato == dato {
			previo.Siguiente = nil
			l.Tamanio--
		}
	}
}

func (l *Lista_simple) Imprimir() {
	actual := l.Cabeza
	if l.Cabeza == nil {
		println("[]")
		return
	}
	print("[")
	for actual.Siguiente != nil {
		print(actual.Dato, ", ")
		actual = actual.Siguiente
	}
	print(actual.Dato, "]\n")
}

func (l *Lista_simple) Graficar() {
	os.Mkdir("Reportes", os.ModePerm)

	dot := "digraph G{\n"
	dot += "	rankdir=LR;\n"
	dot += "	node [shape = record, height = 0.1]\n"

	actual := l.Cabeza
	for i := 0; i < l.Tamanio; i++ {
		dot += fmt.Sprintf("node%d [label = \"{<f0>%d|<f1>}\"];\n", i, actual.Dato)
		actual = actual.Siguiente
	}

	actual = l.Cabeza
	for i := 0; i < l.Tamanio-1; i++ {
		dot += fmt.Sprintf("node%d:f1 -> node%d:f0;\n", i, i+1)
		actual = actual.Siguiente
	}
	// if l.Tamanio > 0 {
	// 	dot += fmt.Sprintf("node%d:f1 -> node0:f0 [contraint=false];\n", l.Tamanio-1)
	// }

	dot += "nil_final [label = \"nil\", shape = square];\n"
	if l.Tamanio > 0 {
		dot += fmt.Sprintf("node%d -> nil_final;\n", l.Tamanio-1)
	}

	dot += "}"

	f, err := os.Create("Reportes/lista_simple.dot")
	if err != nil {
		fmt.Println("Error creando archivo:", err)
		return
	}
	defer f.Close()
	f.WriteString(dot)

	cmd := exec.Command("dot", "-Tpng", "Reportes/lista_simple.dot", "-o", "Reportes/lista_simple.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando Graphviz:", err)
		return
	}
	fmt.Println("Imagen generada exitosamente")
}

func main() {
	lista_simple := Lista_simple{}
	lista_simple.Insertar(1)
	lista_simple.Insertar(2)
	lista_simple.Insertar(3)
	lista_simple.Insertar(4)
	lista_simple.Insertar(5)

	lista_simple.Imprimir()
	lista_simple.Graficar()

	println(lista_simple.Buscar(5))

	lista_simple.Eliminar(1)
	lista_simple.Imprimir()

	lista_simple.Eliminar(3)
	lista_simple.Imprimir()

	lista_simple.Eliminar(5)
	lista_simple.Imprimir()
	println("Tamaño lista:", lista_simple.Tamanio)

	println("------------------------ Más pruebas ---------------------------------")
	lista := &Lista_simple{}

	fmt.Println("=== PRUEBA: Insertar ===")
	lista.Insertar(1)
	lista.Insertar(2)
	lista.Insertar(3)
	lista.Imprimir() // Esperado: 3 -> 2 -> 1 -> nil

	fmt.Println("=== PRUEBA: Buscar ===")
	fmt.Println("Buscar 2:", lista.Buscar(2)) // true
	fmt.Println("Buscar 5:", lista.Buscar(5)) // false

	fmt.Println("=== PRUEBA: Eliminar ===")
	lista.Eliminar(2)
	lista.Imprimir() // Esperado: 3 -> 1 -> nil

	lista.Eliminar(3)
	lista.Imprimir() // Esperado: 1 -> nil

	lista.Eliminar(1)
	lista.Imprimir() // Esperado: nil

	fmt.Println("Buscar en lista vacía:", lista.Buscar(10)) // false

}
