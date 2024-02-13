package configuration

import (
	"fmt"
	"os"
)

const (
	envHttpPort    = "FP_HTTP_PORT"
	envStaticDir   = "FP_STATIC_DIR"
	envTemplateDir = "FP_TEMPLATE_DIR"
	envDataDir     = "FP_DATA_DIR"

	defaultHttpPort    = "8080"
	defaultStaticDir   = "./html/static"
	defaultTemplateDir = "./html/tmpl"
	defaultDataDir     = "./data"
)

type Config struct {
	HttpPort    string
	StaticDir   string
	TemplateDir string
	DataDir     string
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewConfig() *Config {
	fmt.Println("Creating config")
	return &Config{
		HttpPort:    getEnvOrDefault(envHttpPort, defaultHttpPort),
		StaticDir:   getEnvOrDefault(envStaticDir, defaultStaticDir),
		TemplateDir: getEnvOrDefault(envTemplateDir, defaultTemplateDir),
		DataDir:     getEnvOrDefault(envDataDir, defaultDataDir),
	}
}
