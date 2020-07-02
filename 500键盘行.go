package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	keyboard := map[int][]byte{
		1: []byte{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		2: []byte{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
		3: []byte{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
	}

	keyIndexs := map[byte]int{}

	for k, v := range keyboard {
		for _, x := range v {
			keyIndexs[x] = k
		}
	}
	res := []string{}
	for _, w := range words {
		word := strings.ToLower(w)
		flag := true
		first := keyIndexs[word[0]]
		for i := 1; i < len(word); i++ {
			cur := keyIndexs[word[i]]
			if cur != first {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, w)
		}
	}

	return res
}

func main() {
	words := []string{"zZxcvbnm"}
	fmt.Println(findWords(words))
}
