package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    wordcounts := make(map[string]int)
    for _, word := range strings.Fields(s) {
        count, _ := wordcounts[word]
        wordcounts[word] = count + 1
    }
    return wordcounts
}

func main() {
    wc.Test(WordCount)
}

