package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/udittyagi/lenslocked/views"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	tpl, err := views.Parse(filePath)
	if err != nil {
		log.Printf("Error while Parsing Template: %v", err)
		http.Error(w, "Error While Parsing Template", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contacthandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	galleryId := chi.URLParam(r, "id")
	fmt.Fprintf(w, "<h1>Gallery Id %v</h1>", galleryId)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contacthandler)
	r.Get("/faq", faqHandler)
	r.Get("/gallery/{id}", galleryHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "PAGE NOT FOUND", http.StatusNotFound)
	})

	fmt.Println("Server Running on Port 3000")
	http.ListenAndServe(":3000", r)
}
