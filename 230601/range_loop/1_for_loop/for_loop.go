package main

import "fmt"

func main() {
	s := []int{0, 1, 2}

	//for range s {
	//	s = append(s, 10)
	//}

	for i := 0; i < len(s); i++ {
		s = append(s, 10)
	}
	fmt.Println(s)
}
