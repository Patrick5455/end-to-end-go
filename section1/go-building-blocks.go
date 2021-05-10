package section1

import "fmt"
var x  uint8 = 3
func main()  {
	fmt.Println(x)
	fmt.Println(compareNumbers2(3, 3))
	//type inference
	x:=2
	switch x {
	case 3:
		fmt.Println("I am 3")
	case 2:
		fmt.Println("I am 2")
	case 4:
		fmt.Println("I am 4")
	}

	switch x := 4; x {
	case 3:
		fmt.Println("I am 3")
	case 2:
		fmt.Println("I am 2")
	case 4:
		fmt.Println("I am 4")
	}

	forLoop()
	whileLoopBasedOnForLoop()
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

// conditional statements and loops

// switch statements

func compareNumbers2(t1, t2 int) (bool, int)  {
// switch statement in go unlike java does not need a break clause
	switch {
	case t1 > t2:
		fmt.Println("first number is greater than second number")
		return false, t1 - t2
	case t2 > t1:
		fmt.Println("second number is greater than first number")
		return false, t2- t1
	}
	fmt.Println("numbers are equal")
	return true, t1 - t1
}

//loops : the only keyword used for loops in go is for

func forLoop()  {
	for x:=0; x<5; x++ {
		fmt.Println("values of x is ",x)
	}
}

func whileLoopBasedOnForLoop()  {
	i:=1
	fmt.Println("increment")
	for i<=10 {
		fmt.Print(i)
		i++
	}

	fmt.Println("\ndecrement")
	for i:=10; i>=1; {
		i--
		fmt.Print(i)
	}
}

