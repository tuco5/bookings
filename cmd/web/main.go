package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tuco5/bookings/pkg/config"
	"github.com/tuco5/bookings/pkg/handlers"
	"github.com/tuco5/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const port = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting application on http://localhost%s\n", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
