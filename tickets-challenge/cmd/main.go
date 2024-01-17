package main

import (
	"fmt"
	"os"
	"tickets-challenge/internal/application"
)

func main() {
	// env
	// ...

	// application
	// - config
	cfg := &application.ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := application.NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println("run setup", err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println("run app", err)
		return
	}
}
