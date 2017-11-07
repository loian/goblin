package main

import (
	"fmt"
	"os"
	"goblin/console"
)

func main() {
	fmt.Printf("Goblin v0.01\n")
	console.Start(os.Stdin, os.Stdout)
}
