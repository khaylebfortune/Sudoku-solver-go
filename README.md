# Sudoku Solver Using Backtracking in Go

## Overview

This project implements a solver for the standard 9×9 Sudoku puzzle using a recursive backtracking algorithm. Empty cells in the puzzle are represented using the `.` character. The program takes a Sudoku grid as command-line input, attempts to find a valid solution that satisfies all Sudoku constraints, and prints the solved grid if its a valid sudoku as it were, otherwise it prints error showing that the solution doesn't exist.

The implementation is written in Go and uses only the Go standard library.

---

## Mathematical Idea

Sudoku can be formulated as a constraint satisfaction problem. The puzzle consists of a 9×9 grid divided into nine 3×3 subgrids. Each cell in the grid represents a variable whose possible values are the digits 1 through 9.

The solution must satisfy the following constraints:
- Each row contains the digits 1 to 9 exactly once.
- Each column contains the digits 1 to 9 exactly once.
- Each 3×3 subgrid contains the digits 1 to 9 exactly once.

The solver searches for an assignment of digits to empty cells such that all these constraints are satisfied.

---

## Algorithm

The program uses a backtracking algorithm to solve the puzzle. The algorithm works by locating an empty cell and attempting to place each possible digit from 1 to 9 in that cell. For each attempted placement, the algorithm checks whether the digit is valid with respect to the row, column, and subgrid constraints.

If a valid digit is found, the algorithm proceeds recursively to solve the rest of the grid. If a dead end is reached, the algorithm backtracks by removing the previously placed digit and trying a different one. This process continues until either a complete solution is found or all possibilities are exhausted.

---

## How to Run
The program validates its input and prints Error if the provided grid is invalid or if no solution exists.
Clone the repository and run the program using the Go toolchain:

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"


