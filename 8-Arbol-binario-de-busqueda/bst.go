package main

import (
	"fmt"
	"os"
	"os/exec"
)

// ArbolBST: representa un arbol binario de búsqueda
type ArbolBST struct {
	Raiz *Nodo
}

// NuevoArbolBST: crea un nuevo arbol binario de búsqueda vacío
func NuevoArbolBST() *ArbolBST {
	return &ArbolBST{}
}

// Insertar: inserta un nuevo valor en el arbol binario de búsqueda
func (arbol *ArbolBST) Insertar(valor int) {
	arbol.Raiz = arbol.insertarRec(arbol.Raiz, valor)
}

// insertarRec: es la función recursiva que inserta un nodo en el arbol
func (arbol *ArbolBST) insertarRec(nodo *Nodo, valor int) *Nodo {
	// Si el nodo es nulo, crear un nuevo nodo
	if nodo == nil {
		return NuevoNodo(valor)
	}
	// Si el valor es menor, insertar en el subárbol izquierdo
	if valor < nodo.Valor {
		nodo.Izquierdo = arbol.insertarRec(nodo.Izquierdo, valor)
		// Si el valor es mayor, insertar en el subárbol derecho
	} else if valor > nodo.Valor {
		nodo.Derecho = arbol.insertarRec(nodo.Derecho, valor)
	}
	// Si el valor ya existe, no hacer nada (no permitir duplicados)
	return nodo
}

// Eliminar: elimina un valor del arbol binario de búsqueda
func (arbol *ArbolBST) Eliminar(valor int) {
    arbol.Raiz = arbol.eliminarRec(arbol.Raiz, valor)
}

// eliminarRec: es la función recursiva que elimina un nodo del arbol
func (arbol *ArbolBST) eliminarRec(nodo *Nodo, valor int) *Nodo {
    // Caso base: si el nodo es nil, el valor no existe
    if nodo == nil {
        return nil
    }

    // Si el valor a eliminar es menor, buscar en el subárbol izquierdo
    if valor < nodo.Valor {
        nodo.Izquierdo = arbol.eliminarRec(nodo.Izquierdo, valor)
    
    // Si el valor a eliminar es mayor, buscar en el subárbol derecho
    } else if valor > nodo.Valor {
        nodo.Derecho = arbol.eliminarRec(nodo.Derecho, valor)
    
    // Si encontramos el nodo a eliminar
    } else {
        // Caso 1: Nodo sin hijo izquierdo
        if nodo.Izquierdo == nil {
            temp := nodo.Derecho
            nodo = nil // En Go el garbage collector se encarga de la memoria
            return temp
        
        // Caso 2: Nodo sin hijo derecho
        } else if nodo.Derecho == nil {
            temp := nodo.Izquierdo
            nodo = nil // En Go el garbage collector se encarga de la memoria
            return temp
        
        // Caso 3: Nodo con dos hijos
        } else {
            // Obtener el mayor de los menores (predecesor)
            predecesor := arbol.obtenerMayorDeMenores(nodo.Izquierdo)
            // Reemplazar el valor del nodo actual con el del predecesor
            nodo.Valor = predecesor.Valor
            // Eliminar el predecesor del subárbol izquierdo
            nodo.Izquierdo = arbol.eliminarRec(nodo.Izquierdo, predecesor.Valor)
        }
    }

    return nodo
}

// obtenerMayorDeMenores: equivalente al predecesor - el nodo mas a la derecha del subarbol izquierdo
func (arbol *ArbolBST) obtenerMayorDeMenores(nodo *Nodo) *Nodo {
    actual := nodo
    // Ir hacia la derecha hasta encontrar el último nodo
    for actual.Derecho != nil {
        actual = actual.Derecho
    }
    return actual
}

// PreOrder: realiza un recorrido en preorden (raíz, izquierdo, derecho)
func (arbol *ArbolBST) PreOrder() {
	fmt.Print("PreOrder: ")
	arbol.preOrderRec(arbol.Raiz)
	fmt.Println()
}

// preOrderRec: es la función recursiva para el recorrido en preorden
func (arbol *ArbolBST) preOrderRec(nodo *Nodo) {
	if nodo != nil {
		fmt.Print(nodo.Valor, ", ")
		arbol.preOrderRec(nodo.Izquierdo)
		arbol.preOrderRec(nodo.Derecho)
	}
}

// InOrder: realiza un recorrido en orden (izquierdo, raíz, derecho)
func (arbol *ArbolBST) InOrder() {
	fmt.Print("InOrder: ")
	arbol.inOrderRec(arbol.Raiz)
	fmt.Println()
}

// inOrderRec: es la función recursiva para el recorrido en orden
func (arbol *ArbolBST) inOrderRec(nodo *Nodo) {
	if nodo != nil {
		arbol.inOrderRec(nodo.Izquierdo)
		fmt.Print(nodo.Valor, ", ")
		arbol.inOrderRec(nodo.Derecho)
	}
}

// PostOrder: realiza un recorrido en postorden (izquierdo, derecho, raíz)
func (arbol *ArbolBST) PostOrder() {
	fmt.Print("PostOrder: ")
	arbol.postOrderRec(arbol.Raiz)
	fmt.Println()
}

// postOrderRec: es la función recursiva para el recorrido en postorden
func (arbol *ArbolBST) postOrderRec(nodo *Nodo) {
	if nodo != nil {
		arbol.postOrderRec(nodo.Izquierdo)
		arbol.postOrderRec(nodo.Derecho)
		fmt.Print(nodo.Valor, ", ")
	}
}

// Graficar: genera un archivo .dot y una imagen PNG usando Graphviz
func (arbol *ArbolBST) Graficar() {
	// Crear carpeta Reportes si no existe
	os.MkdirAll("Reportes", os.ModePerm)

	contenido := "digraph G {\n"
	if arbol.Raiz != nil {
		arbol.graficarRec(arbol.Raiz, &contenido)
	}
	contenido += "}\n"

	// Escribir archivo .dot
	err := os.WriteFile("Reportes/bst.dot", []byte(contenido), 0644)
	if err != nil {
		fmt.Println("Error escribiendo archivo:", err)
		return
	}

	// Generar imagen con Graphviz
	cmd := exec.Command("dot", "-Tpng", "Reportes/bst.dot", "-o", "Reportes/bst.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando dot:", err)
		return
	}
	fmt.Println("Gráfico generado en Reportes/bst.png")
}

// graficarRec: es la función recursiva que genera el código DOT para visualizar el arbol
func (arbol *ArbolBST) graficarRec(nodo *Nodo, contenido *string) {
	if nodo != nil {
		idNodo := fmt.Sprintf("%p", nodo)
		*contenido += fmt.Sprintf("    \"%s\" [label=\"%d\"];\n", idNodo, nodo.Valor)

		if nodo.Izquierdo != nil {
			idIzquierdo := fmt.Sprintf("%p", nodo.Izquierdo)
			*contenido += fmt.Sprintf("    \"%s\" -> \"%s\";\n", idNodo, idIzquierdo)
		}

		if nodo.Derecho != nil {
			idDerecho := fmt.Sprintf("%p", nodo.Derecho)
			*contenido += fmt.Sprintf("    \"%s\" -> \"%s\";\n", idNodo, idDerecho)
		}

		arbol.graficarRec(nodo.Izquierdo, contenido)
		arbol.graficarRec(nodo.Derecho, contenido)
	}
}
