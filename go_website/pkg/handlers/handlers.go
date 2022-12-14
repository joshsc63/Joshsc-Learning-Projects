package handlers

import (
	"net/http"

	"github.com/joshsc63/go-website/pkg/config"
	"github.com/joshsc63/go-website/pkg/models"
	"github.com/joshsc63/go-website/pkg/render"
)

// the repository used by the handlers
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// Creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Handler Functions
// uppercase becomes public/visible outside of package
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// permorm some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"

	// send data to template

	// render
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
