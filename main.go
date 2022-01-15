package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	if number%2 == 0 && number != 0 {
		fmt.Printf("%d is even", number)
	} else if number == 0 {
		fmt.Println("number is 0")
	} else {
		fmt.Printf("%d is odd", number)
	}

}
