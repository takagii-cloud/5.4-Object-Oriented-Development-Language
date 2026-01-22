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
	// (On garde aussi la map de l'exo 1 si ton programme l'avait déjà)
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
	_ = communes // pour éviter "declared and not used communes"

	// Mappage régions (inchangé)
	regions := map[string][]string{
		"Auvergne-Rhône-Alpes": {"L'Abergement-Clémenciat", "Affoux", "Aigueperse"},
		"Occitanie":            {"Les Aires", "Alès", "Aguts", "Aiguefonde"},
		"Hauts-de-France":      {"Marseille-en-Beauvaisis", "Martincourt"},
		"Nouvelle-Aquitaine":   {"Saint-Laurent-Bretagne"},
	}

	// Demande à l'utilisateur une région
	fmt.Print("Région : ")
	reg := readLine() // Lis la réponse utiisateur

	if liste, ok := regions[reg]; ok {
		fmt.Println("Nombre de communes :", len(liste))
	} else {
		fmt.Println("Région inconnue")
	}
}
