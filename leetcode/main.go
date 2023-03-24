package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("please input name:")
	var name string
	_, _ = fmt.Scanf("%s", &name)
	replaceName(name)
}

func replaceName(name string) {
	name = strings.ReplaceAll(name, "-", "_")
	fmt.Println(name)
}
