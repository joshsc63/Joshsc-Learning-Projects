package main

import "fmt"

func main() {
	var whatToSay string
	var i int

	whatToSay = "Goodbye world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingSaid := saySomething()

	fmt.Println("function returned", whatWasSaid, theOtherThingSaid)

}

func saySomething() (string, string) {
	return "something", "else"
}
