package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	matriz := NuevaMatrizDispersa()

	// matriz.Insertar(10, 10, '*')
	// matriz.Graficar("PrimerNodo")
	// matriz.Insertar(1, 1, '*')
	// matriz.Insertar(1, 2, '*')
	// matriz.Insertar(1, 3, '*')
	// matriz.Insertar(1, 4, '*')
	// matriz.Insertar(2, 1, '*')
	// matriz.Insertar(2, 2, '*')
	// matriz.Insertar(2, 3, '*')
	// matriz.Insertar(2, 4, '*')
	// matriz.Insertar(3, 1, '*')
	// matriz.Insertar(3, 2, '*')
	// matriz.Insertar(3, 3, '*')
	// matriz.Insertar(3, 4, '*')
	// matriz.Insertar(4, 1, '*')
	// matriz.Insertar(4, 2, '*')
	// matriz.Insertar(4, 3, '*')
	// matriz.Insertar(4, 4, '*')
	// matriz.Insertar(8, 9, '*')
	// matriz.Insertar(9, 8, '*')
	// matriz.Graficar("Final")

	// Equivalente a insertaTodo()
	// insertaTodo(matriz)
	// Equivalente a insertaSeleccion()
	insertaSeleccion(matriz)

}

func insertaTodo(matriz *MatrizDispersa) {
    file, err := os.Open("ArchivosDeEntrada/carita.txt")
    if err != nil {
        fmt.Println("err abriendo archivo:", err)
        return
    }
    defer file.Close()

    lectorLineas := bufio.NewScanner(file)
    numeroFila := 0
    for lectorLineas.Scan() {
        textoLinea := lectorLineas.Text()
        numeroFila++
        numeroColumna := 0
        for _, caracter := range textoLinea {
            if caracter != '\n' {
                numeroColumna++
                matriz.Insertar(numeroFila, numeroColumna, caracter)
            }
        }
    }
    matriz.Graficar("carita")
}

func insertaSeleccion(matriz *MatrizDispersa) {
    file, err := os.Open("ArchivosDeEntrada/gengar.txt")
    if err != nil {
        fmt.Println("err abriendo archivo:", err)
        return
    }
    defer file.Close()

    lectorLineas := bufio.NewScanner(file)
    numeroFila := 0
    for lectorLineas.Scan() {
        textoLinea := lectorLineas.Text()
        numeroFila++
        numeroColumna := 0
        for _, caracter := range textoLinea {
            if caracter != '\n' {
                numeroColumna++
                if caracter == '*' {
                    matriz.Insertar(numeroFila, numeroColumna, caracter)
                }
            }
        }
    }
    matriz.Graficar("gengar")
}
