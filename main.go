package main

import "fmt"

// 1 add files
// 2 commit
// 3 push

func sub(a, b int) int {
	return a - b
}

func sum(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func dev(a, b float64) float64 {
	return a / b
}

func display(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(sum(5, 3))
	fmt.Println(sub(5, 3))
	fmt.Println(mul(5, 3))
	fmt.Println(dev(5, 3))
	display("Hello Gitflow")

}
