package main

import (
	"context"
	"fmt"

	"github.com/tooluloope/go-rest-http/internal/comment"
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

	cmtService := comment.NewService(db)
	cmt, err := cmtService.PostComment(
		context.Background(),
		comment.Comment{
			Slug:   "test",
			Body:   "test",
			Author: "TOlulope",
			ID:     "42c21d77-9759-4658-bd33-e415a5bf9011",
		},
	)

	fmt.Println(cmt, err)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"42c21d77-9759-4658-bd33-e415a5bf9011",
	))
	return nil
}

func main() {
	fmt.Println("Hello world")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
