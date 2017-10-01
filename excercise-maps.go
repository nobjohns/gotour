package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCound(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCound)
}
