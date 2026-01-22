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

	// Comparaison
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
		return
	}

	if p1 > p2 {
		fmt.Println("Plus peuplée :", c1, "avec", p1, "habitants")
	} else if p2 > p1 {
		fmt.Println("Plus peuplée :", c2, "avec", p2, "habitants")
	} else {
		fmt.Println("Égalité :", c1, "et", c2, "ont", p1, "habitants")
	}
}
