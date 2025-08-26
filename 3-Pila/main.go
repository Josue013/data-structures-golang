package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Nodo struct {
	Dato      int
	Siguiente *Nodo
}

type Pila struct {
	Top     *Nodo
	Tamanio int
}

func (p *Pila) Push(dato int) {
	nuevo := &Nodo{Dato: dato}

	if p.Top == nil {
		p.Top = nuevo
		p.Tamanio++
	} else {
		nuevo.Siguiente = p.Top
		p.Top = nuevo
		p.Tamanio++
	}
}

func (p *Pila) Pop() (int, bool) {
	if p.Top == nil {
		return 0, false
	}

	aux := p.Top
	p.Top = p.Top.Siguiente
	p.Tamanio--
	return aux.Dato, true
}

func (p *Pila) Peek() (int, bool) {
	if p.Top == nil {
		return 0, false
	}
	return p.Top.Dato, true
}

func (p *Pila) Size() int {
	return p.Tamanio
}

func (p *Pila) IsEmpty() bool {
	if p.Tamanio == 0 {
		return true
	} else {
		return false
	}
}

func (p *Pila) Imprimir() {
	if p.Top == nil {
		fmt.Println("[]")
		return
	}

	actual := p.Top
	fmt.Print("[")
	for actual.Siguiente != nil {
		fmt.Print(actual.Dato, ", ")
		actual = actual.Siguiente
	}
	fmt.Print(actual.Dato)
	fmt.Print("]\n")
}

func (p *Pila) Graficar() {
	os.Mkdir("Reportes", os.ModePerm)

	dot := "digraph G{\n"
	dot += "	rankdir=TB;\n"
	dot += "	node [shape = record, height = 0.1]\n"

	actual := p.Top
	for i := p.Tamanio; i > 0; i-- {
		dot += fmt.Sprintf("node%d [label = \"{<f0>%d | <f1>}\"];\n", i, actual.Dato)
		actual = actual.Siguiente
	}

	actual = p.Top
	for i := p.Tamanio; i > 0; i-- {

		if actual.Siguiente == nil {
			dot += "nil_final [label = \"nil\", shape = square];\n"
			dot += fmt.Sprintf("node%d:f1 -> nil_final;\n", i)
			break
		}
		dot += fmt.Sprintf("node%d:f1 -> node%d:f0;\n", i, i-1)
		actual = actual.Siguiente
	}

	dot += "}"

	f, err := os.Create("Reportes/pila.dot")
	if err != nil {
		fmt.Println("Error creando archivo:", err)
		return
	}
	defer f.Close()
	f.WriteString(dot)

	cmd := exec.Command("dot", "-Tpng", "Reportes/pila.dot", "-o", "Reportes/pila.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando Graphviz:", err)
		return
	}
	fmt.Println("Imagen generada exitosamente")
}

func main() {
	pila := Pila{}

	fmt.Println("¿Está vacía?:", pila.IsEmpty())
	pila.Push(1)
	pila.Push(2)
	pila.Push(3)
	pila.Push(4)
	pila.Push(5)
	pila.Push(6)
	pila.Push(7)
	pila.Graficar()

	pila.Imprimir()

	fmt.Println(pila.Pop())
	pila.Imprimir()

	fmt.Println(pila.Pop())
	pila.Imprimir()

	fmt.Println(pila.Peek())
	pila.Imprimir()

	fmt.Println("El tamaño de la pila es:", pila.Size())
	fmt.Println("¿Está vacía?:", pila.IsEmpty())

	// pila.Graficar()
}
