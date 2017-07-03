package main

import (
	"fmt"
	"testgo/databasehandling"
	"testgo/httphandling"
)

func main() {
	fmt.Println("Hello, Main!")
	maintest()
	databasehandling.PrintDBtestMessage()

	httphandling.CreateHTTPListener()
}

func maintest() {
	fmt.Println("Hello, maintest!")
}
