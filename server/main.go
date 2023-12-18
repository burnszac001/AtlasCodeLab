package main

import "fmt"

func main() {
	app := NewApp()
	if err := app.Run(":8080"); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
