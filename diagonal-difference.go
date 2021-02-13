package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {

	var lsum, rsum int32 = 0, 0
	var nTimes = len(arr)

	for i := 0; i < nTimes; i++ {
		for j := 0; j < nTimes; j++ {

			if i == j {
				lsum += arr[i][j]
			}

			if i+j == nTimes-1 {
				rsum += arr[i][j]
			}

		}
	}

	diff := math.Abs(float64(lsum - rsum))

	return int32(diff)
}

func main() {

	arr := [][]int32{
		{8, 9, 1},
		{4, 5, 6},
		{9, 8, 9},
	}

	result := diagonalDifference(arr)
	fmt.Println(result)
}
