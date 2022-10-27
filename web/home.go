package web

import (
	"net/http"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	data := app.newTemplateData()
	app.render(w, http.StatusOK, "home.tmpl", data)
}
