package fs_helpers

import (
	"log"
	"os"
	"path/filepath"
)

// FindFilesInDir reads all files in a given directory that match the specified pattern,
// or all files if the pattern is empty, and returns a slice of filenames.
func FindFilesInDir(dirPath, pattern string) ([]string, error) {
	var filenames []string

	// Check if the pattern is empty
	if pattern == "" {
		// If the pattern is empty, read all files
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				filenames = append(filenames, entry.Name())
			}
		}
	} else {
		// If there is a pattern, match files against the pattern
		fullPathPattern := filepath.Join(dirPath, pattern)
		matches, err := filepath.Glob(fullPathPattern)
		if err != nil {
			return nil, err
		}

		for _, match := range matches {
			if info, err := os.Stat(match); err == nil && !info.IsDir() {
				filenames = append(filenames, filepath.Base(match))
			}
		}
	}

	return filenames, nil
}

// AssertDirExists checks if a directory exists, and logs a fatal error if it doesn't.
func AssertDirExists(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", dirPath)
	} else if err != nil {
		log.Fatalf("Error checking directory: %s", err)
	}
}

// EnsureDirExists checks if a directory exists, and creates it if it doesn't.
// Logs a fatal error if any operation fails.
func EnsureDirExists(dirPath string) {
	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Directory does not exist, attempt to create it
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			log.Fatalf("Failed to create directory: %s, error: %s", dirPath, err)
		}
		log.Printf("Directory created: %s", dirPath)
	} else if err != nil {
		// An error occurred during os.Stat (other than NotExist)
		log.Fatalf("Error checking directory: %s, error: %s", dirPath, err)
	}
}
