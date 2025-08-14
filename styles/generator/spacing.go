package generator

import (
	"fmt"
	"strings"
)

type SpacePropertyMap map[string]string

func spaceBetweenMap() SpacePropertyMap {
	m := SpacePropertyMap{}

	m["0"] = "calc(0px * var(--tc-space-x-reverse))"
	for i := 1; i < 13; i++ {
		m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%vrem ", float64(i)*0.25)
	}

	for i := 16; i < 66; i += 4 {
		m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%vrem ", float64(i)*0.25)
	}
	i := 14
	m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%vrem", float64(i)*0.25)

	i = 72

	m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%vrem", float64(i)*0.25)
	i = 80

	m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%vrem", float64(i)*0.25)

	i = 96

	m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%vrem", float64(i)*0.25)

	return m
}

func (cg *CSSBuilder) GenerateSpaceBetweenUtilities(prefixes ...string) {
	class := ""
	prefix := GetPrefix(prefixes...)

	if prefix != "" {
		class = fmt.Sprintf("%s\\:", prefix)
	}

	for key, p := range spaceBetweenMap() {
		fmt.Fprintf(cg.Output, ".%sspace-x-%s > * {\n --tc-space-x-reverse:0;\n margin-left:calc(%s * var(--tc-space-x-reverse));\n margin-right:calc(%s * calc(1 - var(--tc-space-x-reverse)));\n}\n\n", class, key, p, p)

	}

	fmt.Fprintf(cg.Output, ".%sspace-x-0\\.5 > * {\n --tc-space-x-reverse: 0;\n margin-right: calc(-0.125rem * var(--tc-space-x-reverse));\\n margin-left: calc(-0.125rem * calc(1 - var(--tc-space-x-reverse)));\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%sspace-x-1\\.5 > * {\n --tc-space-x-reverse: 0;\n margin-right: calc(-0.375rem * var(--tc-space-x-reverse));\\n margin-left: calc(-0.375rem * calc(1 - var(--tc-space-x-reverse)));\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%sspace-x-2\\.5 > * {\n --tc-space-x-reverse: 0;\n margin-right: calc(-0.625rem * var(--tc-space-x-reverse));\\n margin-left: calc(-0.625rem * calc(1 - var(--tc-space-x-reverse)));\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%sspace-x-3\\.5 > * {\n --tc-space-x-reverse: 0;\n margin-right: calc(-0.875rem * var(--tc-space-x-reverse));\\n margin-left: calc(-0.875rem * calc(1 - var(--tc-space-x-reverse)));\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%sspace-y-0\\.5 > * {\n --tc-space-y-reverse: 0;\n margin-top: calc(-0.125rem * var(--tc-space-y-reverse));\\n margin-bottom: calc(-0.125rem * calc(1 - var(--tc-space-y-reverse)));\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%sspace-y-1\\.5 > * {\n --tc-space-y-reverse: 0;\n margin-top: calc(-0.375rem * var(--tc-space-y-reverse));\\n margin-bottom: calc(-0.375rem * calc(1 - var(--tc-space-y-reverse)));\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%sspace-y-2\\.5 > * {\n --tc-space-y-reverse: 0;\n margin-top: calc(-0.625rem * var(--tc-space-y-reverse));\\n margin-bottom: calc(-0.625rem * calc(1 - var(--tc-space-y-reverse)));\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%sspace-y-3\\.5 > * {\n --tc-space-y-reverse: 0;\n margin-top: calc(-0.875rem * var(--tc-space-y-reverse));\\n margin-bottom: calc(-0.875rem * calc(1 - var(--tc-space-y-reverse)));\n}\n\n", class)

	fmt.Fprintf(cg.Output, ".%sspace-x-reverse > * {\n --tc-space-x-reverse: 1;\n}\n\n", class)

	for key, p := range spaceBetweenMap() {
		fmt.Fprintf(cg.Output, ".%sspace-y-%s > * {\n --tc-space-y-reverse:0;\n margin-top:calc(%s * var(--tc-space-y-reverse));\n margin-bottom:calc(%s * calc(1 - var(--tc-space-y-reverse)));\n}\n\n", class, key, p, p)
	}

	fmt.Fprintf(cg.Output, ".%sspace-y-reverse > * {\n --tc-space-y-reverse: 1;\n}\n\n", class)

}

func (cg *CSSBuilder) GenerateSpacingUtilityClasses(prefixes ...string) {
	prefix := GetPrefix(prefixes...)
	properties := []string{
		"padding",
		"margin",
	}

	for _, property := range properties {

		var classes map[string]string
		switch property {
		case "padding":
			classes = map[string]string{"p": "padding", "pl": "padding-left", "pr": "padding-right", "pt": "padding-top", "pb": "padding-bottom", "px": "padding-x", "py": "padding-y"}
		case "margin":
			classes = map[string]string{"m": "margin", "ml": "margin-left", "mr": "margin-right", "mt": "margin-top", "mb": "margin-bottom", "mx": "margin-x", "my": "margin-y"}
		default:
			return
		}

		for name, value := range spacingMap() {
			for prop, val := range classes {
				className := fmt.Sprintf(".%s-%s", prop, name)
				if prefix != "" {
					className = fmt.Sprintf(".%s\\:%s-%s", prefix, prop, name)
				}

				switch true {
				case strings.Contains(prop, "x"):
					fmt.Fprintf(cg.Output, "%s {\n %s-inline:%s;\n", className, property, value)
				case strings.Contains(prop, "y"):
					fmt.Fprintf(cg.Output, "%s {\n %s-block:%s;\n", className, property, value)
				default:
					fmt.Fprintf(cg.Output, "%s {\n %s:%s;\n}\n\n", className, val, value)
				}

			}
		}
	}

}
