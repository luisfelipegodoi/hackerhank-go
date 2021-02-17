package main

import (
	"errors"
	"fmt"
	"math"
)

func height(root []int) int32 {

	if len(root) == 0 {
		errors.New("Array deve ser maior que 0")
	}

	depthLeft := 4
	depthRight := 7
	depth := math.Max(float64(depthLeft), float64(depthRight)) + 1

	return int32(depth)
}

func main() {

	arrTree := []int{3, 5, 2, 1, 4, 6, 7}

	result := height(arrTree)

	fmt.Println(result)
}
