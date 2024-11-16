package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Error while Parsing Template: %v", err)
		http.Error(w, "Error While Parsing Template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, "a string")
	if err != nil {
		log.Printf("Error while Executing Template: %v", err)
		http.Error(w, "Error While Executing Template", http.StatusInternalServerError)
	}
}

func contacthandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contenct-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><a href=\"https://www.google.com\">Google</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
  <ul>
	<li>
	  <b>Is there a free version?</b>
	  Yes! We offer a free trial for 30 days on any paid plans.
	</li>
	<li>
	  <b>What are your support hours?</b>
	  We have support staff answering emails 24/7, though response
	  times may be a bit slower on weekends.
	</li>
	<li>
	  <b>How do I contact support?</b>
	  Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	</li>
  </ul>
  `)
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
