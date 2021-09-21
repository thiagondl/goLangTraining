package main

import "fmt"

func main() {
	x := make([]float64, 5)
	x[0] = 1
  fmt.Println(x)

	arr := [5]float64{1,2,3,4,5}
	y := arr[0:5]
	arr[0] = 2
	fmt.Println(y)

	k := make([]float64, 5)
	w := append(k, 5)
	fmt.Println(w)

	slice1 := []int{1,2,3}
  slice2 := make([]int, 2)
  copy(slice2, slice1)
  fmt.Println(slice1, slice2)
}