package main

import (
	"fmt"
	"os"
	"goblin/console"
)

func main() {
	fmt.Printf("Goblin v0.01\n")
	fmt.Println("------------")

	fmt.Println("crtl+d to quit.")
	console.Start(os.Stdin, os.Stdout)
}
