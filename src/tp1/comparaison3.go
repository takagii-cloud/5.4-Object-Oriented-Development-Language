package main

import "fmt"

func main() {
	var prenom1, prenom2 string                              // Déclaration des variables pour les prénoms
	fmt.Print("Entrez le prenom de la première personne : ") // Demande du prenom de la première personne
	fmt.Scan(&prenom1)                                       // Lecture du prénom de la première personne

	var genre1, genre2 string                                                               // Déclaration des variables pour le genre
	fmt.Print("Quel est le genre de ", prenom1, " ? M pour masculin, et F pour féminin : ") // Demande le genre de la première personne
	fmt.Scan(&genre1)                                                                       // Lecture du genre de la première personne

	var taille1, taille2 float64
	fmt.Print("Entrez la taille de la première personne : ") // Demande de la taille de la première personne
	fmt.Scan(&taille1)                                       // Lecture de la première taille

	fmt.Print("Entrez le prenom de la deuxième personne : ") // Demande du prenom de de la deuxième personne
	fmt.Scan(&prenom2)                                       // Lecture du deuxième prénom

	fmt.Print("Quel est le genre de ", prenom2, " ? M pour masculin, F pour féminin : ") // Demande le genre de la deuxième personne
	fmt.Scan(&genre2)                                                                    // Lecture du genre de la deuxième personne

	fmt.Print("Entrez la taille de la deuxième personne : ") // Demande de la taille de la deuxième personne
	fmt.Scan(&taille2)                                       // Lecture de la deuxième taille

	var langue string                                                                                           // Déclation de la variable langue
	fmt.Print("Dans quelle langue souhaitez-vous afficher le message ? FR pour Français et EN pour Anglais : ") // Demande la langue à l'utilisateur
	fmt.Scan(&langue)                                                                                           // Lecture de la langue

	var phrase string // plus grand ou plus grande
	if genre1 == "M" && taille1 > taille2 && langue == "FR" {
		phrase = " est plus grand que "
	} else if genre1 == "M" && taille1 > taille2 && langue == "EN" {
		phrase = " is taller than "
	} else if genre2 == "M" && taille2 > taille1 && langue == "FR" {
		phrase = (" est plus grand que ")
	} else if genre2 == "M" && taille2 > taille1 && langue == "EN" {
		phrase = (" is taller than ")
	} else if genre1 == "F" && taille1 > taille2 && langue == "FR" {
		phrase = (" est plus grande que ")
	} else if genre1 == "F" && taille1 > taille2 && langue == "EN" {
		phrase = (" is taller than ")
	} else if genre2 == "F" && taille2 > taille1 {
		phrase = (" est plus grande que ")
	} else if genre2 == "F" && taille2 > taille1 && langue == "FR" {
		phrase = (" est plus grande que ")
	} else if genre2 == "F" && taille2 > taille1 && langue == "EN" {
		phrase = (" is taller than ")
	} else if taille1 == taille2 && langue == "FR" {
		phrase = ("font la même taille ")
	} else {
		phrase = (" are the same  height ")
	}

	if taille1 > taille2 { // Comparaison des tailles
		fmt.Print(prenom1, phrase, prenom2, "\n")
	} else if taille2 > taille1 {
		fmt.Print(prenom2, phrase, prenom1, "\n")
	} else {
		fmt.Print(prenom1, " et ", prenom2, phrase, "\n")
	}
}
