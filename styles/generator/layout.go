package generator

import "fmt"

func (cg *CSSBuilder) GenerateContainerCSS(prefixes ...string) {
	prefix := GetPrefix(prefixes...)

	switch prefix {
	case "sm":
		cg.Output.WriteString(".sm\\:container {\n max-width: 640px;\n}\n\n")
	case "md":
		cg.Output.WriteString(".md\\:container {\n max-width: 768px;\n}\n\n")
	case "lg":
		cg.Output.WriteString(".lg\\:container {\n max-width: 1024px;\n}\n\n")
	case "xl":
		cg.Output.WriteString(".xl\\:container {\n max-width: 1280px;\n}\n\n")
	case "2xl":
		cg.Output.WriteString(".2xl\\:container {\n max-width: 1536px;\n}\n\n")
	default:
		cg.Output.WriteString(".container {\n width: 100%;\n}\n\n")
	}

}

func (cg *CSSBuilder) BoxDecorationBreak(prefixes ...string) {
	prefix := GetPrefix(prefixes...)

	if prefix != "" {
		fmt.Fprintf(cg.Output, ".%s\\:decoration-slice {\n -webkit-box-decoration-break: slice;\n -moz-box-decoration-break: slice;\n box-decoration-break: slice;\n}\n\n", prefix)
		fmt.Fprintf(cg.Output, ".%s\\:decoration-clone {\n -webkit-box-decoration-break: clone;\n -moz-box-decoration-break: clone;\n box-decoration-break: clone;\n}\n\n", prefix)
		return
	}
	cg.Output.WriteString(".decoration-slice {\n -webkit-box-decoration-break: slice;\n -moz-box-decoration-break: slice;\n box-decoration-break: slice;\n}\n\n")
	cg.Output.WriteString(".decoration-clone {\n -webkit-box-decoration-break: clone;\n -moz-box-decoration-break: clone;\n box-decoration-break: clone;\n}\n\n")
}
