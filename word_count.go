package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sp:= strings.Split(s," ")
	for _,c := range sp{		
		v, ok := m[c]
		if ok{
			v=v+1
			m[c]=v
		}else{
		m[c]=1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
