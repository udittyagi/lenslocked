package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/udittyagi/lenslocked/controllers"
	"github.com/udittyagi/lenslocked/templates"
	"github.com/udittyagi/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(
		views.Must(
			views.ParseFs(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(
			views.ParseFs(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FaqHandler(views.Must(
		views.ParseFs(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.Get("/signup", controllers.FaqHandler(views.Must(
		views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "PAGE NOT FOUND", http.StatusNotFound)
	})

	fmt.Println("Server Running on Port 3000")
	http.ListenAndServe(":3000", r)
}
