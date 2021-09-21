package main

import "fmt"

func main() {
  i := 1
  for i <= 10 { //while
    fmt.Println(i)
    i = i + 1
  }

	i = 0
  for i <= 10 { // while switch
    switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
		}
    i = i + 1
  }

	for i := 1; i <= 10; i++ { //for index if
		if i % 2 == 0 {
			fmt.Println("2", i)
		} else if i % 3 == 0 {
			fmt.Println("3", i)
		}
  }
	x := [5]float64{ 98, 93, 77, 82, 83 }
	var total float64 = 0
	for _, value := range x { //for each (i, e)
		total += float64(value)
	}
	fmt.Println(total / float64(len(x)))

}