package main

import "fmt"
var x  uint8 = 3
func main()  {
	fmt.Println(x)
	fmt.Println(compareNumbers(3, 3))
	//type inference
	x:=2
	fmt.Println(x)
}



func compareNumbers(i1, i2 int) (bool,int) {
	if i1 > i2 {
		fmt.Println("first number is greater than second number")
		return false, i1 - i2
	} else if i1 < i2 {
		fmt.Println("second number is greater than first number")
		return false, i2-i1
	}
	fmt.Println("numbers are equal")
	return true, 0
}

