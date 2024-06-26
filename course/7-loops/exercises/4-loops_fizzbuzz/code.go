package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++{
		if i % 3 && i % 5{
			fmt.Println("fizzbuzz")
		}else if i % 3{
			fmt.Println("fizz")
		}else if i % 5{
			fmt.Println("buzz")
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
