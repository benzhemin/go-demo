package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := strings.Fields(s)
	wordMap := make(map[string]int)

	for _, val := range res {
		count, ok := wordMap[val]
		if ok {
			wordMap[val] = count + 1
		} else {
			wordMap[val] = 1
		}
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
