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

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1) Habitants d'une commune")
		fmt.Println("2) Ajouter une commune")
		fmt.Println("3) Supprimer une commune")
		fmt.Println("4) Comparer deux communes")
		fmt.Println("0) Quitter")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		if choix == 0 {
			return
		}

		switch choix {
		case 1:
			fmt.Print("Commune : ")
			nom := readLine()
			if pop, ok := communes[nom]; ok {
				fmt.Println(pop)
			} else {
				fmt.Println("Commune inconnue")
			}

		case 2:
			fmt.Print("Nom de la commune à ajouter : ")
			nom := readLine()
			fmt.Print("Nombre d'habitants : ")
			var pop int
			fmt.Scanln(&pop)
			communes[nom] = pop
			fmt.Println("Commune ajoutée.")

		case 3:
			fmt.Print("Nom de la commune à supprimer : ")
			nom := readLine()
			if _, ok := communes[nom]; ok {
				delete(communes, nom)
				fmt.Println("Commune supprimée.")
			} else {
				fmt.Println("Commune inconnue")
			}

		case 4:
			fmt.Print("Commune 1 : ")
			c1 := readLine()
			fmt.Print("Commune 2 : ")
			c2 := readLine()

			p1, ok1 := communes[c1]
			p2, ok2 := communes[c2]

			if !ok1 || !ok2 {
				if !ok1 {
					fmt.Println("Commune inconnue :", c1)
				}
				if !ok2 {
					fmt.Println("Commune inconnue :", c2)
				}
				break
			}

			if p1 > p2 {
				fmt.Println("Plus peuplée :", c1, "avec", p1, "habitants")
			} else if p2 > p1 {
				fmt.Println("Plus peuplée :", c2, "avec", p2, "habitants")
			} else {
				fmt.Println("Égalité :", c1, "et", c2, "ont", p1, "habitants")
			}

		default:
			fmt.Println("Choix invalide")
		}
	}
}
