package web

import (
	"net/http"
	"website/ui"
)

func (app *App) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)

	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("/static/", fileServer)

	return mux
}
