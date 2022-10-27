package web

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *App) render(w http.ResponseWriter, status int, page string, data interface{}) {
	ts, ok := app.TemplateCache[page]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", page))
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *App) newTemplateData() *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}

func (app *App) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *App) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *App) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
