package config

import (
	"html/template"
)

// Dont import any other pkgs. Only STD library to avoid depedency loops

// Holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	//InfoLog       *log.Logger
}
