package main

import "interface_example/interfaces"

func main() {
	var s interfaces.Shape // Declare an interface variable
	s = &interfaces.Rectangle{Length: 10, Breadth: 20}

	s = &interfaces.Circle{}
	print(s.Area())
}
