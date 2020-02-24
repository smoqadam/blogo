package handlers

import (
	"github.com/smoqadam/blogo/lib"
	"net/http"
)

func Index(app *lib.App, w http.ResponseWriter, r *http.Request) {
	app.Response.Html("Index")
}
