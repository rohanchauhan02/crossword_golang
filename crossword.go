package main

import (
	"fmt"
	"os"
)

func display(crossword [][]string) {

	for i := 0; i < len(crossword); i++ {
		for _, j := range crossword[i] {
			if j == "+" || j == "-" {
				fmt.Print("  ")
			} else {
				fmt.Printf("%s ", string(j))

			}
		}
		fmt.Println()
	}
}

func solve(crossword [][]string, words []string, idx int) {
	if idx == len(words) {
		display(crossword)
		os.Exit(200)
	}
	word := words[idx]
	for i := 0; i < len(crossword); i++ {
		for k, j := range crossword[i] {
			if j == "-" || j == string(word[0]) {
				if canPlaceHorizontally(crossword, word, i, k) {
					wePlaced := placeHorizontally(crossword, word, i, k)
					solve(crossword, words, idx+1)
					unplacedHorizontally(crossword, wePlaced, i, k)
				}
				if canPlaceVertically(crossword, word, i, k) {
					wePlaced := placeVertically(crossword, word, i, k)
					solve(crossword, words, idx+1)
					unplacedVertically(crossword, wePlaced, i, k)
				}
			}
		}
	}
}

func canPlaceHorizontally(crossword [][]string, word string, i, j int) bool {
	if j >= 1 && crossword[i][j-1] != "+" {
		return false
	} else if j+len(word) <= len(crossword[0]) && crossword[i][j+len(word)] != "+" {
		return false
	}
	for jj := 0; jj < len(word); jj++ {
		if j+jj >= len(crossword[0]) {
			return false
		}
		if crossword[i][j+jj] == "-" || crossword[i][j+jj] == string(word[jj]) {
			continue
		} else {
			return false
		}
	}
	return true
}
func canPlaceVertically(crossword [][]string, word string, i, j int) bool {
	if i >= 1 && crossword[i-1][j] != "+" {
		return false
	} else if i+len(word) <= len(crossword) && crossword[i+len(word)][j] != "+" {
		return false
	}
	for ii := 0; ii < len(word); ii++ {
		if i+ii >= len(crossword) {
			return false
		}
		if crossword[i+ii][j] == "-" || crossword[i+ii][j] == string(word[ii]) {
			continue
		} else {
			return false
		}
	}
	return true
}

func placeHorizontally(crossword [][]string, word string, i, j int) []bool {
	wePlaced := []bool{}
	for jj := 0; jj < len(word); jj++ {
		if crossword[i][j+jj] == "-" {
			crossword[i][j+jj] = string(word[jj])
			wePlaced = append(wePlaced, true)
		} else {
			wePlaced = append(wePlaced, false)
		}
	}
	return wePlaced
}
func unplacedHorizontally(crossword [][]string, wePlaced []bool, i, j int) {
	for jj := 0; jj < len(wePlaced); jj++ {
		if wePlaced[jj] {
			crossword[i][j+jj] = "-"
		}
	}
}
func placeVertically(crossword [][]string, word string, i, j int) []bool {
	wePlaced := []bool{}
	for ii := 0; ii < len(word); ii++ {
		if crossword[i+ii][j] == "-" {
			crossword[i+ii][j] = string(word[ii])
			wePlaced = append(wePlaced, true)

		} else {
			wePlaced = append(wePlaced, false)
		}
	}
	return wePlaced
}
func unplacedVertically(crossword [][]string, wePlaced []bool, i, j int) {
	for ii := 0; ii < len(wePlaced); ii++ {
		if wePlaced[ii] {
			crossword[i+ii][j] = "-"
		}
	}
}

func main() {
	empty_crossword := [][]string{
		{"+", "-", "+", "+", "+", "+", "+", "+", "+", "+"},
		{"+", "-", "+", "+", "+", "+", "+", "+", "+", "+"},
		{"+", "-", "-", "-", "-", "-", "-", "-", "+", "+"},
		{"+", "-", "+", "+", "+", "+", "+", "+", "+", "+"},
		{"+", "-", "+", "+", "+", "+", "+", "+", "+", "+"},
		{"+", "-", "-", "-", "-", "-", "-", "+", "+", "+"},
		{"+", "-", "+", "+", "+", "-", "+", "+", "+", "+"},
		{"+", "+", "+", "+", "+", "-", "+", "+", "+", "+"},
		{"+", "+", "+", "+", "+", "-", "+", "+", "+", "+"},
		{"+", "+", "+", "+", "+", "+", "+", "+", "+", "+"}}
	word := []string{"agra", "norway", "england", "gwalior"}
	solve(empty_crossword, word, 0)
	// fmt.Println(empty_crossword)
	// display(empty_crossword)

}
