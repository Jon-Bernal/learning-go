// Exercise: Maps

// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

// You might find strings.Fields helpful.

package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	//fmt.Println("arr: ", arr)
	var sMap = make(map[string]int)
	
	for i := 0; i < len(arr); i++ {
		sMap[arr[i]] = sMap[arr[i]] + 1
	}
		//fmt.Println("sMap: ", sMap)
	return sMap
}

func main() {
	wc.Test(WordCount)
}
