package main

import (
	"fmt"
)

func main() {
	bf := NewBloomFilter(10)
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	check, prob := bf.Check([]byte("hello"))
	if check {
		fmt.Printf("Word exists with false positive probability of %v\n", prob)
	} else {
		fmt.Printf("Word does not exists.")
	}
	bf.PrintStats()
}
