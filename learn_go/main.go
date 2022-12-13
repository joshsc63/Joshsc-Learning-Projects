// FUNC EXAMPLE for reference
// starting lowercase makes it private. Can only be called in this package
// adds two ints & returns sum
// func addValues(x, y int) int {
// 	//var sum int
// 	//sum = x + y
// 	//return sum
// 	return x + y
// }
// 
// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 10.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "cannot divide by zero")
// 		return
// 	}
// 
// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f if %f", 100.0, 10.0, f))
// }

//func divideValues(x, y float32) (float32, error) {
//	// dont divide by 0
//	if y <= 0 {
//		err := errors.New("cannot divide by zero")
//		return 0, err
//	}
//
//	result := x / y
//	return result, nil
//}


func main() {

	// HELLO WORLD
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "hello world!")
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	//	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	//})

	http.HandleFunc("/", Home)
	http.HandleFunc("/divide", About)

	// start web server
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
