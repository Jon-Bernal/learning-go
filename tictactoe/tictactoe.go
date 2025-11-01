package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	printBoard(board)
	continueGame := true
	for continueGame {
		if checkWin(board) || !checkMovesLeft(board) {
			fmt.Println("Game Over")
			continueGame = false
		} else {
			row, col, mark := getPlayerMove(board)
			placeMark(board, row, col, mark)
			printBoard(board)
		}
	}
}

func getPlayerMove(board [][]string) (int, int, string) {
	var row, col, _ int
	var mark string
	var err error

	// row
	fmt.Println("row? 0 | 1 | 2")
	_, err = fmt.Scanf("%d", &row)

	// col
	fmt.Println("col? 0 | 1 | 2")
	_, err = fmt.Scanf("%d", &col)

	// x or o
	fmt.Println("mark? x | o")
	_, err = fmt.Scanf("%s", &mark)

	if err != nil {
		fmt.Println(err)
	}

	// col
	return row, col, mark
}

func placeMark(board [][]string, row, col int, mark string) {
	board[row][col] = mark
}

func checkMovesLeft(board [][]string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == "_" {
				return true
			}
		}
	}
	fmt.Println("")
	return false
}

func checkWin(board [][]string) bool {
	for i := 0; i < len(board); i++ {
		switch {
		case checkRowWin(board, i):
			return true
		case checkColWin(board, i):
			return true
		}
	}
	if checkDiagonals(board) {
		return true
	}
	return false
}

func checkRowWin(board [][]string, row int) bool {
	if board[row][0] != "_" && board[row][0] == board[row][1] && board[row][1] == board[row][2] {
		return true
	}
	return false
}
func checkColWin(board [][]string, col int) bool {
	if board[0][col] != "_" && board[0][col] == board[1][col] && board[1][col] == board[2][col] {
		return true
	}
	return false
}
func checkDiagonals(board [][]string) bool {
	if board[0][0] != "_" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] != "_" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	return false
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println("")
}
