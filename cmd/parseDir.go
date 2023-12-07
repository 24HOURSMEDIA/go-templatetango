/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"templatetango/tango"
	"templatetango/tango/fs_helpers"
)

// parseFileCmd represents the parseFile command
var parseDirCmd = &cobra.Command{
	Use:   "parse:dir",
	Short: "Parses all twig files in a directory and output to another directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("parse:template requires exactly two arguments")
			return
		}
		fileMask := "*.twig"
		fileStrip := ".twig"
		sourceDir := args[0]
		targetDir := args[1]

		// assert sourceDir is a dir and exists
		fs_helpers.AssertDirExists(sourceDir)
		fs_helpers.EnsureDirExists(targetDir)

		sourceFiles, err := fs_helpers.FindFilesInDir(sourceDir, fileMask)
		if err != nil {
			log.Fatalf("Error reading directory %s", sourceDir)
		}
		log.Printf("Found %d template files in directory %s\n", len(sourceFiles), sourceDir)

		filesMap := make(map[string]string)
		parsedMap := make(map[string]string)
		templateParams := tango.CreateParams()

		// create an environment using the source directory as the working directory
		stick := tango.CreateStickWithWorkDir(sourceDir)
		// parse all files
		for _, sourceFile := range sourceFiles {
			targetFile := strings.TrimSuffix(sourceFile, fileStrip)
			filesMap[sourceFile] = targetFile

			parsed, err := tango.ParseWithStickEnv(sourceFile, templateParams, stick)
			if err != nil {
				log.Fatalf("Error parsing file %s - %s", sourceFile, err)
			}
			parsedMap[sourceFile] = parsed
		}

		// write all parsed results
		for sourceFile, targetFile := range filesMap {
			fullSourcePath := filepath.Join(sourceDir, sourceFile)
			fullTargetPath := filepath.Join(targetDir, targetFile)
			parsed := parsedMap[sourceFile]
			fmt.Printf("Writing file %s to %s (%d bytes)\n", fullSourcePath, fullTargetPath, len(parsed))
			if err := os.WriteFile(fullTargetPath, []byte(parsed), 0644); err != nil {
				log.Fatalf("Error writing file %s", fullTargetPath)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(parseDirCmd)
}
