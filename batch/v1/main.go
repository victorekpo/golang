package main

import (
	"batch-v1/internal/app"
	"fmt"
)

func init() {
	// Initialization code here, can go for any package
	fmt.Println("Welcome to my App - init")
}

func main() {
	defer func() {
		//fmt.Println("Post script code here")
	}()
	//app.TestWorkers()
	app.App()
	// server.Server()
}

/*
In Go, the init function is a special function that is automatically executed when a
package is initialized. It is used to set up any necessary state or perform any initialization
tasks before the main execution begins. Each package can have multiple init functions, and
they are executed in the order they appear within the package.

There is no direct equivalent of a "post script" function that runs after the main function.
However, you can achieve similar behavior by placing code at the end of the main function or by
using deferred functions.
*/
