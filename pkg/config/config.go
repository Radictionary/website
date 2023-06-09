package config

import (
	"fmt"
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

func Handle(err error, description string) {
	if err != nil {
		fmt.Printf("%v:%v\n", description, err)
	}
}
