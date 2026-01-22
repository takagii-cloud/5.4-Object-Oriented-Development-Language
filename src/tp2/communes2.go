package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLine() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = line[:len(line)-1]
	return line
}

func main() {
	communes := map[string]int{
		"L'Abergement-Clémenciat": 832,
		"Les Aires":               608,
		"Marseille-en-Beauvaisis": 1443,
		"Martincourt":             135,
		"Saint-Laurent-Bretagne":  445,
		"Alès":                    43892,
		"Affoux":                  395,
		"Aigueperse":              243,
		"Aguts":                   236,
		"Aiguefonde":              2510,
	}

	// Demande à l'utilisateur
	fmt.Print("Nom de la commune à ajouter : ")
	nom := readLine()

	fmt.Print("Nombre d'habitants : ")
	var pop int
	fmt.Scanln(&pop)

	communes[nom] = pop // insertion
	fmt.Println("Commune ajoutée.")
}
