package main

import (
	"log"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	_path := strings.Split(r.URL.Path, "/")[1:]
	_extt := strings.Split(_path[len(_path)-1], ".")
	_ext := _extt[len(_extt)-1]

	log.Printf("Request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)

	switch _ext {
	case "org":
		ServeMarkup(w, r)
	case "md":
		ServeMarkup(w, r)
	case "raw":
		ServeRaw(w, r)
	case "ico":
		ServeFavico(w, r)
	default:
		ServeIndex(w, r)
	}
}

func ServeOrg(w http.ResponseWriter, r *http.Request) {
	return
}

func ServeMdown(w http.ResponseWriter, r *http.Request) {
	return
}

func ServeRaw(w http.ResponseWriter, r *http.Request) {
	return
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	return
}

func ServeFavico(w http.ResponseWriter, r *http.Request) {
	return
}
