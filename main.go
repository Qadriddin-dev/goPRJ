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

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 NOT FOUND:(", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :8000...")

	http.ListenAndServe(":8000", r)
}
