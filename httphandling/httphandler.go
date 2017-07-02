package httphandling

import "fmt"

//PrintHTTPtestMessage prints a nice message
func PrintHTTPtestMessage() {
	fmt.Println("Hello, http!")

}

func httpStupidFunction(someNumber int) int {
	return someNumber + 3
}
