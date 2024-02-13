package model

import (
	"fmt"
	"html/template"
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
	fmt.Println("Reading template from path", p)
	var templates = template.Must(template.ParseFiles(p))
	return templates, nil
}
