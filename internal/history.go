package internal

import (
	"fmt"
)

// History is the cache map
// This file contains the functions associated
// For the type definition see types.go

// Create a new empty cache map with key deriv of type m
func NewHistory(m string) *History {
	return &History{
		count:  0,
		seen:   make(map[string]string),
		method: m,
	}
}

// Get the count
func (h *History) Count() int {
	return h.count
}

// Clear the history
func (h *History) Clear() {
	h.seen = make(map[string]string)
	h.count = 0
}

// Get the key for a file
// If entry is in the history simply return a blank key
func (h *History) newKey(f string) (string, error) {
	key, err := hash(f, h.method)
	if err != nil || key == "" {
		return "", err
	}
	// Check if key in the cache
	_, ok := h.seen[key]
	if ok {
		return "", nil
	}

	return key, nil
}

// Add an entry to the cache map
func (h *History) addEntry(k string, f string) error {
	_, ok := h.seen[k]
	if ok {
		return fmt.Errorf("issue adding entry %v as it already exists", f)
	}

	h.seen[k] = f
	h.count += 1

	return nil
}
