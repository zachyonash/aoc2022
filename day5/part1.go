package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func pop(slice []string) (string, []string) {
	var box string
	box, slice = strings.Join(slice[:1], ""), slice[1:]
	return box, slice
}

func prepend(slice []string, box string) []string {
	return append([]string{box}, slice...) // don't understand this magic quite yet
}

func moveBoxes(count, from, to string, slices [][]string) [][]string {
	convCount, _ := strconv.Atoi(count)
	convFrom, _ := strconv.Atoi(from)
	convTo, _ := strconv.Atoi(to)
	var box string
	for i := 0; i < convCount; i++ {
		box, slices[convFrom-1] = pop(slices[convFrom-1])
		slices[convTo-1] = prepend(slices[convTo-1], box)
	}
	return slices
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	count := 0
	slices := make([][]string, 9)
	for i := 0; i < 9 ; i++ {
		slices[i] = []string{}
	}
	for scanner.Scan() {
		if scanner.Text() == "" { continue }
		if scanner.Text()[:1] == " " { 
			continue
		} else if scanner.Text()[:1] == "m" {
			// follow directions, move boxes
			split := strings.Split(scanner.Text(), " ")
			slices = moveBoxes(split[1], split[3], split[5], slices)
		} else {
			for i := range make([]int, 9) {
				if strings.TrimSpace(scanner.Text()[count:count+3]) == "" {
					// Do nothing!
				} else {
					slices[i] = append(slices[i], scanner.Text()[count:count+3])
				}
				count += 4
			}
			count = 0
		}
	}
	for _, slice := range slices {
		fmt.Println(slice)
	}
}