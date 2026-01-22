package main

import "fmt"

func main() {
	var prenom1, prenom2 string // Déclaration des variables pour les prénoms
	fmt.Print("Entrez le prenom de la première personne : ") // Demande
	fmt.Scan(&prenom1) // Lecture du premier prénom

	fmt.Print("Entrez le prenom de la deuxième personne : ") // Demande
	fmt.Scan(&prenom2) // Lecture du deuxième prénom

	var taille1, taille2 float64
	fmt.Print("Entrez la taille de la première personne : ") // Demande
	fmt.Scan(&taille1) // Lecture de la première taille

	fmt.Print("Entrez la taille de la deuxième personne : ") // Demande
	fmt.Scan(&taille2) // Lecture de la deuxième taille

	if taille1 > taille2 { // Comparaison des tailles
		fmt.Print(prenom1," est plus grand que ",prenom2, "\n")
	} else if taille2 > taille1 {
		fmt.Print(prenom2," est plus grand que ",prenom1, "\n")
	} else {
		fmt.Print(prenom1," et ",prenom2, " font la même taille.\n")
	}
}

