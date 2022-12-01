package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
	"sort"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	s := []int {}
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			cur, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			sum += cur

		} else {
			s = append(s, sum)
			sum = 0
		}
	}
	sort.Ints(s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s[0] + s[1] + s[2])
}