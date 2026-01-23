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
	// Map commune -> habitants
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

	// Map région -> liste de communes (inchangé)
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

	// Commune la plus peuplée
	maxNom := ""
	maxPop := -1

	// Nombre total de personne dans la population
	total := 0

	// Calcul du nombre total d'habitants
	for i := 0; i < len(liste); i++ { // En partant de i = 0 jusqu'à i inférieur à la longueur de la liste
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
		fmt.Println("Aucune commune trouvée.")
	} else {
		fmt.Println("Commune la plus peuplée :", maxNom)
	}

	fmt.Println("Nombre total d'habitants :", total)
}
