package icons

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/a-h/templ"
)

// iconContents caches the fully generated SVG strings from icons that have been used, keyed by the composite key of name and props to handle different styles.
var (
	iconContents = make(map[string]string)
	iconMutex    sync.RWMutex
)

const (
	BX     = "bxicons"
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

// Icon returns a function that generates a templ.Component for the specified icon name.
func Icon(name string, family ...string) func(...Props) templ.Component {
	var iconFamily = "bxicons"

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

	if props.Solid {
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
// It reads from the pre-generated internalSvgData map from icon_data.go.
func getIconContent(name string, family string) (string, error) {
	content := ""
	exists := false
	switch family {
	case BX:
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

var customSvgData = map[string]string{}

func AddSvg(name, content string) {
	customSvgData[name] = content
}
