package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	frequencyMap:=make(map[string]int)
	words:=strings.Fields(text)

	for _,word = range words{
		frequencyMap[word]++
	}
	return frequencyMap
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}