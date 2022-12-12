package main

import "fmt"

//if type with function getGreeting, then that type is also a bot interface
// You can also use printGreeting bc it accepts type bot
// becomes interface type. Are NOT generic types
type bot interface {
	getGreeting() string
	//getGreeting(string, int) (string, error) // (arguments) (returns)
	
}

type englishBot struct {} //becomes a concrete type along side map, struct, int, string
type spanishBot struct {}

func main () {
	eb := englishBot{}
	sb := spanishBot{}
	
	printGreeting(eb)
	printGreeting(sb)
}

// Double funcs doing same thing... integrate interface to re-use logic across different types
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func (eb englishBot) getGreeting() string { // eb not in use, so can omit from receiver
func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola amigo!"
}

