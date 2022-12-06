package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func pop(count int, slice []string) ([]string, []string) {
	var boxes []string
	boxes, slice = slice[len(slice)-count:], slice[:len(slice)-count]
	return boxes, slice
}

func moveBoxes(count, from, to string, slices map[string][]string) map[string][]string {
	convCount, _ := strconv.Atoi(count)
	convFrom, _ := strconv.Atoi(from)
	convTo, _ := strconv.Atoi(to)
	from = strconv.Itoa(convFrom - 1)
	to = strconv.Itoa(convTo - 1)
	var boxes []string
	if len(slices[from]) >= convCount {
		boxes, slices[from] = pop(convCount, slices[from])
		slices[to] = append(slices[to], boxes...)
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
	slices := map[string][]string{}
	stop := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if scanner.Text()[:1] == " " {
			// reverse slices
			for _, s := range slices {
				for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
					s[i], s[j] = s[j], s[i]
				}
			}
			continue
		} else if scanner.Text()[:1] == "m" {
			stop++
			// follow directions, move boxes
			split := strings.Split(scanner.Text(), " ")
			slices = moveBoxes(split[1], split[3], split[5], slices)
			// if stop > 8 {
			// 	os.Exit(0)
			// }
		} else {
			for i := range make([]int, 9) {
				if strings.TrimSpace(scanner.Text()[count:count+3]) == "" {
					// Do nothing!
				} else {
					slices[strconv.Itoa(i)] = append(slices[strconv.Itoa(i)], scanner.Text()[count:count+3])
				}
				count += 4
			}
			count = 0
		}
	}
	for i := range make([]int, 9) {
		fmt.Println(slices[strconv.Itoa(i)])
	}
}