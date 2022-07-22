package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	// lista de nombres
	lista := "Luis,Camilo,Andres,Laura,Ernesto,Andrea"

	nombres, cant, err := separarNombres(lista)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	// respuesta
	fmt.Println("Nombres: ", nombres)
	fmt.Println("La cantidad es: ", cant)
}

func separarNombres(lista string) (nombres []string, cantidad int, err error) {
	// separar la lista por ,
	nombres = strings.Split(lista, ",")
	if nombres[0] == "" {
		return nil, 0, errors.New("La lista está vacia")
	}

	// ordenar alfabeticamente
	sort.Sort(sort.StringSlice(nombres))

	return nombres, len(nombres), nil
}
