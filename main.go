package main

import (
	"fmt"
	"os"
)

func printError() {
	fmt.Println("Error")
	os.Exit(0)
}

func validInput(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for _, row := range args {
		if len(row) != 9 {
			return false
		}
		for i := 0; i < 9; i++ {
			c := row[i]
			if c != '.' && (c < '1' || c > '9') {
				return false
			}
		}
	}
	return true
}

func buildGrid(args []string) [9][9]byte {
	var g [9][9]byte
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			g[r][c] = args[r][c]
		}
	}
	return g
}

func initialValid(g *[9][9]byte) bool {
	var seenRow, seenCol, seenBox [9][10]bool
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c] == '.' {
				continue
			}
			d := int(g[r][c] - '0')
			if seenRow[r][d] || seenCol[c][d] {
				return false
			}
			seenRow[r][d] = true
			seenCol[c][d] = true

			box := (r/3)*3 + (c / 3)
			if seenBox[box][d] {
				return false
			}
			seenBox[box][d] = true
		}
	}
	return true
}

func canPlace(g *[9][9]byte, r, c int, digit byte) bool {
	for i := 0; i < 9; i++ {
		if g[r][i] == digit || g[i][c] == digit {
			return false
		}
	}
	br := (r / 3) * 3
	bc := (c / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g[br+i][bc+j] == digit {
				return false
			}
		}
	}
	return true
}

func solve(g *[9][9]byte) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c] == '.' {
				for d := byte('1'); d <= '9'; d++ {
					if canPlace(g, r, c, d) {
						g[r][c] = d
						if solve(g) {
							return true
						}
						g[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func printGrid(g *[9][9]byte) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if c > 0 {
				fmt.Print(" ") // add space between numbers
			}
			fmt.Printf("%c", g[r][c])
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]
	if !validInput(args) {
		printError()
	}
	grid := buildGrid(args)
	if !initialValid(&grid) {
		printError()
	}
	if !solve(&grid) {
		printError()
	}
	printGrid(&grid)
}
