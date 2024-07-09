package config

import (
	"os"
	"path/filepath"
	"testing"
)

func createTestConfigFile(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	filePath := filepath.Join(dir, "config.json")

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	return filePath
}

func TestNewConfig(t *testing.T) {
	content := `{"key1": "value1", "key2": 9999, "key3": true}`
	configFilePath := createTestConfigFile(t, content)

	config, err := NewConfig("config.json", []string{filepath.Dir(configFilePath)})
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	if len(config.data) != 3 {
		t.Errorf("Expected 3 keys in config, got %d", len(config.data))
	}
}

func TestGetString(t *testing.T) {
	content := `{"key1": "value1"}`
	configFilePath := createTestConfigFile(t, content)

	config, err := NewConfig("config.json", []string{filepath.Dir(configFilePath)})
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	val := config.GetString("key1")
	if val != "value1" {
		t.Errorf("Expected 'value1', got '%s'", val)
	}

	_, err = config.GetStringWithError("key2")
	if err == nil {
		t.Error("Expected error for non-existent key, got nil")
	}
}

func TestGetInt(t *testing.T) {
	content := `{"key2": 9999, "key3": "8888"}`
	configFilePath := createTestConfigFile(t, content)

	config, err := NewConfig("config.json", []string{filepath.Dir(configFilePath)})
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	val := config.GetInt("key2")
	if val != 9999 {
		t.Errorf("Expected 9999, got %d", val)
	}

	val, err = config.GetIntWithError("key3")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 8888 {
		t.Errorf("Expected 8888, got %d", val)
	}

	_, err = config.GetIntWithError("key1")
	if err == nil {
		t.Error("Expected error for non-existent key, got nil")
	}
}

func TestGetBool(t *testing.T) {
	content := `{"key3": true}`
	configFilePath := createTestConfigFile(t, content)

	config, err := NewConfig("config.json", []string{filepath.Dir(configFilePath)})
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	val := config.GetBool("key3")
	if val != true {
		t.Errorf("Expected true, got %v", val)
	}

	_, err = config.GetBoolWithError("key1")
	if err == nil {
		t.Error("Expected error for non-existent key, got nil")
	}
}

func TestCheckKeys(t *testing.T) {
	content := `{"key1": "value1", "key2": 9999, "key3": true}`
	configFilePath := createTestConfigFile(t, content)

	config, err := NewConfig("config.json", []string{filepath.Dir(configFilePath)})
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	err = config.CheckKeys([]string{"key1", "key2", "key3"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = config.CheckKeys([]string{"key1", "key4"})
	if err == nil {
		t.Error("Expected error for non-existent key, got nil")
	}

	content = `{"key1": ""}`
	configFilePath = createTestConfigFile(t, content)

	config, err = NewConfig("config.json", []string{filepath.Dir(configFilePath)})
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	err = config.CheckKeys([]string{"key1"})
	if err == nil {
		t.Error("Expected error for empty string key, got nil")
	}
}
