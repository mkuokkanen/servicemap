package model

import (
	"github.com/mkuokkanen/servicemap/configuration"
	"html/template"
	"log"
	"sync"
)

type Model struct {
	config    *configuration.Config
	rwm       sync.RWMutex
	templates *template.Template
	data      *Data
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

func (m *Model) TemplateAndData() (*template.Template, *Data) {
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
	d, err := loadData(m.config.DataDir)
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
