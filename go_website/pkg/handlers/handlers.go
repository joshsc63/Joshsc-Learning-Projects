package handlers

import (
	"net/http"

	"github.com/joshsc63/go-website/pkg/render"
)

// Handler Functions
// uppercase becomes public/visible outside of package
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
