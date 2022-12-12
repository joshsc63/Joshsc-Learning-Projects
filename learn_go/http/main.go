package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main () {
	resp, err := http.Get("http://google.com") 
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
	//bs := []byte{}
	//bs := make([]byte, 99999) // make takes type (of slice) & argument of # of elements to be init'd w/ empty elements
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	
	//func Copy(dst Writer, src Reader)
	// os.Stdout -> Value of type File -> File has function 'Write' -> Implements 'Write' interface
	//io.Copy(os.Stdout, resp.Body)
}

//This integrates the type Writer interface bc logWriter has func called Write
//type Writer interface {
//    Write(p []byte) (n int, err error)
//}
func (logWriter) Write(bs []byte) (int, error) {
	//return 1, nil	
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))	
	return len(bs), nil
}