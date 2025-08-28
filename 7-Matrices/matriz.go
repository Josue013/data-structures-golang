package main

import (
	"fmt"
	"os"
	"os/exec"
)

// MatrizDispersa: representa la matriz dispersa con listas de encabezados para filas y columnas
type MatrizDispersa struct {
	Filas    ListaEncabezado
	Columnas ListaEncabezado
}

// NuevaMatrizDispersa: crea e inicializa una nueva matriz dispersa
func NuevaMatrizDispersa() *MatrizDispersa {
	return &MatrizDispersa{}
}

// Insertar: inserta un nuevo nodo en la matriz dispersa
func (m *MatrizDispersa) Insertar(posX, posY int, caracter rune) {
	nuevo := &NodoInterno{X: posX, Y: posY, Valor: caracter}

	// Buscar o crear encabezado fila
	nodoX := m.Filas.Buscar(posX)
	if nodoX == nil {
		nodoX = &NodoEncabezado{ID: posX}
		m.Filas.Insertar(nodoX)
	}

	// Buscar o crear encabezado columna
	nodoY := m.Columnas.Buscar(posY)
	if nodoY == nil {
		nodoY = &NodoEncabezado{ID: posY}
		m.Columnas.Insertar(nodoY)
	}

	// INSERTAR EN FILA
	if nodoX.Acceso == nil {
		nodoX.Acceso = nuevo
	} else {
		if nuevo.Y < nodoX.Acceso.Y {
			nuevo.Derecha = nodoX.Acceso
			nodoX.Acceso.Izquierda = nuevo
			nodoX.Acceso = nuevo
		} else {
			tmp := nodoX.Acceso
			for tmp != nil {
				if nuevo.Y < tmp.Y {
					nuevo.Derecha = tmp
					nuevo.Izquierda = tmp.Izquierda
					tmp.Izquierda.Derecha = nuevo
					tmp.Izquierda = nuevo
					break
				} else if nuevo.X == tmp.X && nuevo.Y == tmp.Y {
					// Nodo repetido, no insertar
					return
				} else {
					if tmp.Derecha == nil {
						tmp.Derecha = nuevo
						nuevo.Izquierda = tmp
						break
					} else {
						tmp = tmp.Derecha
					}
				}
			}
		}
	}

	// INSERTAR EN COLUMNA
	if nodoY.Acceso == nil {
		nodoY.Acceso = nuevo
	} else {
		if nuevo.X < nodoY.Acceso.X {
			nuevo.Abajo = nodoY.Acceso
			nodoY.Acceso.Arriba = nuevo
			nodoY.Acceso = nuevo
		} else {
			tmp2 := nodoY.Acceso
			for tmp2 != nil {
				if nuevo.X < tmp2.X {
					nuevo.Abajo = tmp2
					nuevo.Arriba = tmp2.Arriba
					tmp2.Arriba.Abajo = nuevo
					tmp2.Arriba = nuevo
					break
				} else if nuevo.X == tmp2.X && nuevo.Y == tmp2.Y {
					// Nodo repetido, no insertar
					break
				} else {
					if tmp2.Abajo == nil {
						tmp2.Abajo = nuevo
						nuevo.Arriba = tmp2
						break
					} else {
						tmp2 = tmp2.Abajo
					}
				}
			}
		}
	}
}

