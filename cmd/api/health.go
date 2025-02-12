package main

import "net/http"

func (app *application) healthCheckHundler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}