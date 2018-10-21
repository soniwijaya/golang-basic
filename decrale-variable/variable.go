package main

import "fmt"

var zero int // this is zero value with type data integer

var u = 12 // global variable
var car string = `How this posibble, " String is in here "` // back tick is allowed wrap the text

func main() {

	// decrale variable and signing variable : =

	x := 42 // for number
	x = 99  // replace number

	fmt.Printf("%T ", x) // Print F for know the format variable
	fmt.Println(x)

	y := 20 + 10

	fmt.Println(y)

	z := "String harus double petik's "
	fmt.Println(z)

	var o = 13 // juga dapat digunakan didalam func
	fmt.Println(u, o)

	// variable car cannot be change in here because is set to be String not Number
	// example :
		// car = 12 // is not possible
	fmt.Println(car)
}
