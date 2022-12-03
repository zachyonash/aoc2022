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

func match(group []string) (value int) {
	for i, c := range group[0] {
		fmt.Println(i) // why do i have to do this?
		if strings.Contains(group[1], string(c)) && strings.Contains(group[2], string(c)) {
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
	pos := 0
	group := []string{}
	for scanner.Scan() {
		group = append(group, scanner.Text())
		pos++
		if pos == 3 {
			pos = 0
			sum += match(group)
			group = nil
		}
	}
	fmt.Println(sum)
}
