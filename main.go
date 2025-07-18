package main

import "fmt"

func main() {

	matrix := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{7, 8, 9},
	}

	fmt.Println(maximumWealth(matrix))
}

func maximumWealth(accounts [][]int) int {
	maximum := 0
	for i := range accounts {
		sum := 0
		for j := range accounts[i] {
			sum += accounts[i][j]
		}
		if sum > maximum {
			maximum = sum
		}
	}
	return maximum
}
