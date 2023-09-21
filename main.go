package main

import "fmt"

func main(){
	fmt.Println("Hello 🚗💶")

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("%d \t %b \t %#x\n", a, a, a)
	fmt.Printf("%d \t %b \t %#x\n", b, b, b)
	fmt.Printf("%d \t %b \t %#x\n", c, c, c)
	fmt.Printf("%d \t %b \t %#x\n", d, d, d)
	fmt.Printf("%d \t %b \t %#x\n", e, e, e)
	fmt.Printf("%d \t %b \t %#x\n", f, f, f)
}

