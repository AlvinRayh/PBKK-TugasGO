package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan minimal 3 kata): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

	input = strings.TrimSpace(input)
	words := strings.Fields(input)
	if len(words) < 3 {
		fmt.Println("Input minimal 3 kata.")
		return
	}

	for _, word := range words {
		fmt.Println(reverse(word))
	}
}
