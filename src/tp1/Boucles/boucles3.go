package main

import "fmt"

func main() {
	// Initialisation du tableau de 8 entiers
	var tailles [8]int

	// Demande à l'utilisateurs
	for i := 0; i < len(tailles); i++ {
		fmt.Printf("Entrez la taille %d en cm: ", i+1)
		fmt.Scan(&tailles[i])
	}

	// Affichage de la taille en centimètres
	for i := 0; i < len(tailles); i++ {
		t := tailles[i]
		metres := t / 100
		cm := t % 100

		fmt.Print(metres, "m")
		if cm < 10 {
			fmt.Print("0")
		}
		fmt.Println(cm)
	}
}
