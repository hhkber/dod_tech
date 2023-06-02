package main

import "fmt"

func main() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
		3: true,
		4: false,
	}

	m2 := m
	for k, v := range m {
		if v {
			m2[10+k] = true
		}
	}

	fmt.Println(len(m2))
	fmt.Println(m2)
}
