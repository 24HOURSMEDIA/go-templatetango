package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// addTemplatesDirOption adds the --templates-AbsDir option to the command
func addTemplatesDirOption(cmd *cobra.Command) {
	cmd.Flags().StringP("templates-AbsDir", "d", "", "Root directory for templates. "+
		"Leave empty to use the directory of the file or directory to parse.",
	)
}

// resolveAbsFilePath returns the absolute path of the file path.
// If the file path is not absolute, it will be resolved relative to the workDirIfNotAbs directory.
func resolveAbsFilePath(filePath string, workDirIfNotAbs string) (string, error) {
	if filepath.IsAbs(filePath) {
		return filePath, nil
	}
	oldWorkDir, err := os.Getwd()
	if err != nil {
		return filePath, errors.New("Could not get current working directory")
	}
	if err := os.Chdir(workDirIfNotAbs); err != nil {
		return filePath, errors.New("Could not change working directory")
	}
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return filePath, errors.New("Could not get absolute path")
	}
	if err := os.Chdir(oldWorkDir); err != nil {
		return filePath, errors.New("Could not change working directory")
	}
	return absPath, nil
}

type fileResolveResult struct {
	AbsDir   string
	Relative string
}

func resolveTemplateFileDirAndPath(cmd *cobra.Command, filePath string) (fileResolveResult, error) {
	// Retrieve the current working directory and restore it later
	initialWorkDir, err := os.Getwd()
	if err != nil {
		return fileResolveResult{}, err
	}
	defer func() {
		if err := os.Chdir(initialWorkDir); err != nil {
			panic(err)
		}
	}()

	// Is there a custom directory for the templates?
	// If so, resolve filePath relative to it
	customTemplatesDir, err := cmd.Flags().GetString("templates-AbsDir")
	if err != nil {
		return fileResolveResult{}, err
	}
	if customTemplatesDir != "" {
		absCustomTemplatesDir, err := resolveAbsFilePath(customTemplatesDir, initialWorkDir)
		if err != nil {
			return fileResolveResult{}, err
		}
		if err := os.Chdir(customTemplatesDir); err != nil {
			return fileResolveResult{}, err
		}

		if filepath.IsAbs(filePath) {
			relativeFilePath, err := filepath.Rel(absCustomTemplatesDir, filePath)
			if err != nil {
				return fileResolveResult{}, err
			}
			return fileResolveResult{AbsDir: absCustomTemplatesDir, Relative: relativeFilePath}, nil
		}

		relativeFilePath, err := filepath.Rel(".", filePath)
		if err != nil {
			return fileResolveResult{}, err
		}
		return fileResolveResult{AbsDir: absCustomTemplatesDir, Relative: relativeFilePath}, nil
	}

	// No custom directory for the templates
	// Get the directory of the absolute filepath as the templates directory
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return fileResolveResult{}, err
	}
	absDir := filepath.Dir(absFilePath)
	if err != nil {
		return fileResolveResult{}, err
	}
	relativeFilePath, err := filepath.Rel(absDir, absFilePath)
	if err != nil {
		return fileResolveResult{}, err
	}
	return fileResolveResult{AbsDir: absDir, Relative: relativeFilePath}, nil
}
