package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

const (
	DataYaml = "data.yaml"
)

type Data struct {
	Title  string
	Groups []GroupData `yaml:"groups,flow"`
}

type GroupData struct {
	Name     string
	Sort     int
	Cols     int           `yaml:"cols"`
	Services []ServiceData `yaml:"services,flow"`
}

type ServiceData struct {
	Name        string
	Url         string
	Description string
	Components  []ComponentData
}

type ComponentData struct {
	Name        string
	Description string
	Url         string
}

func loadData(dataDir string) (*Data, error) {
	// Read file
	fileBytes, err := readFile(dataDir, DataYaml)
	if err != nil {
		return nil, err
	}
	// yaml bytes to struct
	data := new(Data)
	err = yaml.Unmarshal(fileBytes, data)
	if err != nil {
		return nil, err
	}
	// return
	return data, nil
}

func readFile(root string, filename string) ([]byte, error) {
	p, err := filepath.Abs(filepath.Join(root, filename))
	if err != nil {
		return nil, err
	}
	log.Printf("Reading data from path %s", p)
	d, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return d, nil
}
