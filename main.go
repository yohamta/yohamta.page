package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"website/web"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		infoLog.Printf("Defaulting to port %s", port)
	}

	templateCache, err := web.NewTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &web.App{
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		TemplateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     net.JoinHostPort("", port),
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("Listening on port %s", port)
	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}
