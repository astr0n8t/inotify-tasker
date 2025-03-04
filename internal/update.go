package internal

import (
	"fmt"
	"os"
)

// This function actually does the touching of the file
// Basically just opens the file, seeks to the end, writes nothing, and closes it
//
// Arguments:
// f string: the full file path to the file to update
// m os.FileMode: the OS mask of the file so that perms are preserved
func UpdateFile(f string, m os.FileMode) error {
	// Try to open the file with its normal perms
	file, err := os.OpenFile(f, os.O_RDWR, m)
	if err != nil {
		return fmt.Errorf("issue opening file: %v", err)
	}

	// Try to seek to the end of the file
	_, err = file.Seek(0, os.SEEK_END)
	if err != nil {
		return fmt.Errorf("Error seeking to the end of the file: %v", err)
	}

	// Try to write nothing to the file
	nothing := []byte{}
	_, err = file.Write(nothing)
	if err != nil {
		return fmt.Errorf("Error writing to the file: %v", err)
	}

	// Close the file
	file.Close()

	return nil
}
