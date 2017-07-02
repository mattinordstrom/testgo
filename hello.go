package main

import (
	"fmt"
	"gotest/databasehandling"
	"gotest/httphandling"
)

func main() {
	fmt.Println("Hello, Main!")
	maintest()
	httphandling.PrintHTTPtestMessage()
	databasehandling.PrintDBtestMessage()
}

func maintest() {
	fmt.Println("Hello, maintest!")
}
