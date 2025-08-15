package generator

import (
	"fmt"
	"strings"
)

func BoxDecorationBreak(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	if prefix != "" {
		fmt.Fprintf(st, ".%s\\:decoration-slice {\n -webkit-box-decoration-break: slice;\n -moz-box-decoration-break: slice;\n box-decoration-break: slice;\n}\n\n", prefix)
		fmt.Fprintf(st, ".%s\\:decoration-clone {\n -webkit-box-decoration-break: clone;\n -moz-box-decoration-break: clone;\n box-decoration-break: clone;\n}\n\n", prefix)
		return st.String()
	}
	st.WriteString(".decoration-slice {\n -webkit-box-decoration-break: slice;\n -moz-box-decoration-break: slice;\n box-decoration-break: slice;\n}\n\n")
	st.WriteString(".decoration-clone {\n -webkit-box-decoration-break: clone;\n -moz-box-decoration-break: clone;\n box-decoration-break: clone;\n}\n\n")
	return st.String()
}

func ClearUtilities(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"left", "right", "both", "none"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%sclear-%s {\n clear:%s;\n}\n\n", class, option, option)

	}
	return st.String()
}

func ContainerCSS(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	switch prefix {
	case "sm":
		st.WriteString(".sm\\:container {\n max-width: 640px;\n}\n\n")
	case "md":
		st.WriteString(".md\\:container {\n max-width: 768px;\n}\n\n")
	case "lg":
		st.WriteString(".lg\\:container {\n max-width: 1024px;\n}\n\n")
	case "xl":
		st.WriteString(".xl\\:container {\n max-width: 1280px;\n}\n\n")
	case "2xl":
		st.WriteString(".2xl\\:container {\n max-width: 1536px;\n}\n\n")
	default:
		st.WriteString(".container {\n width: 100%;\n}\n\n")
	}

	return st.String()

}

func DisplayUtilities(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"block", "inline", "inline-block", "inline-flex", "inline-grid", "flex", "table", "table-cell", "table-column", "table-caption", "table-column-group", "grid", "hidden", "subgrid"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%s%s {\n display:%s;\n}\n\n", class, option, option)

	}

	return st.String()
}

func FloatUtilities(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"left", "right", "none"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%sfloat-%s {\n float:%s;\n}\n\n", class, option, option)

	}
	return st.String()
}

func IsolationUtilities(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	class := "."
	if prefix != "" {
		class = fmt.Sprintf(".%s\\:", prefix)
	}
	fmt.Fprintf(st, "%sisolation-%s {\n isolation:%s;\n}\n\n", class, "auto", "auto")
	fmt.Fprintf(st, ".%sisolate {\n isolation:isolate;\n}\n\n", class)
	return st.String()
}

func ObjectFit(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"contain", "cover", "fill", "none", "scale-down"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%sobject-%s {\n object-fit:%s;\n}\n\n", class, option, option)

	}
	return st.String()
}

func ObjectPosition(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"bottom", "center", "left", "left bottom", "left top", "right", "right bottom", "right top", "top"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%sobject-%s {\n object-position:%s;\n}\n\n", class, option, option)

	}
	return st.String()
}

func OverScroll(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"auto", "contain", "none"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%soverscroll-%s {\n overscroll-behavior:%s;\n}\n\n", class, option, option)
		fmt.Fprintf(st, "%soverscroll-x-%s {\n overscroll-behavior-x:%s;\n}\n\n", class, option, option)
		fmt.Fprintf(st, "%soverscroll-y-%s {\n overscroll-behavior-y:%s;\n}\n\n", class, option, option)

	}
	return st.String()
}

func OverflowUtilities(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"visible", "hidden", "scroll", "auto"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%soverflow-%s {\n overflow:%s;\n}\n\n", class, option, option)
		fmt.Fprintf(st, "%soverflow-x-%s {\n overflow-x:%s;\n}\n\n", class, option, option)
		fmt.Fprintf(st, "%soverflow-y-%s {\n overflow-y:%s;\n}\n\n", class, option, option)
	}
	return st.String()
}

func PositionUtilities(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"static", "relative", "absolute", "fixed", "sticky"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%s%s {\n position:%s;\n}\n\n", class, option, option)

	}
	return st.String()

}

