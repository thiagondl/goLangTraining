package main

import "fmt"

var x string = "Hello World"

func main() {
  x := "outra coisa"
  fmt.Println(x)
	f()
}

func f() {
  fmt.Println(x)
}