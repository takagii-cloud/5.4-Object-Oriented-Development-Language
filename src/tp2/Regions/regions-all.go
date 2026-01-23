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
	// Mapage Commune -> Nb Habitants
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

	// Mappage région -> Commune
	regions := map[string][]string{
		"Auvergne-Rhône-Alpes": {"L'Abergement-Clémenciat", "Affoux", "Aigueperse"},
		"Occitanie":            {"Les Aires", "Alès", "Aguts", "Aiguefonde"},
		"Hauts-de-France":      {"Marseille-en-Beauvaisis", "Martincourt"},
		"Nouvelle-Aquitaine":   {"Saint-Laurent-Bretagne"},
	}

	fmt.Print("Région : ")
	reg := readLine()

	liste, ok := regions[reg]
	if !ok {
		fmt.Println("Région inconnue")
		return
	}

	// Nb Habitants dans une communce
	fmt.Println("Nombre de communes :", len(liste))

	// Commune la plus peuplée
	maxNom := ""
	maxPop := -1

	// Total d'habitants
	total := 0

	for i := 0; i < len(liste); i++ {
		nom := liste[i]
		pop, exists := communes[nom]
		if exists {
			total = total + pop

			if pop > maxPop {
				maxPop = pop
				maxNom = nom
			}
		}
	}

	if maxNom == "" {
		fmt.Println("Commune la plus peuplée : aucune")
	} else {
		fmt.Println("Commune la plus peuplée :", maxNom)
	}

	fmt.Println("Nombre total d'habitants :", total)
}
