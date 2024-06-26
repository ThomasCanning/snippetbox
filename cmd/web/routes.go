package main

import (
	"github.com/justinas/alice"
	"net/http"
)

/*
Used for keeping track of the pages and resources of the website
*/

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) //strip the "/static" prefix from the request URL before the request reaches the fileServer
	mux.HandleFunc("/", app.home)                                   //subtree path as it ends with a trailing slash, so it will match any path that has a prefix of "/"
	mux.HandleFunc("/snippet/view", app.snippetView)                //fixed path, so it will only match the exact path "/snippet/view"
	mux.HandleFunc("/snippet/create", app.snippetCreate)            //longer patterns take precedence over shorter ones

	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Return the 'standard' middleware chain followed by the servemux.
	return standard.Then(mux)
}
