package main

import (
	"errors"
	"fmt"
	"github.com/ThomasCanning/snippetbox/internal/models"
	"net/http"
	"strconv"
)

/*
Each page has a handler that takes in a http writer and a request path. Checks can be done on the request path to make sure requested path is correct, e.g. for subtree paths restricting their catch-all behaviour.
The rest of the function then dictates what happens on that page, with the http writer being used to write to the page.
*/

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { //has effect of overwriting subtree behavior to make it exact path match
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "home.tmpl", &templateData{
		Snippets: snippets,
	})
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, http.StatusOK, "view.tmpl", &templateData{
		Snippet: snippet,
	})
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	//temp dummy data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
