package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const LowerDelta = 96
const UpperDelta = 38

func match(first, second string) (value int) {
	for i, c := range first {
		fmt.Println(i) // why do i have to do this?
		if strings.Contains(second, string(c)) {
			fmt.Printf("Similar letter is %s\n", string(c))
			if unicode.IsLower(c) {
				return int(c) - LowerDelta
			} else {
				return int(c) - UpperDelta
			}
		}
	}
	return 1
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		len := len(scanner.Text())
		sum += match(scanner.Text()[:len/2], scanner.Text()[len/2:])
	}
	fmt.Println(sum)
}
