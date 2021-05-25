package main

import "github.com/selimq/go-mux-postgre/app"

func main() {
	a := app.App{}
	a.Initialize(
		"postgres", "", "postgres")

	a.Run(":8010")
}
