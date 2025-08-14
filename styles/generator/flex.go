package generator

import (
	"fmt"
	"strings"
)

var flexProperties = map[string]CSSPropertyMap{
	"flex-direction": {
		"row":         "row",
		"row-reverse": "row-reverse",
		"col":         "column",
		"col-reverse": "column-reverse",
	},
	"flex-wrap": {
		"nowrap":       "nowrap",
		"wrap":         "wrap",
		"wrap-reverse": "wrap-reverse",
	},
	"justify-content": {
		"start":   "flex-start",
		"end":     "flex-end",
		"center":  "center",
		"between": "space-between",
		"around":  "space-around",
		"evenly":  "space-evenly",
	},

	"align-content": {
		"start":   "flex-start",
		"end":     "flex-end",
		"center":  "center",
		"between": "space-between",
		"around":  "space-around",
		"evenly":  "space-evenly",
		"stretch": "stretch",
	},
	"align-self": {
		"auto":     "auto", // Default, uses align-items value of parent
		"start":    "flex-start",
		"end":      "flex-end",
		"center":   "center",
		"baseline": "baseline",
		"stretch":  "stretch",
	},
	"flex-shrink": {
		"shrink":   "1",
		"shrink-0": "0",
	},
	"flex-grow": {
		"grow":   "1",
		"grow-0": "0",
	},
	"flex-basis": {
		"auto":    "auto", // Default, size based on content or width/height
		"0":       "0",    // Zero initial size (flex item will rely on grow/shrink)
		"full":    "100%", // 100% of container's main axis size
		"1_2":     "50%",  // 50%
		"1_3":     "33.333333%",
		"1_4":     "25%",
		"3_4":     "75%",
		"content": "content",
		"1":       "0.25rem",
		"2":       "0.5rem",
		"3":       "0.75rem",
		"4":       "1rem",
	},
	"order": generateOrders(),
}

func generateOrders() CSSPropertyMap {
	m := CSSPropertyMap{}
	m["first"] = "-9999"
	m["last"] = "9999"
	m["none"] = "0"
	for i := 1; i < 12; i++ {
		m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%v", i)
	}

	return m
}

func (cg *CSSBuilder) generateFlexUtilityClasses(prefixes ...string) {
	prefix := GetPrefix(prefixes...)

	for key, p := range flexProperties {
		for name, value := range p {

			if strings.Contains(key, "wrap") || strings.Contains(key, "direction") || strings.Contains(key, "grow") || strings.Contains(key, "shrink") {
				if prefix != "" {
					fmt.Fprintf(cg.Output, ".%sflex-%s {\n %s:%s;\n}\n\n", prefix, name, key, value)
				} else {
					fmt.Fprintf(cg.Output, ".flex-%s {\n %s:%s;\n}\n\n", name, key, value)
				}
				continue
			}

			if prefix != "" {
				fmt.Fprintf(cg.Output, ".%s%s-%s {\n %s:%s;\n}\n\n", prefix, key, name, key, value)
				continue
			}
			fmt.Fprintf(cg.Output, ".%s-%s {\n %s:%s;\n}\n\n", key, name, key, value)

		}
	}

	for i := 0; i < 10; i++ {
		if prefix != "" {
			fmt.Fprintf(cg.Output, ".%sorder-%d {\n order:%d;\n}\n\n", prefix, i, i)
			continue
		}
		fmt.Fprintf(cg.Output, ".order-%d {\n order:%d;\n}\n\n", i, i)
	}

}
