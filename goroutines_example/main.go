package main

import "goroutines_example/mypack"

func main() {
	a := &mypack.Cell{}

	b := new(mypack.Cell)
	print(a)
	print(b)
	print("Hello")
}
