package generator

import (
	"fmt"
	"strconv"
	"strings"
)

var innerSizeMap = map[string]CSSPropertyMap{
	"width": widthValues(),
	"min-width": {
		"0":    "0px",
		"full": "100%",
		"min":  "min-content",
		"max":  "max-content",
	},
	"max-width": {
		"0":          "0",
		"none":       "none",
		"xs":         "20rem",
		"sm":         "24rem",
		"md":         "28rem",
		"lg":         "32rem",
		"xl":         "36rem",
		"2xl":        "42rem",
		"3xl":        "48rem",
		"4xl":        "56rem",
		"5xl":        "64rem",
		"6xl":        "72rem",
		"7xl":        "80rem",
		"full":       "100%",
		"max":        "max-content",
		"min":        "min-content",
		"prose":      "65ch",
		"screen-sm":  "640px",
		"screen-md":  "768px",
		"screen-lg":  "1024px",
		"screen-xl":  "1280px",
		"screen-2xl": "1536px",
	},
	"min-height": {
		"0":              "0",
		"full":           "100%",
		"screen":         "100vh",
		"dynamic-screen": "100dvh",
	},
	"height":    heightValues(),
	"max-heiht": heightValues(),
}

func widthValues() CSSPropertyMap {
	m := CSSPropertyMap{}
	for i := 0; i < 17; i++ {
		val := float64(i) * .25

		valStr := strconv.FormatFloat(val, 'f', 3, 64) + "rem"

		m[fmt.Sprintf("%v", i)] = valStr
	}

	for i := 20; i < 66; i += 4 {
		val := i / 4

		m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%vrem", val)
	}
	m["72"] = "18rem"
	m["80"] = "20rem"
	m["96"] = "24rem"
	m["auto"] = "auto"
	m["px"] = "1px"
	m["0\\.5"] = "0.125rem"
	m["1\\.5"] = "0.375rem"
	m["2\\.5"] = "0.625rem"
	m["3\\.5"] = "0.875rem"

	return m
}

func heightValues() CSSPropertyMap {
	m := CSSPropertyMap{}
	for i := 0; i < 17; i++ {
		val := float64(i) * .25

		valStr := strconv.FormatFloat(val, 'f', 3, 64) + "rem"

		m[fmt.Sprintf("%v", i)] = valStr
	}

	for i := 20; i < 66; i += 4 {
		val := i / 4

		m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%vrem", val)
	}
	m["72"] = "18rem"
	m["80"] = "20rem"
	m["96"] = "24rem"
	m["auto"] = "auto"
	m["px"] = "1px"
	m["0\\.5"] = "0.125rem"
	m["1\\.5"] = "0.375rem"
	m["2\\.5"] = "0.625rem"
	m["3\\.5"] = "0.875rem"
	m["full"] = "100%"
	m["screen"] = "100vh"
	m["dynamic-screen"] = "100dvh"
	return m
}

func (cg *CSSBuilder) GenerateSizing(prefixes ...string) {
	prefix := GetPrefix(prefixes...)

	class := ""

	if prefix != "" {
		class = fmt.Sprintf("%s\\:", prefix)
	}

	for key, p := range innerSizeMap {
		switch true {
		case strings.Contains(key, "width"):
			class += "w"
		}
		for name, value := range p {
			fmt.Fprintf(cg.Output, ".%s-%s {\n %s:%s;\n}\n\n", class, name, key, value)
		}
	}
}
