package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	tokens := strings.Fields(s)
	stat := make(map[string]int)

	for i := range tokens {
		stat[tokens[i]] += 1
	}

	return stat
}

func main() {
	wc.Test(WordCount)
}

// $ go run moretypes23.go
// PASS
//  f("I am learning Go!") =
//   map[string]int{"am":1, "learning":1, "Go!":1, "I":1}
// PASS
//  f("The quick brown fox jumped over the lazy dog.") =
//   map[string]int{"brown":1, "fox":1, "jumped":1, "dog.":1, "The":1, "quick":1, "lazy":1, "over":1, "the":1}
// PASS
//  f("I ate a donut. Then I ate another donut.") =
//   map[string]int{"I":2, "ate":2, "a":1, "donut.":2, "Then":1, "another":1}
// PASS
//  f("A man a plan a canal panama.") =
//   map[string]int{"man":1, "a":2, "plan":1, "canal":1, "panama.":1, "A":1}
