package main

import "fmt"

func main() {
	arr := []int{6, 5, 3, 4, 65, 8, 9, 10, 3, 4, 1, 34, 21, 24, 64, 13}
	//arr = []int{6, 9, 5, 4, 8, 4}
	s := quicksort(arr)
	fmt.Println(s)
}

func partition(arr []int) (int, []int, []int) {
	p := arr[len(arr)-1]
	lArr := []int{}
	rArr := []int{}
	l := 0
	r := len(arr) - 1 - 1
	for l <= r {
		if l == r {
			if arr[r] <= p {
				lArr = append(lArr, arr[r])
			} else {
				rArr = append(rArr, arr[r])
			}
			break
		}
		if arr[l] <= p {
			lArr = append(lArr, arr[l])
		}
		if arr[r] <= p {
			lArr = append(lArr, arr[r])
		}
		if arr[l] > p {
			rArr = append(rArr, arr[l])
		}
		if arr[r] > p {
			rArr = append(rArr, arr[r])
		}
		l++
		r--
	}
	return p, lArr, rArr
}

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	p, l, r := partition(arr)
	return append(append(quicksort(l), p), quicksort(r)...)
}
