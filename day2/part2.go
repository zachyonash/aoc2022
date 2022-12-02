package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Define outcomes
const RockDraw = "AY"
const PaperDraw = "BY"
const ScissorsDraw = "CY"
const RockWin = "CZ"
const PaperWin = "AZ"
const ScissorsWin = "BZ"
const RockLoss = "BX"
const PaperLoss = "CX"
const ScissorsLoss = "AX"
const Draw = 3
const Win = 6
const Loss = 0
const Rock = 1
const Paper = 2
const Scissors = 3

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")[0] + strings.Split(scanner.Text(), " ")[1]
		switch round {
		case RockDraw:
			sum += Rock + Draw
		case PaperDraw:
			sum += Paper + Draw
		case ScissorsDraw:
			sum += Scissors + Draw
		case RockWin:
			sum += Rock + Win
		case PaperWin:
			sum += Paper + Win
		case ScissorsWin:
			sum += Scissors + Win
		case RockLoss:
			sum += Rock + Loss
		case PaperLoss:
			sum += Paper + Loss
		case ScissorsLoss:
			sum += Scissors + Loss
		}
	}
	fmt.Println(sum)
}
