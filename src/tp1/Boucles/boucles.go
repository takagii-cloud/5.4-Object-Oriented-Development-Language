package main

import "fmt"

func main() {
	tailles := [...]int{164, 208, 150, 90, 123, 156, 175, 187} // Initialisation d'un tableau d'entiers

	somme := 0 // Initialisation d'une variable somme à 0

	for i := 0; i < len(tailles); i++ { // On boucle tant que l'on ne dépasse pas la taille du tableau "tailles"
		somme += tailles[i] // On ajoute les valeurs des éléments à la position [i]
	}

	moyenne := float64(somme) / float64(len(tailles)) // Calcul de la moyenne

	fmt.Printf("Taille moyenne: %.2f cm\n", moyenne) // Affichage
}
