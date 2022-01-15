package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	fmt.Scanln(&number)
	var x float64 = 2
	var prima bool = true
	for x < math.Sqrt(number) {
		if math.Mod(number, x) == 0 {
			prima = false
			break
		} else {
			x++
		}
	}
	if prima {
		fmt.Printf("%f is prime", number)
	} else {
		fmt.Printf("%f isn't prime", number)
	}

}
