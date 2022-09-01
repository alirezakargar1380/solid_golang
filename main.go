package main

import "fmt"

func main() {

	// S
	// var single_response *Command = &Command{
	// 	Args: ,
	// }
	// single_response.Encode([]string{"a", "b", "c", "d"})
	// single_response.Decode()

	// O
	squ := &Square{
		Length: 4,
	}
	circle := &Circle{
		Radius: 4,
	}
	var areaCalculator *AreaCalculator = &AreaCalculator{
		Square: []*Square{squ},
		Circle: []*Circle{circle},
	}

	var sumOfShapes int = areaCalculator.sum()
	fmt.Println(sumOfShapes)
}
