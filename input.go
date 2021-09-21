package main

import "fmt"

func main() {
  fmt.Print("Seu nome é: ")
  var input string
  fmt.Scanf("%s", &input)

  output := "Oi, " + input

  fmt.Println(output)
}