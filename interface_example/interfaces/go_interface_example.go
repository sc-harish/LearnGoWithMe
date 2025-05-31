package interfaces

type Shape interface {
	Area() float32
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius
}

func (r *Rectangle) Area() float32 {
	return r.Length * r.Breadth
}
