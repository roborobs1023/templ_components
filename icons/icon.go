package icons

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/a-h/templ"
)

// iconContents caches the fully generated SVG strings from icons that have been used, keyed by the composite key of name and props to handle different styles.
var (
	iconContents = make(map[string]string)
	iconMutex    sync.RWMutex
)

const (
	BOX    = "bxicons"
	LUCIDE = "lucide"
	CUSTOM = "custom"
)

type Props struct {
	ID          string
	Size        int
	Color       string
	Solid       bool
	Fill        string
	Stroke      string
	StrokeWidth string
	Class       string
}

type IconFunc func(...Props) templ.Component

type SVGMap map[string]string

var DefSvgProps = Props{
	Size:  24,
	Color: "currentColor",
	Solid: false,
}

var defFamily = BOX

func SetDefIconFamily(family string) error {
	switch family {
	case BOX, LUCIDE:
		defFamily = family
		return nil
	default:

		return errors.New("invalid icon family, using bxicons")
	}
}

// Icon returns a function that generates a templ.Component for the specified icon name.
func Icon(name string, family ...string) IconFunc {
	var iconFamily = defFamily

	if len(family) > 0 {
		iconFamily = family[0]
	}

	return func(props ...Props) templ.Component {

		var p Props
		if len(props) > 0 {
			p = props[0]
		}

		cacheKey := fmt.Sprintf("%s|s:%d|c:%s|f:%s|sk:%s|sw:%s|cl:%s", name, p.Size, p.Color, p.Fill, p.Stroke, p.StrokeWidth, p.Class)

		return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			iconMutex.RLock()
			svg, cached := iconContents[cacheKey]
			iconMutex.RUnlock()

			if cached {
				_, err = w.Write([]byte(svg))
				return err
			}

			// Not cached, generate it
			// The actual generation now happens once and is cached.
			generatedSvg, err := generateSVG(name, iconFamily, p) // p (Props) is passed to generateSVG
			if err != nil {
				// Provide more context in the error message
				return fmt.Errorf("failed to generate svg for icon '%s' with props %+v: %w", name, p, err)
			}
			// log.Println(generatedSvg)

			iconMutex.Lock()
			iconContents[cacheKey] = generatedSvg
			iconMutex.Unlock()

			_, err = w.Write([]byte(generatedSvg))
			return err
		})
	}
}

func generateSVG(name string, family string, props Props) (string, error) {
	// Get the raw, inner SVG content for the icon name from our internal data map.
	content, err := getIconContent(name, family) // This now reads from internalSvgData
	if err != nil {
		log.Println(err.Error())
		return "", err // Error from getIconContent already includes icon name
	}

	size := props.Size
	if size <= 0 {
		size = 24 // Default size
	}

	stroke := props.Stroke
	if stroke == "" {
		stroke = props.Color // Fallback to Color if Stroke is not set
	}
	if stroke == "" && props.Color == "" {
		stroke = "currentColor" // Default stroke color
	}

	strokeWidth := props.StrokeWidth
	if strokeWidth == "" {
		strokeWidth = "1" // Default stroke width
		if family == LUCIDE {
			strokeWidth = "2"
		}
	}

	fill := props.Fill
	if fill == "" {
		fill = "none" // Default fill
	}

	if props.Solid || strings.Contains(name, "bxs") {
		fill = stroke
	}

	// Construct the final SVG string.
	// The data-lucide attribute helps identify these as Lucide icons if needed.
	var res = "<svg"
	if props.ID != "" {
		res += " id=" + props.ID
	}
	res += fmt.Sprintf(" xmlns=\"http://www.w3.org/2000/svg\" width=\"%d\" height=\"%d\" viewBox=\"0 0 24 24\" style=\"fill:%s; stroke:%s; stroke-width:%s;\" class=\"%s\" data-lucide=\"icon\">%s</svg>",
		size, size, fill, stroke, strokeWidth, props.Class, content)

	// fmt.Println(res)
	return res, nil
}

// getIconContent retrieves the raw inner SVG content for a given icon name.
// It reads from the pre-generated SvgData map from bxIconSvgData.go, and lucideIconSvgData.go.
func getIconContent(name string, family string) (string, error) {
	content := ""
	exists := false
	switch family {
	case BOX:
		content, exists = bxIconSvgData[name]
	case LUCIDE:
		content, exists = lucideSvgData[name]
	case CUSTOM:
		content, exists = customSvgData[name]
	default:
		content, exists = bxIconSvgData[name]
	}
	if !exists {
		return "", fmt.Errorf("icon '%s' not found in %s svg data map", name, family)
	}
	return content, nil
}

var customSvgData = SVGMap{}

func AddSvg(name, content string) {
	customSvgData[name] = content
}

// GetCustomSvgsFromFolder scans the selected folder for .svg files
func GetCustomSvgsFromFolder(folderPath string) (SVGMap, error) {
	// Initialize the map
	customSvgData = make(SVGMap)

	// Define the regex to find all <path> tags (including self-closing and multi-line)
	// (?s) makes '.' match newlines, allowing multi-line matches.
	regex := regexp.MustCompile(`(?s)<path[^>]*?>.*?</path>|<path[^>]*?/>`)

	// Read all files from the specified folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", folderPath, err)
	}

	// Loop through each file to build the map entries
	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		if !strings.HasSuffix(strings.ToLower(file.Name()), ".svg") {
			fmt.Printf("Skipping non-SVG file: %s\n", file.Name())
			continue // Skip this file if it's not an SVG
		}

		filePath := filepath.Join(folderPath, file.Name())
		key := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) // The key is the filename WITHOUT the extension

		// Read the entire file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Warning: Could not read file %s: %v\n", filePath, err)
			continue // Skip this file and proceed to the next
		}

		// Find all matches of the regex in the content
		matches := regex.FindAllString(string(content), -1)

		// Join all found <path> tags together into a single string
		extractedPaths := strings.Join(matches, "")

		// Add the entry to the map
		customSvgData[key] = extractedPaths
	}

	return customSvgData, nil
}

// WriteIconMapToFile writes a given SVGMap to a file in JSON format.
// The data will be indented for readability.
func WriteIconMapToFile(filePath string) error {
	jsonData, err := json.MarshalIndent(customSvgData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal map to JSON: %w", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON data to file %s: %w", filePath, err)
	}

	return nil
}

// ReadIconMapFromFile reads a JSON file from the given path and unmarshals it
// into a SVGMap.
func ReadIconMapFromFile(filePath string) (SVGMap, error) {
	// Read the entire file content
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Declare a variable of the target type (SVGMap)
	var loadedMap SVGMap

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(jsonData, &loadedMap)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data from %s: %w", filePath, err)
	}

	return loadedMap, nil
}

func SetCustomSVGMap(svgs SVGMap) error {
	if len(svgs) < 1 {
		return errors.New("svgs is empty. Please provide a map of SVG Files using ReadIconMapFromFile or GetCustomSVGsFromFolder")
	}
	customSvgData = svgs

	return nil
}
