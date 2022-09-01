package main

type Square struct {
	Length int
}

func (shape *Square) SquareArea() int {
	// count area of an square
	return shape.Length * 2

}

type Circle struct {
	Radius int
}

func (shape *Circle) CircleArea() int {
	return shape.Radius / 2
}

type AreaCalculator struct {
	Square []*Square
	Circle []*Circle
}

func (shapes *AreaCalculator) sum() int {
	var sum int = 0
	for _, shape := range shapes.Circle {
		sum = sum + shape.CircleArea()
	}
	for _, shape := range shapes.Square {
		sum = sum + shape.SquareArea()
	}
	return sum
}
