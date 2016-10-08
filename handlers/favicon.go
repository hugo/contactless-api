package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Favicon(rw http.ResponseWriter, r *http.Request) {
	file, err := os.Open("public/favicon.ico")
	if err != nil {
		log.Printf("Failed to open favicon file: %s\n", err.Error())
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	io.Copy(rw, file)
	return
}
