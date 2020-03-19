package main

import "fmt"

func main() {
	result := sum(5, 5)
	fmt.Println(result)
}

func sum(v1 int, v2 int) int {
	return v1 + v2
}