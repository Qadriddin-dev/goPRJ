package main

import (
	"fmt"
	"github.com/Qadriddin-dev/goPRJ/controllers"
	"github.com/Qadriddin-dev/goPRJ/models"
	"github.com/Qadriddin-dev/goPRJ/templates"
	"github.com/Qadriddin-dev/goPRJ/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	sessionService := models.SessionService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS,
		"signup.gohtml", "tailwind.gohtml"))

	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS,
		"signin.gohtml", "tailwind.gohtml"))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :8080...")

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMW := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false))
	http.ListenAndServe(":8080", csrfMW(r))
}
