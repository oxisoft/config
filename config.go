package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	data map[string]interface{}
}

// NewConfig creates a new Config instance, loading the JSON configuration from possible paths
func NewConfig(configName string, configPaths []string) (*Config, error) {
	var configData map[string]interface{}
	for _, path := range configPaths {
		// Expand environment variables in the path
		expandedPath := os.ExpandEnv(path)
		fullPath := filepath.Join(expandedPath, configName)
		if _, err := os.Stat(fullPath); err == nil {
			file, err := os.ReadFile(fullPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read config file: %w", err)
			}
			err = json.Unmarshal(file, &configData)
			if err != nil {
				return nil, fmt.Errorf("failed to parse config file: %w", err)
			}
			return &Config{data: configData}, nil
		}
	}
	return nil, fmt.Errorf("config file %s not found in paths: %v", configName, configPaths)
}

// GetString returns the string value associated with the given key
func (c *Config) GetString(key string) string {
	value, _ := c.GetStringWithError(key)
	return value
}

// GetStringWithError returns the string value associated with the given key and an error
func (c *Config) GetStringWithError(key string) (string, error) {
	if value, exists := c.data[key]; exists {
		if strValue, ok := value.(string); ok {
			return strValue, nil
		}
		return "", fmt.Errorf("value for key '%s' is not a string", key)
	}
	return "", fmt.Errorf("key '%s' not found", key)
}

// GetInt returns the int64 value associated with the given key, even if it's a string representation of an integer
func (c *Config) GetInt(key string) int64 {
	value, _ := c.GetIntWithError(key)
	return value
}

// GetIntWithError returns the int64 value associated with the given key and an error
func (c *Config) GetIntWithError(key string) (int64, error) {
	if value, exists := c.data[key]; exists {
		switch v := value.(type) {
		case float64:
			return int64(v), nil
		case int64:
			return v, nil
		case string:
			if intValue, err := strconv.ParseInt(v, 10, 64); err == nil {
				return intValue, nil
			} else {
				return 0, fmt.Errorf("value for key '%s' is not an int", key)
			}
		default:
			return 0, fmt.Errorf("value for key '%s' is not an int", key)
		}
	}
	return 0, fmt.Errorf("key '%s' not found", key)
}

// GetBool returns the boolean value associated with the given key
func (c *Config) GetBool(key string) bool {
	value, _ := c.GetBoolWithError(key)
	return value
}

// GetBoolWithError returns the boolean value associated with the given key and an error
func (c *Config) GetBoolWithError(key string) (bool, error) {
	if value, exists := c.data[key]; exists {
		if boolValue, ok := value.(bool); ok {
			return boolValue, nil
		}
		return false, fmt.Errorf("value for key '%s' is not a boolean", key)
	}
	return false, fmt.Errorf("key '%s' not found", key)
}

// CheckKeys checks that all given keys exist and are not empty
func (c *Config) CheckKeys(keys []string) error {
	for _, key := range keys {
		value, exists := c.data[key]
		if !exists || value == nil {
			return fmt.Errorf("key '%s' does not exist or is empty", key)
		}
		if strValue, ok := value.(string); ok && strValue == "" {
			return fmt.Errorf("key '%s' is an empty string", key)
		}
	}
	return nil
}
