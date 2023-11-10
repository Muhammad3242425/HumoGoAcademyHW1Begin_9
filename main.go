package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	
	fmt.Print("a = ")
	
	fmt.Scan(&a)
	fmt.Print("b =")
	fmt.Scan(&b)


	fmt.Println("Среднее геометрическое = ",math.Sqrt(a*b)) 
	

}
