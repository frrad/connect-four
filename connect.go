package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const (
	boardFile = "board.txt"
	rows      = 6
	columns   = 7
)

func readState() [rows][columns]int {
	board, err := ioutil.ReadFile(boardFile)
	if err != nil {
		return [rows][columns]int{}
	}

	ret := [rows][columns]int{}

	str := string(board)
	data := strings.Split(str, "\n")
	for i, line := range data {
		for j, val := range line {
			if i < 6 && j < 7 {
				if let := string(val); let == "X" {
					ret[i][j] = 1
				} else if let == "O" {
					ret[i][j] = -1
				}
			}

		}
	}

	return ret
}

func display(state [rows][columns]int) {
	for _, row := range state {
		for _, pos := range row {
			if pos == -1 {
				fmt.Print("O")
			} else if pos == 1 {
				fmt.Print("X")
			}
		}
		fmt.Print("\n")
	}

}

func main() {
	starttime := time.Now()

	state := readState()
	display(state)

	fmt.Print(state)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
