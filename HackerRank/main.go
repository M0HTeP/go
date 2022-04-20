package main

import (
	"fmt"
)

func repeatedString(s string, n int64) int64 {
	// Write your code here
	var newStr string = ""
	var counter int64 = 0
	var y int64 = n % int64(len(s))
	var i int64 = 0
	println(n - y)
	for i < (n-y)/int64(len(s)) {
		newStr = newStr + s
		fmt.Println(newStr)
		i++
	}
	i = 0

	for i < y {
		newStr += string(s[i])
		fmt.Println(newStr)
		i++
	}
	i = 0
	for i, _ := range newStr {
		if newStr[i] == 'a' {
			counter++
		}
	}
	return counter

}

func main() {
	fmt.Println(repeatedString("aba", 10))
}

//Year 0: 7
//math.Pow(7, float64(i+1))
