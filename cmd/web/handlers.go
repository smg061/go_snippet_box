package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

func (app *application) home (w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" { 
		app.notFound(w)
		return 
	}

	files := []string{ "./ui/html/home.page.tmpl", "./ui/html/base.layout.tmpl", "./ui/html/footer.partial.tmpl" }



	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serveError(w, err)
		return;
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serveError(w, err)
		return;
	}

	w.Write([]byte("Hello from SnippetBox"))
} 

func (app *application)showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "display a specific snippet with id of %d", id)
}

func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "0 snail"
	content := "O nail \n Climb mount Fuji\nBut slowly, slowly!"
	expires := "7"
	id, err := app.snippets.Insert(title, content,expires)
	if err != nil  {
		app.serveError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

}


func (app *application)downloadHandler(w http.ResponseWriter, r *http.Request) { 
	http.ServeFile(w, r, filepath.Clean("./ui/static/file.zip"))
}

