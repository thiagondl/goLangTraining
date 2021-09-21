package main

import "fmt"

func main() {
	x := [5]int{ 98, 93, 77, 100, 83 }
  fmt.Println(x)

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += float64(x[i])
	}
	fmt.Println(total / float64(len(x)))
}