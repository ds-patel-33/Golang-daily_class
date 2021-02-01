package main

import "fmt"

func main() {
	var y [5]string
	fmt.Println(y)

	var z [5]complex128
	fmt.Println(z)

	var x [5]int

	x[0] = 100
	x[1] = 101

	fmt.Printf("x[0] = %d, x[1] = %d\n", x[0], x[1])
	fmt.Println("x = ", x)

	//declare and initialize
	var a = [5]int{2, 4, 6, 8, 10}
	fmt.Println("a = ", a)

	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1

	a2[1] = "German"

	fmt.Println("a1 = ", a1)
	fmt.Println("a2 = ", a2)

	names := [3]string{"Mark", "Dhiraj", "Biden"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	num := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for _, value := range num {
		sum = sum + value
	}

	fmt.Printf("Sum of all the elements in array %v = %f \n", num, sum)

	//multidirectional
	multi := [2][2]int{
		{3, 5},
		{7, 9},
	}
	fmt.Println(multi)

	multi2 := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(multi2)
}
