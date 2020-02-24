package handlers

import (
	"github.com/smoqadam/blogo/helpers"
	"github.com/smoqadam/blogo/lib"
	"github.com/smoqadam/blogo/models"
	"net/http"
	"strings"
	"time"
)

// View return a post by its slug
func View(app *lib.App, w http.ResponseWriter, r *http.Request) {
	vars := app.Vars(r)
	slug := vars["slug"]
	var p models.Post
	app.Model.DB.Preload("Tags").Where(&models.Post{Slug: slug}).First(&p)
	app.Response.Json(p)
}

// Posts return all posts
func Posts(app *lib.App, w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	app.Model.DB.Preload("Tags").Find(&posts)
	app.Response.Json(posts)
}

// Create create a new post
func Create(app *lib.App, w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	r.ParseForm()
	post.Content = r.FormValue("content")
	post.Title = r.FormValue("title")
	post.Slug = helpers.ToSlug(r.FormValue("title"))
	post.Tags = parseTags(app, r.FormValue("tags"))
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	app.Model.DB.Create(&post)
	app.Response.Json(post)
}

func parseTags(app *lib.App, s string) []models.Tag {
	var tags []models.Tag
	for _, t := range strings.Split(s, ",") {
		tag := models.Tag{
			Tag: t,
		}
		app.Model.DB.FirstOrCreate(&tag, tag)
		tags = append(tags, tag)
	}
	return tags
}
