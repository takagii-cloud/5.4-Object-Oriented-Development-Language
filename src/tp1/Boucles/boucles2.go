package main

import "fmt"

func main() {
	tailles := [...]int{164, 208, 150, 90, 123, 156, 175, 187} // Initialisation d'un tableau d'entiers

	for i := 0; i < len(tailles); i++ { // On boucle tant que l'on ne dépasse pas la taille du tableau "tailles"
		t := tailles[i] // On ajoute les valeurs des éléments à la position [i]

		metres := t / 100 // Nombre de mètres
		cm := t % 100     // Nombre de centimètres

		fmt.Print(metres, "m") // On affiche la taille en mètre
		if cm < 10 {           // Si la taille n'est pas d'1m ex : 90cm
			fmt.Print("0")
		}
		fmt.Println(cm)
	}
}
