package main

import (
	"fmt"

	"github.com/tooluloope/go-rest-http/internal/db"
)

// Run Func would instantiate and startup the server

func Run() error {
	fmt.Println("Starting server...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Error connecting to db")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Error migrating db")
		return err
	}

	fmt.Println("Connected to db and Pinged successfully")
	return nil
}



func main() {
	fmt.Println("Hello world")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}