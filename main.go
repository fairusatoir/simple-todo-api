package main

import (
	"simple-to-do/utilities"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := InitializedApp()
	err := server.ListenAndServe()
	utilities.PanicOnError(err)
}
