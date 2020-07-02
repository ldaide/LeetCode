package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	temp := make(map[int32]int)
	for _, v := range S {
		if _, ok := temp[v]; ok {
			temp[v]++
		} else {
			temp[v] = 1
		}
	}
	ret := 0
	for _, v := range J {
		if n, ok := temp[v]; ok {
			ret += n
		}
	}
	return ret
}

func main() {

	j := "aA"
	s := "aAAbbbb"
	fmt.Println(numJewelsInStones(j, s))
}
