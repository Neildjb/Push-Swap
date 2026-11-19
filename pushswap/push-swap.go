package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mots := []string{}
	liste_nombres := []int{}

	for _, elt := range os.Args[1:] {
		mots = append(mots, strings.Fields(elt)...)
	}

	for _, mot := range mots {
		nombre, erreur := strconv.Atoi(mot)
		if erreur != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		} else {
			for _, anciennombre := range liste_nombres {
				if anciennombre == nombre {
					fmt.Fprintln(os.Stderr, "Error")
					os.Exit(1)
				}
			}
			liste_nombres = append(liste_nombres, nombre)
		}
	}
	fmt.Println(liste_nombres)
}
