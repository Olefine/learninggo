package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Sym   string
	Turns map[int]int
}

type Board [][]string

func (p *Player) Turn(x, y int, board *Board) {
	if x > 2 && x < 0 || y > 2 && y < 0 {
		panic("board is 3 x 3")
	}

	(*board)[x][y] = p.Sym
}

func main() {
	// Create a tic-tac-toe board.
	board := Board{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	player1 := Player{Sym: "X"}
	player2 := Player{Sym: "O"}

	scanner := bufio.NewScanner(os.Stdin)

	displayBoard(board)

	for {
		for scanner.Scan() {
			scannerText := scanner.Text()
			//Pass data if form "x y PlayerSymbol" e.g "0 0 X"
			if scannerText == "Q" {
				fmt.Println("Leaving game")
				return
			}

			parsedInput := strings.Split(scannerText, " ")

			fmt.Println(parsedInput)

			x, _ := strconv.Atoi(parsedInput[0])
			y, _ := strconv.Atoi(parsedInput[1])
			sym := parsedInput[2]

			if sym == "X" {
				player1.Turn(x, y, &board)
			} else {
				player2.Turn(x, y, &board)
			}

			displayBoard(board)

			if checkWinnerBySymbol(board, sym) {
				fmt.Printf("Player %s won the game\n", sym)
				os.Exit(0)
			}
		}
	}
}

func displayBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("")
}

func checkWinnerBySymbol(board [][]string, symbol string) bool {
	repitedSym := fmt.Sprintf("%[1]s%[1]s%[1]s", symbol)
	return checkRows(board, repitedSym) || checkRows(transpose(board), repitedSym) || checkDiagolans(board, repitedSym)
}

func transpose(b [][]string) [][]string {
	length := len(b[0])
	res := make([][]string, length)
	copy(res, b)
	for i := 0; i < length; i++ {
		res[i] = []string{b[0][i], b[1][i], b[2][i]}
	}

	return res
}

func checkRows(board [][]string, signature string) bool {
	return strings.Join(board[0], "") == signature || strings.Join(board[1], "") == signature || strings.Join(board[2], "") == signature
}

func checkDiagolans(board [][]string, signature string) bool {
	firstDiagonal := []string{board[0][0], board[1][1], board[2][2]}
	secondDiagonal := []string{board[0][2], board[1][1], board[2][0]}
	return strings.Join(firstDiagonal, "") == signature || strings.Join(secondDiagonal, "") == signature
}
