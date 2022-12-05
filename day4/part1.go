package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func contains(elf []int, s int) bool {
	for _, i := range elf {
		if i == s {
			return true
		}
	}
	return false
}

func fullyContained(elf1Lo, elf1Hi, elf2Lo, elf2Hi int) int {
	elf1contains := true
	elf2contains := true
	elf1arr := make([]int, elf1Hi-elf1Lo+1)
	for i := range elf1arr {
		elf1arr[i] = i + elf1Lo
	}
	elf2arr := make([]int, elf2Hi-elf2Lo+1)
	for i := range elf2arr {
		elf2arr[i] = i + elf2Lo
	}

	for _, i := range elf1arr {
		if !contains(elf2arr, i) {
			elf1contains = false
			break
		}
	}
	if elf1contains { return 1 }
	for _, i := range elf2arr {
		if !contains(elf1arr, i) {
			elf2contains = false
			break
		}
	}
	if elf2contains { return 1 }
	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), ",")
		elf1 := strings.Split(temp[0], "-")
		elf2 := strings.Split(temp[1], "-")
		elf1Lo, _ := strconv.Atoi(elf1[0])
		elf1Hi, _ := strconv.Atoi(elf1[1])
		elf2Lo, _ := strconv.Atoi(elf2[0])
		elf2Hi, _ := strconv.Atoi(elf2[1])
		count += fullyContained(elf1Lo, elf1Hi, elf2Lo, elf2Hi)
	}
	fmt.Println(count)
}
