package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Printf("index=%d, value=%s\n", i, v)
	}
}
