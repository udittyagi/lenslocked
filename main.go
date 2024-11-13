package main

import (
	"fmt"
	"net/http"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, "<h1>Hello World Udit Tyagi</h1>")
// }

// func contacthandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Contenct-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, "<h1>Contact Page</h1><a href=\"https://www.google.com\">Google</a>")
// }

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

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hello World Udit Tyagi</h1>")
	case "/contact":
		fmt.Fprint(w, "<h1>Contact Page</h1><a href=\"https://www.google.com\">Google</a>")
	default:
		{
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "<h1>NOT FOUND</h1>")
		}

	}
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hello World Udit Tyagi</h1>")
	case "/contact":
		fmt.Fprint(w, "<h1>Contact Page</h1><a href=\"https://www.google.com\">Google</a>")
	case "/faq":
		faqHandler(w, r)
	default:
		{
			http.Error(w, "PAGE NOT FOUND", http.StatusNotFound)
		}
	}
}

//HandlerFunc type with base type as function ==> this is how http.HandlerFunc is implemented in go
// type HandlerFunc func(http.ResponseWriter, *http.Request)

// func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	f(w, r)
// }

func main() {
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contacthandler)
	// http.HandleFunc("/", pathHandler)
	fmt.Println("Server Running on Port 3000")
	// http.ListenAndServe(":3000", nil)

	//----------------------
	var handler Router
	http.ListenAndServe(":3000", handler)

	//---------------
	// http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
