package main

import (
	"github.com/smoqadam/blogo/handlers"
	"github.com/smoqadam/blogo/lib"
	"github.com/smoqadam/blogo/models"
)

func main() {
	app := lib.NewApp()

	// Register new routes
	app.Get("/", handlers.Index)
	app.Get("/posts", handlers.Posts)
	app.Get("/post/{slug}", handlers.View)
	app.Post("/post", handlers.Create)

	// Add model here for migration
	app.Model.DB.AutoMigrate(models.Post{}, models.Tag{})

	// Run the web server
	app.Run(":3000")
}
