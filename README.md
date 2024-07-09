
# Config

`config` is a Go package for reading and parsing JSON configuration files from multiple paths. It provides convenient methods to retrieve configuration values as strings, integers, and booleans, as well as a method to check the existence and validity of keys.

## Installation

To install the `config` package, use the following command:

```sh
go get github.com/oxisoft/config
```

## Usage

Here's an example of how to use the `config` package in your Go project.

### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/oxisoft/config"
)

func main() {
	configPaths := []string{"./config", "/etc/myapp", "$HOME/.myapp","."}
	cfg, err := config.NewConfig("config.json", configPaths)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Get string value
	dbHost := cfg.GetString("db_host")
	fmt.Printf("Database Host: %s\n", dbHost)

	// Get string value with error handling
	dbHost, err = cfg.GetStringWithError("db_host")
	if err != nil {
		log.Fatalf("Error getting db_host: %v", err)
	}
	fmt.Printf("Database Host: %s\n", dbHost)

	// Get int value
	port := cfg.GetInt("port")
	fmt.Printf("Port: %d\n", port)

	// Get int value with error handling
	port, err = cfg.GetIntWithError("port")
	if err != nil {
		log.Fatalf("Error getting port: %v", err)
	}
	fmt.Printf("Port: %d\n", port)

	// Get bool value
	debugMode := cfg.GetBool("debug")
	fmt.Printf("Debug Mode: %v\n", debugMode)

	// Get bool value with error handling
	debugMode, err = cfg.GetBoolWithError("debug")
	if err != nil {
		log.Fatalf("Error getting debug mode: %v", err)
	}
	fmt.Printf("Debug Mode: %v\n", debugMode)

	// Check keys
	requiredKeys := []string{"db_host", "port", "debug"}
	err = cfg.CheckKeys(requiredKeys)
	if err != nil {
		log.Fatalf("Error checking keys: %v", err)
	}
	fmt.Println("All required keys are present and valid")
}
```

## Methods

### `NewConfig`

```go
func NewConfig(configName string, configPaths []string) (*Config, error)
```

Creates a new `Config` instance, loading the JSON configuration from possible paths. Returns an error if the configuration file is not found or cannot be parsed.

### `GetString`

```go
func (c *Config) GetString(key string) string
func (c *Config) GetStringWithError(key string) (string, error)
```

- `GetString` returns the string value associated with the given key.
- `GetStringWithError` returns the string value associated with the given key and an error.

### `GetInt`

```go
func (c *Config) GetInt(key string) int64
func (c *Config) GetIntWithError(key string) (int64, error)
```

- `GetInt` returns the `int64` value associated with the given key, even if it's a string representation of an integer.
- `GetIntWithError` returns the `int64` value associated with the given key and an error.

### `GetBool`

```go
func (c *Config) GetBool(key string) bool
func (c *Config) GetBoolWithError(key string) (bool, error)
```

- `GetBool` returns the boolean value associated with the given key.
- `GetBoolWithError` returns the boolean value associated with the given key and an error.

### `CheckKeys`

```go
func (c *Config) CheckKeys(keys []string) error
```

Checks that all given keys exist and are not empty. Returns an error indicating which key does not exist or is empty.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to contribute to this project by submitting issues or pull requests.
