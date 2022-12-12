package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact contactInfo
	contactInfo
}

func main() {
	//Ordered arguments, sucks
	//alex := person{"Alex", "Anderson"}

	//alex := person{firstName: "Alex", lastName: "Anderson"}

	//var alex person //vals are 0
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	//fmt.Println(alex) // { }
	//fmt.Printf("%+v", alex) //{firstName: lastName:}

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		//contact: contactInfo{
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 12345,
		},
	}

    // pointer needed to update record in memory
    //jimPointer := &jim // &variable - give memory address that var is pointing at
	//jimPointer.updateName("Jimmy")

	//shortcut - go will convert `person` to `*person` when func receiver is a *pointer
	jim.updateName("jimmy")
	jim.print()
}

// USE POINTERS FOR INT, FLOAT, STRING, BOOL, STRUCTS
// DONT USE POINTERS FOR SLICES, MAPS, CHANNELS, POINTERS, FUNCTIONS

// ADDR            VAL
// 0001      person{firstName: "jim"}
// Turn addres into value with *address
// Turn value into address with &value
func (pointerToPerson *person) updateName(newFirstName string) { //*person - type of pointer to person
	//p.firstName = newFirstName //wont update due to pointing to p in memory, not jim

	// *pointerToPerson - operator to manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName //*pointer - give me memory address pointing to - turns into the struct val
}

func (p person) print() {
	fmt.Printf(("%+v"), p)
}

