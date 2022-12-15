package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Dont import any other pkgs. Only STD library to avoid depedency loops

// Holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
