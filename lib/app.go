// App is responsible for handling routes, response, models and run the web server
package lib

import (
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router   *mux.Router
	Response *Response
	Model    *Model
}

type httpHandler = func(app *App, w http.ResponseWriter, r *http.Request)

// Return new app
func NewApp() *App {
	a := App{}
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Response = &Response{}
	a.Model = NewModel()
	return &a
}

// HandleRequest handling http requests based on method.
// It will send an instance of app to the handler function.
func (a *App) HandleRequest(path string, method string, handler httpHandler) {
	a.Router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		a.Response.Writer = writer
		handler(a, writer, request)
	}).Methods(method)
}

// Get add GET http handler
func (a *App) Get(path string, f httpHandler) {
	a.HandleRequest(path, "GET", f)
}

// Post add POST http handler
func (a *App) Post(path string, f httpHandler) {
	a.HandleRequest(path, "POST", f)
}

// Vars is equivalent of mux.Vars(r *http.Request)
func (a *App) Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// Run run the web server
func (a *App) Run(addr string) {
	http.ListenAndServe(addr, a.Router)
}
