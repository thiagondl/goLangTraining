package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	x["delete"] = 1
	fmt.Println(x)
	delete(x, "delete")
	fmt.Println(x)

	elements := map[string]map[string]string{
    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
	}
	fmt.Println(elements)
}