package main

import ("fmt"; "math")

type Rectangle struct {
  x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

type Person struct {
  Name string
}
func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
  Person
  Model string
}

func main() {
	r := Rectangle{0, 0, 10, 10} // var r Rectangle | r := new(Rectangle)
	fmt.Println(r.area())
	a := new(Android)
	a.Name = "Thiago"
	a.Talk()
}