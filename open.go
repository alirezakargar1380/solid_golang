package main

type Square struct {
	Length int
}

func (shape *Square) SquareArea() {

}

type Circle struct {
	Radius int
}

func (shape *Circle) CircleArea() {

}

type AreaCalculator struct {
	Square []*Square
	Circle []*Circle
}

func (shapes *AreaCalculator) sum() int {
	var sum int = 0
	for _, itr := range shapes.Circle {
		sum = sum + itr.Radius
	}
	for _, itr := range shapes.Square {
		sum = sum + itr.Length
	}
	return sum
}
