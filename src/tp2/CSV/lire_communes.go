package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("communes_simple.csv")
	fmt.Println(string(data))
}
