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
	count := 0
	for rowPos, row := range treesGrid {
		for columnPos, tree := range row {
			if rowPos == 0 || rowPos == len(treesGrid)-1 || columnPos == 0 || columnPos == len(row)-1 {
				count++
			} else {
				if visible(treesGrid, row, int(tree-'0'), rowPos, columnPos) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func visible(treesGrid [][]rune, row []rune, tree int, rowPos int, columnPos int) bool {
	directions := []string{"left", "right", "up", "down"}
	for _, direction := range directions {
		switch direction {
		case "left":
			leftVisible := true
			index := columnPos
			for index > 0 {
				if tree <= int(row[index-1]-'0') {
					leftVisible = false
				}
				index--
			}
			if leftVisible {
				return true
			}
		case "right":
			rightVisible := true
			index := columnPos
			for index < len(row)-1 {
				if tree <= int(row[index+1]-'0') {
					rightVisible = false
				}
				index++
			}
			if rightVisible {
				return true
			}
		case "up":
			upVisible := true
			index := rowPos
			for index > 0 {
				if tree <= int(treesGrid[index-1][columnPos]-'0') {
					upVisible = false
				}
				index--
			}
			if upVisible {
				return true
			}
		case "down":
			downVisible := true
			index := rowPos
			for index < len(treesGrid)-1 {
				if tree <= int(treesGrid[index+1][columnPos]-'0') {
					downVisible = false
				}
				index++
			}
			if downVisible {
				return true
			}
		}
	}
	return false
}
