package model

import (
	"html/template"
	"log"
	"path/filepath"
)

const (
	MainTemplate = "main.gohtml"
)

func loadTemplates(templateDir string) (*template.Template, error) {
	p, err := filepath.Abs(filepath.Join(templateDir, MainTemplate))
	if err != nil {
		return nil, err
	}
	log.Printf("Reading template from path %s", p)
	var templates = template.Must(template.ParseFiles(p))
	return templates, nil
}
