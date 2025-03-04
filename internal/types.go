package internal

import (
	"os"
)

// cachemap structure
// methods defined in history.go
type History struct {
	// Holds the hashed file path and the file path
	seen   map[string]string `mapstructure:"seen"`
	count  int               `mapstructure:"count"`
	method string            `mapstructure:"method"`
}

// simple file and mode storage
type FileEntry struct {
	Path string      `mapstructure:"path"`
	Mode os.FileMode `mapstructure:"mode"`
}
