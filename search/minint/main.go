package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 4, 6, 8, 9, 7, 11, 10}
	arr = []int{-1, -2, -3, 1, 2, 3, 4}
	fmt.Println(searchMinInt(arr))
	//fmt.Println(1)
}

func searchMinInt(arr []int) int {
	a := make([]int, len(arr)+1)
	isAllNegative := true
	smallestInt := 0
	largestInt := 0
	for _, v := range arr {
		if v > 0 && v < len(a) {
			a[v] = 1
			isAllNegative = false
		}
		if largestInt < v {
			largestInt = v
		}
	}
	if isAllNegative {
		return 1
	}
	for i, v := range a {
		if i == 0 {
			continue
		}
		if v == 0 {
			smallestInt = i
			break
		}
	}
	if smallestInt == 0 {
		return largestInt + 1
	} else {
		return smallestInt
	}
}
