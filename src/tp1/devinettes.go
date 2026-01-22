package main

import "fmt"

func main() {
	secret := 77   // nombre secret
	tentative := 0 // ce sera l'entrée utilisateur

	fmt.Print("Devine le nombre secret : ")
	fmt.Scan(&tentative) // lire l'entrée utilisateur

	// "for" utilisé comme un while : on recommence tant que ce n'est pas bon
	for tentative != secret {
		if tentative < secret {
			fmt.Println("Raté : le nombre secret est supérieur.")
		} else {
			fmt.Println("Raté : le nombre secret est inférieur.")
		}

		fmt.Print("Réessaie : ")
		fmt.Scan(&tentative)f
	}

	fmt.Println("Bravo ! Tu as trouvé le nombre secret.")
}
