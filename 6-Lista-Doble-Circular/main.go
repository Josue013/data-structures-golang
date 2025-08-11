package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Nodo: representa un nodo de la lista doble circular
type Nodo struct {
	Dato      int
	Siguiente *Nodo
	Anterior  *Nodo
}

// ListaDobleCircular: representa la lista doblemente enlazada circular
type ListaDobleCircular struct {
	Primero *Nodo
	Ultimo  *Nodo
	Tamanio int
}

// Agregar: agrega un nuevo nodo a la lista doble circular
func (l *ListaDobleCircular) Agregar(dato int) {
	nuevo := &Nodo{Dato: dato}
	if l.Primero == nil && l.Ultimo == nil {
		l.Primero = nuevo
		l.Ultimo = nuevo
		l.Ultimo.Siguiente = l.Primero
		l.Primero.Anterior = l.Ultimo
	} else {
		l.Ultimo.Siguiente = nuevo
		nuevo.Anterior = l.Ultimo
		l.Ultimo = nuevo
		l.Ultimo.Siguiente = l.Primero
		l.Primero.Anterior = l.Ultimo
	}
	l.Tamanio++
}

// Mostrar: imprime los elementos de la lista en orden
func (l *ListaDobleCircular) Mostrar() {
	actual := l.Primero
	contador := 0
	for contador < l.Tamanio {
		fmt.Println(actual.Dato)
		actual = actual.Siguiente
		contador++
	}
}

// Reversa: imprime los elementos de la lista en orden inverso
func (l *ListaDobleCircular) Reversa() {
	actual := l.Ultimo
	contador := 0
	for contador < l.Tamanio {
		fmt.Println(actual.Dato)
		actual = actual.Anterior
		contador++
	}
}

// ObtenerDato: busca un dato en la lista y lo retorna si existe
func (l *ListaDobleCircular) ObtenerDato(dato int) *int {
	actual := l.Primero
	contador := 0
	for contador < l.Tamanio {
		if actual.Dato == dato {
			return &actual.Dato
		}
		actual = actual.Siguiente
		contador++
	}
	return nil
}

// Graficar: genera un archivo .dot y una imagen PNG usando Graphviz
func (l *ListaDobleCircular) Graficar() {
	os.MkdirAll("reportes", os.ModePerm)
	dot := "digraph G {\n  rankdir=LR;\n  node [shape = record, height = .1]\n"
	actual := l.Primero
	contador := 0
	// Crear nodos
	for contador < l.Tamanio {
		dot += fmt.Sprintf("node%d [label = \"{<f1>|%d|<f2>}\"];\n", contador, actual.Dato)
		actual = actual.Siguiente
		contador++
	}
	// Crear apuntadores
	for i := 0; i < l.Tamanio-1; i++ {
		dot += fmt.Sprintf("node%d:f2 -> node%d:f1 [dir=both];\n", i, i+1)
	}
	// Apuntadores de los extremos
	if l.Tamanio > 0 {
		dot += fmt.Sprintf("node0:f1 -> node%d:f2 [dir=both constraint=false];\n", l.Tamanio-1)
	}
	dot += "}"

	// Escribir archivo .dot
	f, err := os.Create("reportes/lista_doble_circular.dot")
	if err != nil {
		fmt.Println("Error creando archivo:", err)
		return
	}
	defer f.Close()
	f.WriteString(dot)

	// Generar imagen con Graphviz
	cmd := exec.Command("dot", "-Tpng", "reportes/lista_doble_circular.dot", "-o", "reportes/lista_doble_circular.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando Graphviz:", err)
		return
	}
	fmt.Println("Imagen generada en reportes/lista_doble_circular.png")
}

func main() {
	lista := ListaDobleCircular{}
	lista.Agregar(10)
	lista.Agregar(20)
	lista.Agregar(30)
	lista.Agregar(40)
	lista.Agregar(50)

	fmt.Println("Lista en orden:")
	lista.Mostrar()

	fmt.Println("Lista en reversa:")
	lista.Reversa()

	lista.Graficar()

	if dato := lista.ObtenerDato(30); dato != nil {
		fmt.Println("Encontrado:", *dato)
	} else {
		fmt.Println("No encontrado")
	}
}
