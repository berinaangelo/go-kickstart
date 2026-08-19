// Write a program that solves an open-ended knight's tour given an (N x N) matrix, such that  100 > N > 3

package main

import (
	"fmt"
	"math/rand"
)

// todo: find way how to parameterize this
const N int = 8

var MOVES = struct {
	x []int
	y []int
}{
	x: []int{2, 1, -1, -2, -2, -1, 1, 2},
	y: []int{1, 2, 2, 1, -1, -2, -2, -1},
}

func isMoveValid(x int, y int, board [][]int) bool {
	return (0 <= x && x < N) && (0 <= y && y < N) && board[x][y] == -1
}

func getDegree(x int, y int, board [][]int) int {
	count := 0

	// changed to MOVES.x since its a slice and 0-based
	for idx := range MOVES.x {
		if isMoveValid(x+MOVES.x[idx], y+MOVES.y[idx], board) {
			count += 1
		}
	}

	return count
}

func nextMove(board [][]int, posX int, posY int) *[]int {
	minDegIdx := -1
	minDegree := N + 1

	startIndex := rand.Intn(N)

	for i := range MOVES.x {
		idx := (startIndex + i) % N
		nextX, nextY := posX+MOVES.x[idx], posY+MOVES.y[idx]

		if isMoveValid(nextX, nextY, board) {
			degree := getDegree(nextX, nextY, board)

			if degree < minDegree {
				minDegree, minDegIdx = degree, idx
			}
		}
	}

	if minDegIdx == -1 {
		return nil
	}

	return &[]int{(posX + MOVES.x[minDegIdx]), (posY + MOVES.y[minDegIdx])}
}

func createBoard(N int) [][]int {
	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
		for j := range board[i] {
			board[i][j] = -1
		}
	}

	return board
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%2d ", cell)
		}
		fmt.Println()
	}
}

func solveKnightsTour() bool {
	if N < 3 && N > 100 {
		fmt.Println("invalid board size")
		return false
	}

	// initialize the NxN grid
	board := createBoard(N)
	posX, posY := rand.Intn(N-1), rand.Intn(N-1)

	for i := 2; i <= N*N+1; i++ {
		step := nextMove(board, posX, posY)

		// TODO: find workaround on how to deal with tie breakers
		if step == nil {
			fmt.Printf("Tour failed at step %v. Backtracking needed or bad tie-break occurred.", i)
			return false
		}

		posX, posY = (*step)[0], (*step)[1]
		board[posX][posY] = i
	}

	printBoard(board)
	return true
}
