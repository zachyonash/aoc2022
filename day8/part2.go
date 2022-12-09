package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	treesGrid := make([][]rune, 0)
	pos := 0
	for scanner.Scan() {
		treesRow := make([]rune, 0)
		for _, c := range scanner.Text() {
			treesRow = append(treesRow, c)
		}
		treesGrid = append(treesGrid, treesRow)
		pos++
	}

	// Count number of visible trees
	scenicScore := 0
	highScore := 0
	for rowPos, row := range treesGrid {
		for columnPos, tree := range row {
			if rowPos == 0 || rowPos == len(treesGrid)-1 || columnPos == 0 || columnPos == len(row)-1 {
				continue
			} else {
				scenicScore = score(treesGrid, row, int(tree-'0'), rowPos, columnPos) 
				if scenicScore > highScore { highScore = scenicScore }
			}
		}
	}
	fmt.Println(highScore)
}

func score(treesGrid [][]rune, row []rune, tree int, rowPos int, columnPos int) int {
	directions := []string{"left", "right", "up", "down"}
	leftScore, rightScore, upScore, downScore := 0, 0, 0, 0
	for _, direction := range directions {
		switch direction {
		case "left":
			index := columnPos
			for index > 0 {
				if tree <= int(row[index-1]-'0') {
					leftScore++
					break
				}
				leftScore++
				index--
			}
		case "right":
			index := columnPos
			for index < len(row)-1 {
				if tree <= int(row[index+1]-'0') {
					rightScore++
					break
				}
				rightScore++
				index++
			}
		case "up":
			index := rowPos
			for index > 0 {
				if tree <= int(treesGrid[index-1][columnPos]-'0') {
					upScore++
					break
				}
				upScore++
				index--
			}
		case "down":
			index := rowPos
			for index < len(treesGrid)-1 {
				if tree <= int(treesGrid[index+1][columnPos]-'0') {
					downScore++
					break
				}
				downScore++
				index++
			}
		}
	}
	return leftScore * rightScore * upScore * downScore
}
