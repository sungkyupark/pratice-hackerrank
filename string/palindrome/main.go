package main

import (
	"fmt"
)

func main() {
	s := getLargestPalindrome("adddcabbacdddbxxxxxdclkjddjklcdxxxxx")
	fmt.Println(len(s))
}

func getLargestPalindrome(s string) string {
	var largest, fStr1, fStr2, comp string
	for i := 0; i < len(s); i++ {
		var l = i - 1
		var r = i + 1
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			if fStr1 == "" {
				fStr1 = string(s[i])
			}
			fStr1 = string(s[l]) + fStr1 + string(s[r])
			l--
			r++
		}
		l = i
		r = i + 1
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			fStr2 = string(s[l]) + fStr2 + string(s[r])
			l--
			r++
		}
		comp = fStr1
		if len(fStr1) < len(fStr2) {
			comp = fStr2
		}
		if len(largest) < len(comp) {
			largest = comp
		}
		fStr1 = ""
		fStr2 = ""
	}
	return largest
}
