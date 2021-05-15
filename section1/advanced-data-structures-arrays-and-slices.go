package section1

import "fmt"

func PrintArrays() {
	arr:= [5]int{1, 2, 3,4} //arrays
	var slice = []int{1,2,3}

	fmt.Println("printing from arrays")
	for x:= range arr {
		fmt.Println(x)
	}

	fmt.Println("rpinting from slice")
	for y:= range slice{
		fmt.Println(y)
	}
}
