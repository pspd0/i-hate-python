package main

import "fmt"

func hate(str string) (string, string) {
	return "i hate", str
}

func main() {
	fmt.Println(hate("python"))
}
