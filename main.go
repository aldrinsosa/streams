package main

import "fmt"
import "flag"

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("Usage main.go FILE")
		return
	}
	fmt.Println(fileName)
}
