package main

import "fmt"
import _ "dod_tech_230524/230601/init/1/bbb"
import _ "dod_tech_230524/230601/init/1/aaa"

func init() {
	fmt.Println("init-init")
}

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init-1")
}

func main() {
	fmt.Println("main")
}
