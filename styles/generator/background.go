package generator

import (
	"fmt"
	"strings"
)

var backgroundProperties = map[string]CSSPropertyMap{
	"background-attachment": {
		"fixed": "fixed",
		"local": "scroll",
	},
	"background-clip": {
		"border":  "border-box",
		"content": "content-box",
		"padding": "padding-box",
		"text":    "text",
	},
	"bg-opacity": Opacities(),
	"background-origin": {
		"border":  "border-box",
		"content": "content-box",
		"padding": "padding-box",
	},
	"background-position": {
		"bottom":       "bottom",
		"center":       "center",
		"left":         "left",
		"right":        "right",
		"left-top":     "left top",
		"left-bottom":  "left bottom",
		"right-top":    "right top",
		"right-bottom": "right bottom",
		"top":          "top",
	},
	"background-repeat": {
		"":          "repeat",
		"no-repeat": "no-repeat",
		"x":         "repeat-x",
		"y":         "repeat-y",
		"round":     "round",
		"space":     "space",
	},
	"background-size": {
		"auto":    "auto",
		"cover":   "cover",
		"contain": "contain",
	},
	"background-image": {
		"none":           "none",
		"gradiant-to-t":  "linear-gradient(to top, var(--tc-gradient-stops))",
		"gradient-to-tr": "linear-gradient(to top right, var(--tc-gradient-stops))",
		"gradiant-to-r":  "linear-gradient(to right, var(--tc-gradient-stops))",
		"gradient-to-br": "linear-gradient(to bottom right, var(--tc-gradient-stops))",
		"gradient-to-bl": "linear-gradient(to bottom left, var(--tc-gradient-stops))",
		"gradient-to-l":  "linear-gradient(to left, var(--tc-gradient-stops))",
		"gradient-to-tl": "linear-gradient(to top left, var(--tc-gradient-stops))",
	},
}

func Opacities() CSSPropertyMap {
	m := CSSPropertyMap{}

	for i := 0; i <= 100; i += 5 {
		m[fmt.Sprintf("opacity-%v", i)] = fmt.Sprintf("%v", i/100)
	}

	return m
}

func (cg *CSSBuilder) GenerateBackgroundUtilities(prefix string, suffix string) {
	className := ""
	if prefix != "" {
		className += fmt.Sprintf("%s\\:", prefix)
	}
	for k, properties := range backgroundProperties {
		if strings.Contains(k, "background-size") || strings.Contains(k, "background-image") {
			className += "bg"
		} else if strings.Contains(k, "background") {
			className += strings.ReplaceAll(k, "background", "bg")
		} else {
			className += k
		}
		if suffix != "" {
			className += fmt.Sprintf(":%s", suffix)
		}
		for name, value := range properties {
			fmt.Fprintf(cg.Output, ".%s-%s {\n %s:%s;\n}\n\n", className, name, k, value)
		}
	}

}
