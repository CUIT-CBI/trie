package main

import (
	"fmt"
)

func main() {
	words := []string{"abandon", "apple", "band", "banana"}
	root := BuildTrie(words)
	fmt.Println(root)
}
