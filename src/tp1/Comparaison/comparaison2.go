package main

import "fmt"

func main() {
	var prenom1, prenom2 string                              // Déclaration des variables pour les prénoms
	fmt.Print("Entrez le prenom de la première personne : ") // Demande
	fmt.Scan(&prenom1)                                       // Lecture du premier prénom

	var genre1, genre2 string                                                               // Déclaration des variables pour le genre
	fmt.Print("Quel est le genre de ", prenom1, " ? M pour masculin, et F pour féminin : ") // Demande le genre
	fmt.Scan(&genre1)                                                                       // Lecture du genre de la première personne

	var taille1, taille2 float64
	fmt.Print("Entrez la taille de la première personne : ") // Demande
	fmt.Scan(&taille1)                                       // Lecture de la première taille

	fmt.Print("Entrez le prenom de la deuxième personne : ") // Demandes
	fmt.Scan(&prenom2)                                       // Lecture du deuxième prénom

	fmt.Print("Quel est le genre de ", prenom2, " ? M pour masculin, F pour féminin : ") // Demande
	fmt.Scan(&genre2)                                                                    // Lecture du genre de la deuxième personne

	fmt.Print("Entrez la taille de la deuxième personne : ") // Demande
	fmt.Scan(&taille2)                                       // Lecture de la deuxième taille

	var phrase string // plus grand ou plus grande
	if genre1 == "M" && taille1 > taille2 {
		phrase = " est plus grand que "
	} else if genre2 == "M" && taille2 > taille1 {
		phrase = (" est plus grand que ")
	} else if genre1 == "F" && taille1 > taille2 {
		phrase = (" est plus grande que ")
	} else if genre2 == "F" && taille2 > taille1 {
		phrase = (" est plus grande que ")
	} else {
		phrase = ("font la même taille ")
	}

	if taille1 > taille2 { // Comparaison des tailles
		fmt.Print(prenom1, phrase, prenom2, "\n")
	} else if taille2 > taille1 {
		fmt.Print(prenom2, phrase, prenom1, "\n")
	} else {
		fmt.Print(prenom1, "et", prenom2, phrase, "\n")
	}
}
