# Sudoku Solver Using Backtracking in Go

## Overview
This project implements a solver for the standard 9×9 Sudoku puzzle using a recursive backtracking algorithm. Empty cells in the puzzle are represented using the `.` character.

The program accepts a Sudoku grid as command-line input, validates the input, attempts to find a solution that satisfies all Sudoku constraints, and prints the solved grid. If the input is invalid or no solution exists, the program prints `Error`.

The implementation is written in Go and uses only the Go standard library.

---

## Mathematical Idea
Sudoku can be formulated as a constraint satisfaction problem. The puzzle consists of a 9×9 grid divided into nine 3×3 subgrids. Each cell represents a variable whose possible values are the digits 1 through 9.

A valid solution must satisfy the following constraints:
- Each row contains the digits 1 to 9 exactly once.
- Each column contains the digits 1 to 9 exactly once.
- Each 3×3 subgrid contains the digits 1 to 9 exactly once.

The solver searches for an assignment of digits to empty cells such that all these constraints are satisfied simultaneously.

---

## Algorithm
The program uses a recursive backtracking algorithm to solve the puzzle. The algorithm proceeds as follows:

1. Locate the first empty cell in the grid.
2. Attempt to place each digit from 1 to 9 in the cell.
3. Check whether the placement is valid with respect to row, column, and subgrid constraints.
4. If valid, recursively attempt to solve the rest of the grid.
5. If a dead end is reached, backtrack by removing the digit and trying the next possibility.

This process continues until either a complete solution is found or all possibilities are exhausted.

---

## Input Validation and Errors
The program prints `Error` in the following cases:
- The number of input rows is not exactly 9.
- Any row does not contain exactly 9 characters.
- The input contains characters other than digits `1–9` or `.`.
- The initial grid violates Sudoku constraints.
- No valid solution exists for the given puzzle.

---

## How to Run
Clone the repository and run the program using the Go toolchain:

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
