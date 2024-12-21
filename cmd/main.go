package main

import "github.com/develslawer/webcalc/internal/application"

func main() {
	app := application.New()
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
