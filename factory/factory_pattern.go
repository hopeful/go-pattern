package factory

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Circle struct {
	Name string
}

func (c *Circle) Draw() {
	fmt.Printf("我是:%s \n", c.Name)
}

type Square struct {
	Name string
}

func (s *Square) Draw() {
	fmt.Printf("我是:%s \n", s.Name)
}

type Rectangle struct {
	Name string
}

func (r *Rectangle) Draw() {
	fmt.Printf("我是:%s \n", r.Name)
}

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "Circle":
		return &Circle{Name: "Circle"}
	case "Square":
		return &Square{Name: "Square"}
	case "Rectangle":
		return &Rectangle{Name: "Rectangle"}
	}
	return nil
}
