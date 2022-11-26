package main

import (
	"fmt"
	"github.com/Qadriddin-dev/goPRJ/controllers"
	"github.com/Qadriddin-dev/goPRJ/templates"
	"github.com/Qadriddin-dev/goPRJ/views"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))
	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :8000...")
	http.ListenAndServe(":8000", r)
}
