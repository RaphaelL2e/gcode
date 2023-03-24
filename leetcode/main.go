package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "first-unique-character-in-a-string"
	replaceName(name)
}

func replaceName(name string) {
	name = strings.ReplaceAll(name, "-", "_")
	fmt.Println(name)
}
