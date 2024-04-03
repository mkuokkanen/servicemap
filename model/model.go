package model

import (
	"context"
	"github.com/mkuokkanen/servicemap/configuration"
	"github.com/mkuokkanen/servicemap/pkl/gen"
	"html/template"
	"log"
	"sync"
)

type Model struct {
	config    *configuration.Config
	rwm       sync.RWMutex
	templates *template.Template
	data      *gen.Data
}

func NewModel(config *configuration.Config) (*Model, error) {
	log.Printf("Creating model")
	m := &Model{
		config: config,
	}
	err := m.loadResources()
	return m, err
}

func (m *Model) Load() error {
	// Lock readers out when doing changes
	m.rwm.Lock()
	defer m.rwm.Unlock()
	return m.loadResources()
}

func (m *Model) TemplateAndData() (*template.Template, *gen.Data) {
	// Verify that no writer has lock
	m.rwm.RLock()
	defer m.rwm.RUnlock()
	return m.templates, m.data
}

func (m *Model) loadResources() error {
	// load templates
	t, err := loadTemplates(m.config.TemplateDir)
	if err != nil {
		return err
	}
	// load data
	log.Printf("Reading data from path %s", m.config.DataFile)
	d, err := gen.LoadFromPath(context.Background(), m.config.DataFile)
	if err != nil {
		return err
	}
	// sort data
	sortData(d)
	// replace variables with new templates and data
	m.templates = t
	m.data = d
	return nil
}
