package main

import "github.com/selimq/go-mux-postgre/app"

func main() {
	a := app.App{}
	a.Initialize(
		"postgres", "SamplePass", "postgres")

	a.Run(":8010")
}

/*
user: 'postgres',
host: '192.168.79.167',
database: 'postgres',
password: 'SamplePass',
port: 5432, */
