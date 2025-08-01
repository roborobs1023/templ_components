package main

import (
	"bufio"
	"fmt"

	"os"
	"path/filepath"
	"strings"

	"github.com/roborobs1023/templ_components/internal/cfg"
)

func CombineCSSInDirectory(rootPath string) (string, error) {
	var combinedCSS strings.Builder
	var cssFilesFound []string

	// Use filepath.WalkDir for efficient directory traversal
	err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// If an error occurred while walking, log it and decide whether to stop or continue.
			// Here, we'll just print a warning and continue, but you might want to return err
			// if you want to stop on the first directory traversal error.
			fmt.Printf("Warning: error accessing path %s: %v\n", path, err)
			return nil // Don't return error to continue walking
		}

		// Skip directories themselves, we only care about files
		if d.IsDir() {
			return nil
		}

		// Check if the file has a .css extension
		if strings.HasSuffix(strings.ToLower(d.Name()), ".css") {
			cssFilesFound = append(cssFilesFound, path)
		}
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error walking directory %s: %w", rootPath, err)
	}

	// Now that we have all the CSS file paths, use the previous logic to combine them
	if len(cssFilesFound) == 0 {
		fmt.Printf("No .css files found in %s and its subdirectories.\n", rootPath)
		return "", nil // Or return a specific error if no files found is an error for your use case
	}

	for _, filePath := range cssFilesFound {
		// Read the content of the CSS file
		content, readErr := os.ReadFile(filePath)
		if readErr != nil {
			// Log the error but continue to the next file, as one bad file shouldn't stop the whole process
			fmt.Printf("Warning: error reading CSS file %s: %v\n", filePath, readErr)
			continue
		}

		// Append content to the builder, followed by a newline for separation
		// (helps maintain readability in the combined output)
		combinedCSS.Write(content)
		combinedCSS.WriteString(fmt.Sprintf("\n/* --- End of %s --- */\n\n", filepath.Base(filePath))) // Add comment for clarity
	}

	return combinedCSS.String(), nil
}

func main() {
	cssSourceDir := "."

	outputDirInput := cfg.Config.String("file-dirs.css")

	if outputDirInput == "" {
		// --- Ask user for the output directory ---
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the output directory for combined.css (e.g., 'dist' or 'build/css'): ")
		outputDirInput, _ = reader.ReadString('\n')
	}
	outputDirInput = strings.TrimSpace(outputDirInput) // Remove newline and spaces

	if outputDirInput == "" {
		fmt.Fprintf(os.Stderr, "Error: Output directory cannot be empty. Exiting.\n")
		os.Exit(1)
	}
	// --- End user input ---

	outputFile := filepath.Join(outputDirInput, "combined.css")

	fmt.Printf("Generating combined CSS to %s...\n", outputFile)

	// Ensure the output directory exists
	// Use the user-provided outputDirInput here
	if _, err := os.Stat(outputDirInput); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDirInput, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output directory %s: %v\n", outputDirInput, err)
			os.Exit(1)
		}
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Error checking output directory %s: %v\n", outputDirInput, err)
		os.Exit(1)
	}

	combinedContent, err := CombineCSSInDirectory(cssSourceDir) // Call your function
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error combining CSS files: %v\n", err)
		os.Exit(1)
	}

	if _, err = os.Stat(outputFile); !os.IsNotExist(err) {
		os.Remove(outputFile)
	}

	err = os.WriteFile(outputFile, []byte(combinedContent), 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing combined CSS to %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated combined CSS to %s.\n", outputFile)
}
