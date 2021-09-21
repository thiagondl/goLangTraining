package main

import "fmt"

func average(xs []float64) float64 {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}

func f() (int, int) {
  return 5, 6
}

func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func first() {
  fmt.Println("1st")
}

func second() {
  fmt.Println("2nd")
}

func main() {
  xs := []float64{98,93,77,82,3}
  fmt.Println(average(xs))

	x, y := f()
	fmt.Println(x, y)

	fmt.Println(add(1,2,3))

	ys := []int{1,2,3}
  fmt.Println(add(ys...))

	defer second()
  first()

	defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}