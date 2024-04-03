package configuration

import (
	"log"
	"os"
)

const (
	envHttpPort    = "FP_HTTP_PORT"
	envStaticDir   = "FP_STATIC_DIR"
	envTemplateDir = "FP_TEMPLATE_DIR"
	envDataFile    = "FP_DATA_FILE"

	defaultHttpPort    = "8080"
	defaultStaticDir   = "./html/static"
	defaultTemplateDir = "./html/tmpl"
	defaultDataFile    = "./pkl/data/data.pkl"
)

type Config struct {
	HttpPort    string
	StaticDir   string
	TemplateDir string
	DataFile    string
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewConfig() *Config {
	log.Printf("Creating config")
	return &Config{
		HttpPort:    getEnvOrDefault(envHttpPort, defaultHttpPort),
		StaticDir:   getEnvOrDefault(envStaticDir, defaultStaticDir),
		TemplateDir: getEnvOrDefault(envTemplateDir, defaultTemplateDir),
		DataFile:    getEnvOrDefault(envDataFile, defaultDataFile),
	}
}
