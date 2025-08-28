package main

import "fmt"

type Nodo struct {
	Dato      int
	Siguiente *Nodo
}

type Cola struct {
	Primero *Nodo
	Ultimo  *Nodo
	Tamanio int
}

func (c *Cola) Encolar(dato int) {
	nuevo := &Nodo{Dato: dato}

	if c.Primero == nil {
		c.Primero = nuevo
		c.Ultimo = nuevo
		c.Tamanio++
	} else {
		c.Ultimo.Siguiente = nuevo
		c.Ultimo = nuevo
		c.Tamanio++
	}
}

func (c *Cola) Desencolar() (int, bool) {
	if c.Primero == nil {
		return 0, false
	}

	dato := c.Primero.Dato
	c.Primero = c.Primero.Siguiente
	if c.Primero == nil {
		c.Ultimo = nil
	}
	c.Tamanio--
	return dato, true
}

func (c *Cola) Peek() (int, bool) {
	if c.Primero == nil {
		return 0, false
	} else {
		return c.Primero.Dato, true
	}
}

func (c *Cola) IsEmpty() bool {
	return c.Tamanio == 0
}

func (c *Cola) Imprimir() {
	if c.Primero == nil {
		fmt.Println("[]")
		return
	}

	actual := c.Primero
	fmt.Print("[")
	for actual != nil {
		if actual.Siguiente == nil {
			fmt.Print(actual.Dato, "]\n")
			return
		}
		fmt.Print(actual.Dato, ", ")
		actual = actual.Siguiente
	}
}

func main() {
	cola := &Cola{}

	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(5)
	cola.Encolar(6)
	cola.Encolar(7)
	cola.Encolar(8)
	cola.Encolar(9)
	cola.Encolar(10)
	cola.Imprimir()

	fmt.Println(cola.Desencolar())
	cola.Imprimir()
	fmt.Println(cola.Desencolar())
	cola.Imprimir()
	fmt.Println(cola.Desencolar())
	cola.Imprimir()

	dato, existe := cola.Peek()
	fmt.Printf("Primer elemento: Existe: %t, Valor: %d\n", existe, dato)

	fmt.Println("¿Está vacía?:", cola.IsEmpty())

}
