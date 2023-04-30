package tools

import "fmt"

func PrintArgs(i int) {
	_ = fmt.Sprintf("%d", i)
}

func Success() {
	fmt.Println("Success.")
}

func HelloWorld() {
	fmt.Println("Hello World.")
}
