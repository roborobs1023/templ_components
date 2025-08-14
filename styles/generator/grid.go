package generator

import (
	"fmt"
	"strconv"
)

var gridProperties = map[string]CSSPropertyMap{
	"grid-template-columns": {
		"none":        "none",
		"auto":        "auto",
		"min-content": "min-content",
		"max-content": "max-content",
		"fit-content": "fit-content(minmax(0, 1fr))", // Common to use with 1fr
		"1":           "repeat(1, minmax(0, 1fr))",
		"2":           "repeat(2, minmax(0, 1fr))",
		"3":           "repeat(3, minmax(0, 1fr))",
		"4":           "repeat(4, minmax(0, 1fr))",
		"5":           "repeat(5, minmax(0, 1fr))",
		"6":           "repeat(6, minmax(0, 1fr))",
		"7":           "repeat(7, minmax(0, 1fr))",
		"8":           "repeat(8, minmax(0, 1fr))",
		"9":           "repeat(9, minmax(0, 1fr))",
		"10":          "repeat(10, minmax(0, 1fr))",
		"11":          "repeat(11, minmax(0, 1fr))",
		"12":          "repeat(12, minmax(0, 1fr))",
		"dense":       "repeat(auto-fit, minmax(0, 1fr))", // Common for auto-filling grids
	},
	"grid-template-rows": {
		"none":        "none",
		"auto":        "auto",
		"min-content": "min-content",
		"max-content": "max-content",
		"1":           "repeat(1, minmax(0, 1fr))",
		"2":           "repeat(2, minmax(0, 1fr))",
		"3":           "repeat(3, minmax(0, 1fr))",
		"4":           "repeat(4, minmax(0, 1fr))",
		"5":           "repeat(5, minmax(0, 1fr))",
		"6":           "repeat(6, minmax(0, 1fr))",
	},
	"grid-auto-rows": {
		"auto":        "auto",
		"min-content": "min-content",
		"max-content": "max-content",
		"fr":          "minmax(0, 1fr)",
	},
	"grid-auto-cols": {
		"auto":        "auto",
		"min-content": "min-content",
		"max-content": "max-content",
		"fr":          "minmax(0, 1fr)",
	},
	"row-auto": {
		"auto": "auto",
	},
	"row-span":  rowSpans(),
	"row-start": rowStartsAndEnds(),
	"row-end":   rowStartsAndEnds(),

	"grid-column-start": colSpans(),
	"grid-column-end":   colSpans(true),
	"justify-self": {
		"auto":    "auto",
		"start":   "start",
		"end":     "end",
		"center":  "center",
		"stretch": "stretch",
	},
	"align-items": {
		"start":    "flex-start",
		"end":      "flex-end",
		"center":   "center",
		"baseline": "baseline",
		"stretch":  "stretch",
	},
	"place-self": {
		"auto":    "auto auto",
		"start":   "start start",
		"end":     "end end",
		"center":  "center center",
		"stretch": "stretch stretch",
	},
	"gap":        gapValues(),
	"row-gap":    gapValues(),
	"column-gap": gapValues(),
}

func colSpans(end ...bool) CSSPropertyMap {
	m := CSSPropertyMap{
		"auto": "auto",
		"full": "1 / -1", // Common for spanning full grid
	}
	for i := 1; i <= 12; i++ {
		s := fmt.Sprintf("%d", i)
		m[s] = s                   // For line number, e.g., grid-column-start-1
		m["span-"+s] = "span " + s // For spanning, e.g., grid-column-span-2
	}
	if len(end) > 0 {
		m["-1"] = "-1"
	}
	return m
}

func rowSpans() CSSPropertyMap {
	m := CSSPropertyMap{}

	for i := range 7 {
		m[fmt.Sprintf("%v", i)] = fmt.Sprintf("span %v / span %v", i, i)
	}

	m["full"] = "1 / -1"
	return m
}

func rowStartsAndEnds() CSSPropertyMap {
	m := CSSPropertyMap{}
	for i := range 8 {
		m[fmt.Sprintf("%v", i)] = fmt.Sprintf("%v", i)
	}
	m["auto"] = "auto"
	return m
}

func gapValues() CSSPropertyMap {
	m := CSSPropertyMap{}

	for i := range 13 {
		val := float64(i) * .25

		valStr := strconv.FormatFloat(val, 'f', 3, 64) + "rem"

		m[fmt.Sprintf("%v", i)] = valStr
	}
	m["14"] = "3.5rem"
	m["16"] = "4rem"
	for i := 20; i < 68; i += 4 {
		val := float64(i) * .25

		valStr := strconv.FormatFloat(val, 'f', 3, 64) + "rem"

		m[fmt.Sprintf("%v", i)] = valStr
	}

	m["72"] = "18rem"
	m["80"] = "20rem"
	m["96"] = "24rem"

	m["px"] = "1px"
	m["0_5"] = "0.125rem"
	m["1_5"] = "0.375rem"
	m["2_5"] = "0.625rem"
	m["4_5"] = "0.875rem"

	return m
}
