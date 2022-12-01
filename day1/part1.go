package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	highest := 0
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			cur, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			sum += cur

		} else {
			if sum > highest {
				highest = sum
			}
			sum = 0
		}
	}
	fmt.Println(highest)
}
