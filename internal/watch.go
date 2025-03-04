package internal

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/astr0n8t/inotify-tasker/config"
)

// Starts the main watcher loop
//
// No arguments and no return since it loops infinitely
func Watch() {

	// Grab all of the config details we need
	config := config.Config()

	directory := config.GetString("directory")
	log.Printf("Monitoring directory: %v", directory)
	recursive := config.GetBool("recursive")
	log.Printf("Monitoring recursively: %v", recursive)
	interval := config.GetInt("interval")
	log.Printf("Using interval: %v", interval)
	verbose := config.GetBool("verbose")
	log.Printf("Verbose mode: %v", verbose)
	method := config.GetString("method")
	log.Printf("Using hash method: %v", method)
	clearOnEmpty := config.GetBool("clear_on_empty")
	log.Printf("Clearing cache when no files in dir: %v", clearOnEmpty)

	// Create a new cache map using our key deriv method
	history := NewHistory(method)

	// Start our main event loop
	for {
		// Set our initial vars
		count := history.Count()
		fileList := make([]FileEntry, 0)

		// Get the list of files to look at
		makeFileList(directory, recursive, &fileList)

		// If the file list is empty, let's clear the cache if var set
		if len(fileList) == 0 {
			if clearOnEmpty {
				history.Clear()
				if verbose {
					log.Printf("Cleared the history as no dir entries")
				}
			}
		} else {
			// Otherwise process our file list
			processFileList(history, fileList, verbose)
		}

		if verbose {
			log.Printf("Processed %v new entries", history.Count()-count)
		}

		// Sleep the defined amount of time between checks
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

// Returns a slice of FileEntry objects to evaluate which contain full file paths with their masks
//
// Arguments:
//
// directory string: the directory that is being watched
// recursive bool: whether to descend into sub-dirs
// fileList: the slice to store the FileEntry objects in
func makeFileList(directory string, recursive bool, fileList *[]FileEntry) {
	// Check if we should descend into sub-dirs
	if recursive {
		// Use filepath since it will do what we want
		err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
			// We don't want to add directories
			if !f.IsDir() {
				// Append the info we need to our slice
				*fileList = append(*fileList, FileEntry{
					Path: path,
					Mode: f.Mode(),
				})
			}
			return err
		})
		if err != nil {
			log.Fatalf("issue reading watch directory: %v", err)
		}
	} else {
		// Just use a normal os readdir
		entries, err := os.ReadDir(directory)
		if err != nil {
			log.Fatalf("issue reading watch directory: %v", err)
		}
		// Iterate over the results of the readdir
		for _, f := range entries {
			// Again no dirs
			if !f.IsDir() {
				// Add just the info we need
				*fileList = append(*fileList, FileEntry{
					Path: directory + "/" + f.Name(),
					Mode: f.Type(),
				})
			}
		}
	}
}

// Returns null, but updates history in place once entries are processed
//
// Arguments:
//
// history *History: the cache map to check against
// fileList []FileEntry: the slice with FileEntry objects
// verbose: whether or not to print status
func processFileList(history *History, fileList []FileEntry, verbose bool) {
	// Iterate over our FileEntry objects
	for _, e := range fileList {
		if verbose {
			log.Printf("Processing entry: %v", e.Path)
		}
		// Grab the key
		key, err := history.newKey(e.Path)
		// Check to make sure our key actually got created
		if err != nil  {
			log.Printf("ERROR: could not hash entry: %v : %v", e.Path, err)
		} else if key == "" {
			log.Printf("Entry already contained in history or has duplicate hash: %v", e.Path)
		// Entry needs to be added
		} else {
			// Try to update file
			err := UpdateFile(e.Path, e.Mode)
			if err != nil {
				log.Printf("ERROR: issue processing entry: %v : %v", e.Path, err)
			} else {
				// We can add the entry
				history.addEntry(key, e.Path)
				if verbose {
					log.Printf("Processed entry: %v with key: %v", e.Path, key)
				}
			}
		}
	}
}
