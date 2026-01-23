package main

import (
	"fmt"
	"os"
	"strings"
)

func afficheCase(csv string, ligne int, colonne int) {
	lignes := strings.Split(csv, "\n")
	cols := strings.Split(lignes[ligne], ",")
	fmt.Println(cols[colonne])
}

func main() {
	data, _ := os.ReadFile("communes_simple.csv")
	csv := string(data)

	afficheCase(csv, 0, 2)
}
