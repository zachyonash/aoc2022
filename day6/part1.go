package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func allUnique(buffer []string) bool {
	m := make(map[string]int)
	for _, c := range buffer {
		if m[c] == 0 {
			m[c] = 1
			continue
		}
		if m[c] == 1 {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	buffer := []string{}
	for scanner.Scan() {
		for i, c := range scanner.Text() {
			buffer = append(buffer, string(c))
			if len(buffer) == 4 {
				if allUnique(buffer) {
					fmt.Println(i + 1)
					os.Exit(0)
				} else {
					buffer = buffer[1:4]
				}
			}
		}
	}
}
