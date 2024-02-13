package configuration

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	var config = NewConfig()
	if config.DataDir == "" {
		t.Fatalf(`DataDir was empty instead of default`)
	}
	if config.TemplateDir == "" {
		t.Fatalf(`TemplateDir was empty instead of default`)
	}
	if config.StaticDir == "" {
		t.Fatalf(`StaticDir was empty instead of default`)
	}
	if config.HttpPort == "" {
		t.Fatalf(`HttpPort was empty instead of default`)
	}
}

func TestNewConfigWithEnvs(t *testing.T) {
	_ = os.Setenv("FP_HTTP_PORT", "TEST_PORT")
	_ = os.Setenv("FP_STATIC_DIR", "TEST_STATIC")
	_ = os.Setenv("FP_TEMPLATE_DIR", "TEST_TEMPLATE")
	_ = os.Setenv("FP_DATA_DIR", "TEST_DATA")

	config := NewConfig()
	if config.HttpPort != "TEST_PORT" {
		t.Fatalf("FP_HTTP_PORT env variable was not passed to config")
	}
	if config.TemplateDir != "TEST_TEMPLATE" {
		t.Fatalf(`FP_TEMPLATE_DIR env variable was not passed to config`)
	}
	if config.StaticDir != "TEST_STATIC" {
		t.Fatalf(`FP_STATIC_DIR env variable was not passed to config`)
	}
	if config.DataDir != "TEST_DATA" {
		t.Fatalf(`FP_DATA_DIR env variable was not passed to config`)
	}
}
