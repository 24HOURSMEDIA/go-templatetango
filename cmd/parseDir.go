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

		fmt.Println("Tango: parsing directory `" + sourceDir + "` to `" + targetDir + "`")

		templatesDir, err := resolveTemplateFileDirAndPath(cmd, sourceDir)
		dirToScan := filepath.Join(templatesDir.AbsDir, templatesDir.Relative)
		fmt.Println("- using templates dir: " + templatesDir.AbsDir)
		fmt.Println("- directory to scan: " + dirToScan)

		// assert sourceDir is a AbsDir and exists
		fs_helpers.AssertDirExists(dirToScan)
		fs_helpers.EnsureDirExists(targetDir)

		// Make a map of all sourceFiles to targetFiles, stripping the extension
		fmt.Println("- scanning directory " + dirToScan)
		filesMap := makeMapOfFiles(dirToScan, targetDir, fileMask, fileStrip)
		fmt.Printf("- found %d files to parse\n", len(filesMap))

		for sourceFile, targetFile := range filesMap {
			debugPrintf(cmd, "Source file: %s\n", sourceFile)
			debugPrintf(cmd, "Target file: %s\n", targetFile)
			debugPrintln(cmd)
		}

		templateParams := tango.CreateParams()

		type parseResult struct {
			sourceFile string
			targetFile string
			parsed     string
		}
		parsed := make([]parseResult, 0)
		for sourceFile, targetFile := range filesMap {
			if err != nil {
				log.Fatalf("Error resolving file %s - %s", sourceFile, err)
			}

			relSourceFile, err := filepath.Rel(templatesDir.AbsDir, sourceFile)
			stickEnv := tango.CreateStickWithWorkDir(templatesDir.AbsDir)
			parsedContent, err := tango.ParseWithStickEnv(relSourceFile, templateParams, stickEnv)
			if err != nil {
				log.Fatalf("Error parsing file %s - %s", sourceFile, err)
			}

			result := parseResult{
				sourceFile: sourceFile,
				targetFile: targetFile,
				parsed:     parsedContent,
			}
			parsed = append(parsed, result)
		}

		// Write all the parsed results if not in debug mode
		if isInDebugMode(cmd) {
			fmt.Println("Debug mode: not writing files")

			for _, result := range parsed {
				fmt.Println("=====================================================")
				fmt.Printf("parsing %s\n", result.sourceFile)
				fmt.Println(result.targetFile + ":")
				fmt.Println("```")
				fmt.Println(result.parsed)
				fmt.Println("```")
				fmt.Println()
			}

			return
		}

		for _, result := range parsed {
			fmt.Printf("- writing parsed file %s to %s (%d bytes)\n", result.sourceFile, result.targetFile,
				len(result.parsed))
			if err := os.WriteFile(result.targetFile, []byte(result.parsed), 0644); err != nil {
				log.Fatalf("Error writing file %s", result.targetFile)
			}
		}
	},
}

// makeMapOfFiles makes a map of absolute paths source files to absolute paths of target files
// The target file names are the source files with the fileStrip suffix removed
func makeMapOfFiles(sourceDir string, targetDir string, fileMask string, fileStrip string) map[string]string {
	filesMap := make(map[string]string)
	sourceFiles, err := fs_helpers.FindFilesInDir(sourceDir, fileMask)
	if err != nil {
		log.Fatalf("Error reading directory %s", sourceDir)
	}
	absTargetDir, err := filepath.Abs(targetDir)
	if err != nil {
		log.Fatalf("Error getting absolute path of %s", targetDir)
	}
	for _, sourceFile := range sourceFiles {
		absSourceFile, err := filepath.Abs(filepath.Join(sourceDir, sourceFile))
		if err != nil {
			log.Fatalf("Error getting absolute path of %s", sourceFile)
		}
		targetFile := strings.TrimSuffix(sourceFile, fileStrip)
		filesMap[absSourceFile] = filepath.Join(absTargetDir, targetFile)
	}
	return filesMap
}

// init registers the command with Cobra
func init() {
	addTemplatesDirOption(parseDirCmd)
	addDebugOption(parseDirCmd, "Does not write files but prints the parsed content to stdout")
	rootCmd.AddCommand(parseDirCmd)
}
