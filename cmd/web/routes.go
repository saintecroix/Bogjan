package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/", app.home)
	m.HandleFunc("/send-to-mail", app.send_to_mail)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	m.Handle("/static/", http.StripPrefix("/static", fileServer))

	return m
}
