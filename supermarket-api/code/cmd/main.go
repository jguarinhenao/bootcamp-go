package main

import (
	"app/scaffolding/internal/application"
	"fmt"
)

func main() {

	// app
	// - config
	app := application.NewDefaultHTTP(":8080")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
