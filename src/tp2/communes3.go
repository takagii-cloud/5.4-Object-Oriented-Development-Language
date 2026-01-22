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

	// Suppression
	fmt.Print("Nom de la commune à supprimer : ")
	nom := readLine()

	// Vérifier si la commune existe
	if _, ok := communes[nom]; ok {
		delete(communes, nom) // suppression
		fmt.Println("Commune supprimée.")
	} else {
		fmt.Println("Commune inconnue.")
	}
}
