package interface_example

func main() {
	print("Hello")
}

type Shape interface {
	Area() float32
}

type Circle struct {
	radius float32
}

type Rectangle struct {
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius
}
