package styles

import (
	"fmt"
	"slices"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

func MaxWidth(value string) string {
	return fmt.Sprintf("max-width: %v", value)
}

func MinWidth(value string) string {
	return fmt.Sprintf("min-width: %v", value)
}

type Size string

const (
	SM   Size = "sm"
	XS   Size = "xs"
	MD   Size = "md"
	LG   Size = "lg"
	XL   Size = "xl"
	XXL  Size = "2xl"
	XXXL Size = "3xl"
	ZERO Size = "0"
)

var Sizes = map[Size]string{
	XS:   "0.25rem",
	SM:   "0.5rem",
	MD:   "1rem",
	LG:   "1.25rem",
	XL:   "1.5rem",
	XXL:  "2rem",
	XXXL: "2.5rem",
	ZERO: "0",
}

func (s Size) Value() string {
	return Sizes[s]
}

type loc string

const (
	locTop    loc = "top"
	locLeft   loc = "left"
	locRight  loc = "right"
	locBottom loc = "bottom"
	locInline loc = "inline"
	locBlock  loc = "block"
)

var locations = []loc{locTop, locBottom, locLeft, locRight, locInline, locBlock}

func Padding(value string, locs ...string) string {
	var location = ""
	if len(locs) > 0 {
		location = locs[0]
	}

	if location != "" && slices.Contains(locations, loc(location)) {
		property := "padding-" + location
		return getString(property, value)
	}

	return getString("padding", value)
}

func Margin(value string, locs ...string) string {
	var location = ""
	if len(locs) > 0 {
		location = locs[0]
	}

	if location != "" && slices.Contains(locations, loc(location)) {
		property := "margin-" + location
		return getString(property, value)
	}

	return getString("margin", value)
}

func getString(property string, value string) string {
	size := Size(value).Value()
	if size == "" {
		size = value
	}
	return property + ": " + size
}

func Inset(value string) string {

	return fmt.Sprintf("inset: %v", value)
}

func GenerateThemeCSSVariables(baseHexColor string) (string, error) {
	// Parse the base color
	baseColor, err := colorful.Hex(baseHexColor)
	if err != nil {
		return "", fmt.Errorf("invalid base hex color: %w", err)
	}

	// Convert to HSL for easier manipulation
	h, s, l := baseColor.Hsl()

	// Define the color scheme variations
	// You can customize these values heavily based on your design needs.
	// H (Hue): Controls the "color" (red, green, blue, etc.)
	// S (Saturation): Controls the intensity of the color (0 = grayscale, 1 = vivid)
	// L (Lightness): Controls how bright or dark the color is (0 = black, 1 = white)

	// Primary colors
	primaryLight := colorful.Hsl(h, s, l+(1-l)*0.2) // Lighter by 20% of remaining lightness
	primaryDark := colorful.Hsl(h, s, l*0.8)        // Darker by 20%

	// Text/Background colors derived from base lightness
	// Adjust saturation for less vibrant background elements if needed
	backgroundColor := colorful.Hsl(h, s*0.2, 0.95) // Very light, slightly desaturated
	textColor := colorful.Hsl(h, s*0.5, 0.20)       // Dark, slightly desaturated

	// Accent/Complementary color (optional, but good for visual interest)
	// Shift hue by 180 degrees for complementary
	accentColor := colorful.Hsl(h+180, s, l)

	// Neutral colors (often grayscale, can be slightly tinted towards the base hue)
	grayLight := colorful.Hsl(h, s*0.05, 0.90) // Slightly tinted light gray
	grayMedium := colorful.Hsl(h, s*0.05, 0.50)
	grayDark := colorful.Hsl(h, s*0.05, 0.15)

	// Error/Success/Warning colors (often fixed, but can be slightly tinted)
	// For simplicity, we'll use fixed values here, but you could derive them
	// from red, green, yellow hues with adjusted saturation/lightness.
	errorColor, _ := colorful.Hex("#dc3545")   // Bootstrap red
	successColor, _ := colorful.Hex("#28a745") // Bootstrap green
	warningColor, _ := colorful.Hex("#ffc107") // Bootstrap yellow

	// Build the CSS variables string
	var cssVars strings.Builder
	cssVars.WriteString(":root {\n")
	cssVars.WriteString(fmt.Sprintf("  --base-color: %s;\n", baseColor.Hex()))
	cssVars.WriteString(fmt.Sprintf("  --primary-light: %s;\n", primaryLight.Hex()))
	cssVars.WriteString(fmt.Sprintf("  --primary-dark: %s;\n", primaryDark.Hex()))

	cssVars.WriteString(fmt.Sprintf("  --background-color: %s;\n", backgroundColor.Hex()))
	cssVars.WriteString(fmt.Sprintf("  --text-color: %s;\n", textColor.Hex()))

	cssVars.WriteString(fmt.Sprintf("  --accent-color: %s;\n", accentColor.Hex()))

	cssVars.WriteString(fmt.Sprintf("  --gray-light: %s;\n", grayLight.Hex()))
	cssVars.WriteString(fmt.Sprintf("  --gray-medium: %s;\n", grayMedium.Hex()))
	cssVars.WriteString(fmt.Sprintf("  --gray-dark: %s;\n", grayDark.Hex()))

	cssVars.WriteString(fmt.Sprintf("  --error-color: %s;\n", errorColor.Hex()))
	cssVars.WriteString(fmt.Sprintf("  --success-color: %s;\n", successColor.Hex()))
	cssVars.WriteString(fmt.Sprintf("  --warning-color: %s;\n", warningColor.Hex()))
	cssVars.WriteString("}\n")

	return cssVars.String(), nil
}
