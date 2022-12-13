package main

import "log"

func main() {
	var myString string
	myString = "green"

	log.Println("mystring is set to", myString)
	changeUsingPointer(&myString) //pass reference of variable of its memory location.... add &
	log.Println("after func call, myString is set to", myString)
}

// *string is the REFERENCE to place in memory
func changeUsingPointer(s *string) {
	log.Println("s is set to", s)

	newValue := "Red"
	*s = newValue
}