func (m *MatrizDispersa) Graficar(nombre string) {
	contenido := `digraph G{
    node[shape=box, width=0.7, height=0.7, fontname="Arial", fillcolor="white", style=filled]
    edge[style = "bold"]
    node[label = "Root" fillcolor="darkolivegreen1" pos = "-1,1!"]raiz;`

	// Graficar nodos ENCABEZADO FILA
	pivote := m.Filas.Primero
	posX := 0
	for pivote != nil {
		contenido += fmt.Sprintf("\n\tnode[label = \"F%d\" fillcolor=\"azure3\" pos=\"-1,-%d!\" shape=box]x%d;",
			pivote.ID, posX, pivote.ID)
		pivote = pivote.Siguiente
		posX++
	}

	// Enlaces entre encabezados fila
	pivote = m.Filas.Primero
	for pivote != nil && pivote.Siguiente != nil {
		contenido += fmt.Sprintf("\n\tx%d->x%d;", pivote.ID, pivote.Siguiente.ID)
		contenido += fmt.Sprintf("\n\tx%d->x%d[dir=back];", pivote.ID, pivote.Siguiente.ID)
		pivote = pivote.Siguiente
	}
	if m.Filas.Primero != nil {
		contenido += fmt.Sprintf("\n\traiz->x%d;", m.Filas.Primero.ID)
	}

	// Graficar nodos ENCABEZADO COLUMNA
	pivoteY := m.Columnas.Primero
	posY := 0
	for pivoteY != nil {
		contenido += fmt.Sprintf("\n\tnode[label = \"C%d\" fillcolor=\"azure3\" pos = \"%d,1!\" shape=box]y%d;",
			pivoteY.ID, posY, pivoteY.ID)
		pivoteY = pivoteY.Siguiente
		posY++
	}

	// Enlaces entre encabezados columna
	pivoteY = m.Columnas.Primero
	for pivoteY != nil && pivoteY.Siguiente != nil {
		contenido += fmt.Sprintf("\n\ty%d->y%d;", pivoteY.ID, pivoteY.Siguiente.ID)
		contenido += fmt.Sprintf("\n\ty%d->y%d[dir=back];", pivoteY.ID, pivoteY.Siguiente.ID)
		pivoteY = pivoteY.Siguiente
	}
	if m.Columnas.Primero != nil {
		contenido += fmt.Sprintf("\n\traiz->y%d;", m.Columnas.Primero.ID)
	}

	// Graficar nodos internos
	pivote = m.Filas.Primero
	posX = 0
	for pivote != nil {
		pivoteCelda := pivote.Acceso
		for pivoteCelda != nil {
			// Buscar posY
			pivoteY := m.Columnas.Primero
			posYCelda := 0
			for pivoteY != nil {
				if pivoteY.ID == pivoteCelda.Y {
					break
				}
				posYCelda++
				pivoteY = pivoteY.Siguiente
			}

			if pivoteCelda.Valor == '*' {
				contenido += fmt.Sprintf("\n\tnode[label=\"*\" fillcolor=\"black\" pos=\"%d,-%d!\" shape=box]i%d_%d;",
					posYCelda, posX, pivoteCelda.X, pivoteCelda.Y)
			} else {
				contenido += fmt.Sprintf("\n\tnode[label=\" \" fillcolor=\"white\" pos=\"%d,-%d!\" shape=box]i%d_%d;",
					posYCelda, posX, pivoteCelda.X, pivoteCelda.Y)
			}
			pivoteCelda = pivoteCelda.Derecha
		}

		// Enlaces horizontales
		pivoteCelda = pivote.Acceso
		for pivoteCelda != nil {
			if pivoteCelda.Derecha != nil {
				contenido += fmt.Sprintf("\n\ti%d_%d->i%d_%d;",
					pivoteCelda.X, pivoteCelda.Y, pivoteCelda.Derecha.X, pivoteCelda.Derecha.Y)
				contenido += fmt.Sprintf("\n\ti%d_%d->i%d_%d[dir=back];",
					pivoteCelda.X, pivoteCelda.Y, pivoteCelda.Derecha.X, pivoteCelda.Derecha.Y)
			}
			pivoteCelda = pivoteCelda.Derecha
		}

		// Enlace desde encabezado fila
		if pivote.Acceso != nil {
			contenido += fmt.Sprintf("\n\tx%d->i%d_%d;", pivote.ID, pivote.Acceso.X, pivote.Acceso.Y)
			contenido += fmt.Sprintf("\n\tx%d->i%d_%d[dir=back];", pivote.ID, pivote.Acceso.X, pivote.Acceso.Y)
		}

		pivote = pivote.Siguiente
		posX++
	}

	// Enlaces verticales
	pivote = m.Columnas.Primero
	for pivote != nil {
		pivoteCelda := pivote.Acceso
		for pivoteCelda != nil {
			if pivoteCelda.Abajo != nil {
				contenido += fmt.Sprintf("\n\ti%d_%d->i%d_%d;",
					pivoteCelda.X, pivoteCelda.Y, pivoteCelda.Abajo.X, pivoteCelda.Abajo.Y)
				contenido += fmt.Sprintf("\n\ti%d_%d->i%d_%d[dir=back];",
					pivoteCelda.X, pivoteCelda.Y, pivoteCelda.Abajo.X, pivoteCelda.Abajo.Y)
			}
			pivoteCelda = pivoteCelda.Abajo
		}

		// Enlace desde encabezado columna
		if pivote.Acceso != nil {
			contenido += fmt.Sprintf("\n\ty%d->i%d_%d;", pivote.ID, pivote.Acceso.X, pivote.Acceso.Y)
			contenido += fmt.Sprintf("\n\ty%d->i%d_%d[dir=back];", pivote.ID, pivote.Acceso.X, pivote.Acceso.Y)
		}

		pivote = pivote.Siguiente
	}

	contenido += "\n}"

	// Generar archivo DOT y svg
	dotFile := fmt.Sprintf("Reportes/matriz_%s.dot", nombre)
	err := os.WriteFile(dotFile, []byte(contenido), 0644)
	if err != nil {
		fmt.Println("Error escribiendo archivo:", err)
		return
	}

	svgFile := fmt.Sprintf("Reportes/matriz_%s.svg", nombre)
	cmd := exec.Command("neato", "-Tsvg", dotFile, "-o", svgFile)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error ejecutando neato:", err)
		return
	}

	fmt.Println("Generado:", svgFile)
}
