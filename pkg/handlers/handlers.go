package handlers

import (
	"net/http"

	"github.com/tuco5/bookings/pkg/config"
	"github.com/tuco5/bookings/pkg/models"
	"github.com/tuco5/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.Get(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: map[string]string{"test": "hello world"},
		Data:      map[string]interface{}{"remote_ip": remoteIP},
	})
}