func TopRightBottomLeft(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"top", "right", "bottom", "left"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}

		for i := range 65 {

			fmt.Fprintf(st, "%s%s-%d {\n %s:%vrem;\n}\n\n", class, option, i, option, float64(i)*0.25)
		}

		fmt.Fprintf(st, "%s%s-%d {\n %s:%vrem;\n}\n\n", class, option, 72, option, float64(72)*0.25)
		fmt.Fprintf(st, "%s%s-%d {\n %s:%vrem;\n}\n\n", class, option, 80, option, float64(80)*0.25)
		fmt.Fprintf(st, "%s%s-%d {\n %s:%vrem;\n}\n\n", class, option, 96, option, float64(96)*0.25)

	}

	for _, option := range []string{"inset", "inset-x", "inset-y", "top", "right", "bottom", "left"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		property := ""

		switch option {
		case "inset":
			property = "inset"
		case "inset-x":
			property = "inset-inline"
		case "inset-y":
			property = "inset-block"
		default:
			property = option
		}
		for i := range 65 {

			fmt.Fprintf(st, "%s%s-%d {\n %s:%vrem;\n}\n\n", class, option, i, property, float64(i)*0.25)
			fmt.Fprintf(st, "%s-%s-%d {\n %s:-%vrem;\n}\n\n", class, option, i, property, float64(i)*0.25)
		}

		fmt.Fprintf(st, "%s%s-%d {\n %s:%vrem;\n}\n\n", class, option, 72, property, float64(72)*0.25)
		fmt.Fprintf(st, "%s-%s-%d {\n %s:-%vrem;\n}\n\n", class, option, 72, property, float64(72)*0.25)
		fmt.Fprintf(st, "%s-%s-%d {\n %s:-%vrem;\n}\n\n", class, option, 80, property, float64(80)*0.25)
		fmt.Fprintf(st, "%s%s-%d {\n %s:%vrem;\n}\n\n", class, option, 80, property, float64(80)*0.25)
		fmt.Fprintf(st, "%s%s-%d {\n %s:%vrem;\n}\n\n", class, option, 96, property, float64(96)*0.25)
		fmt.Fprintf(st, "%s-%s-%d {\n %s:-%vrem;\n}\n\n", class, option, 96, property, float64(96)*0.25)
		fmt.Fprintf(st, "%s%s-full {\n %s:100%%;\n}\n\n", class, option, property)

		fmt.Fprintf(st, "%s%s-1\\/2 {\n %s:50%%;\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s-%s-1\\/2 {\n %s:-50%%;\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s%s-1\\/3 {\n %s:calc(100%% / 3);\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s-%s-1\\/3 {\n %s:calc(100%% / -3);\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s%s-2\\/3 {\n %s:calc(200%% / 3);\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s-%s-2\\/3 {\n %s:calc(200%% / -3);\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s%s-1\\/4 {\n %s:25%%;\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s-%s-1\\/4 {\n %s:-25%%;\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s%s-2\\/4 {\n %s:50%%;\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s-%s-2\\/4 {\n %s:-50%%;\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s%s-3\\/4 {\n %s:75%%;\n}\n\n", class, option, property)
		fmt.Fprintf(st, "%s-%s-3\\/4 {\n %s:-75%%;\n}\n\n", class, option, property)

	}
	return st.String()

}

func Visibility(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"visible", "hidden"} {
		class := "."
		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		switch option {
		case "visible":
			fmt.Fprintf(st, "%s%s {\n visibility:visible;\n}\n\n", class, option)
		case "hidden":
			fmt.Fprintf(st, "%sinvisible {\n visibility:%s;\n}\n\n", class, option)
		}
	}
	return st.String()
}

func ZIndexUtilities(prefixes ...string) string {
	st := new(strings.Builder)
	prefix := GetPrefix(prefixes...)

	for _, option := range []string{"auto", "0", "10", "20", "30", "40", "50"} {
		class := "."

		if prefix != "" {
			class = fmt.Sprintf(".%s\\:", prefix)
		}
		fmt.Fprintf(st, "%sz-index-%s {\n z-index:%s;\n}\n\n", class, option, option)

	}
	return st.String()
}

func (cg *CSSBuilder) GenerateLayoutUtilities(prefixes ...string) {
	cg.Output.WriteString(BoxDecorationBreak(prefixes...))
	cg.Output.WriteString(ClearUtilities(prefixes...))
	cg.Output.WriteString(ContainerCSS(prefixes...))
	cg.Output.WriteString(DisplayUtilities(prefixes...))
	cg.Output.WriteString(FloatUtilities(prefixes...))
	cg.Output.WriteString(IsolationUtilities(prefixes...))
	cg.Output.WriteString(ObjectFit(prefixes...))
	cg.Output.WriteString(ObjectPosition(prefixes...))
	cg.Output.WriteString(OverScroll(prefixes...))
	cg.Output.WriteString(PositionUtilities(prefixes...))
	cg.Output.WriteString(TopRightBottomLeft(prefixes...))
	cg.Output.WriteString(Visibility(prefixes...))
	cg.Output.WriteString(ZIndexUtilities(prefixes...))
}
