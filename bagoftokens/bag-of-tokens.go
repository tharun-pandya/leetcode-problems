package main

import (
	"fmt"
)

func main() {
	var tokens []int
	var len, power, tdata int
	fmt.Println("Enter length of tokens")
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Print("Enter ", i, "token data :")
		fmt.Scan(&tdata)
		tokens = append(tokens, tdata)
	}
	fmt.Println("Enter power")
	fmt.Scan(&power)
	fmt.Println(bagOfTokensScore(tokens, power))
}
func bagOfTokensScore(tokens []int, power int) int {
	var score int
	var array []int
	for i := 0; i < len(tokens); i++ {
		for j := 0; j < len(tokens); j++ {
			if tokens[i] < tokens[j] {
				ele := tokens[i]
				tokens[i] = tokens[j]
				tokens[j] = ele
			}
		}
	}
	if len(tokens) == 0 {
		return 0
	} else if tokens[0] > power {
		return 0
	} else {
		power = power - tokens[0]
		score++ //face up
		array = append(array, score)
		len := len(tokens) - 1
		for i := 1; i <= len; i++ {
			if tokens[i] > power {
				power = tokens[len] + power
				score--
				array = append(array, score)
				len-- //face down
				i--
			} else {
				power = power - tokens[i]
				score++ //face up
				array = append(array, score)
			}
		}
	}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] < array[j] {
				score = array[j]
			}
		}
	}
	return score
}
