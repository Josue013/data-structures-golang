package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Nodo: representa un nodo de la lista circular
type Nodo struct {
	Dato      int
	Siguiente *Nodo
}

// ListaCircular: representa la lista circular
type ListaCircular struct {
	Primero *Nodo
	Ultimo  *Nodo
	Tamanio int
}

// Agregar: agrega un nuevo nodo a la lista circular
func (l *ListaCircular) Agregar(dato int) {
	nuevo := &Nodo{Dato: dato}
	if l.Primero == nil && l.Ultimo == nil {
		l.Primero = nuevo
		l.Ultimo = nuevo
		l.Ultimo.Siguiente = l.Primero
	} else {
		l.Ultimo.Siguiente = nuevo
		l.Ultimo = nuevo
		l.Ultimo.Siguiente = l.Primero
	}
	l.Tamanio++
}

// Mostrar: imprime los elementos de la lista circular
func (l *ListaCircular) Mostrar() {
	actual := l.Primero
	contador := 0
	for contador < l.Tamanio {
		fmt.Println(actual.Dato)
		actual = actual.Siguiente
		contador++
	}
}

// Graficar: genera un archivo .dot para visualizar la lista circular
func (l *ListaCircular) Graficar() {
	// Crear carpeta Reportes si no existe
	os.MkdirAll("Reportes", os.ModePerm)

	dot := "digraph G {\n  rankdir=LR;\n  node [shape = record, height = .1]\n"
	actual := l.Primero
	for i := 0; i < l.Tamanio; i++ {
		dot += fmt.Sprintf("node%d [label = \"{<f0>%d|<f1>}\"];\n", i, actual.Dato)
		actual = actual.Siguiente
	}
	// Relaciones
	actual = l.Primero
	for i := 0; i < l.Tamanio-1; i++ {
		dot += fmt.Sprintf("node%d:f1 -> node%d:f0;\n", i, i+1)
		actual = actual.Siguiente
	}
	if l.Tamanio > 0 {
		dot += fmt.Sprintf("node%d:f1 -> node0:f0 [constraint=false];\n", l.Tamanio-1)
	}
	dot += "}"

	// Escribir archivo .dot
	f, err := os.Create("Reportes/lista_circular.dot")
	if err != nil {
		fmt.Println("Error creando archivo:", err)
		return
	}
	defer f.Close()
	f.WriteString(dot)

	// Generar imagen con Graphviz
	cmd := exec.Command("dot", "-Tpng", "Reportes/lista_circular.dot", "-o", "Reportes/lista_circular.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando Graphviz:", err)
		return
	}
	fmt.Println("Imagen generada en Reportes/lista_circular.png")
}

// ObtenerDato: busca un dato en la lista y lo retorna si existe
func (l *ListaCircular) ObtenerDato(dato int) *int {
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

func main() {
	lista := ListaCircular{}
	lista.Agregar(1)
	lista.Agregar(2)
	lista.Agregar(3)
	lista.Agregar(4)
	lista.Agregar(5)

	lista.Mostrar()

	lista.Graficar()

	if dato := lista.ObtenerDato(1); dato != nil {
		fmt.Println("Encontrado:", *dato)
	} else {
		fmt.Println("No encontrado")
	}
}
