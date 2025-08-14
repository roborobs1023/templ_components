package generator

import "fmt"

func (cg *CSSBuilder) GenerateTableUtilities(prefixes ...string) {
	class := ""
	prefix := GetPrefix(prefixes...)

	if prefix != "" {
		class = prefix + "\\:"
	}
	fmt.Fprintf(cg.Output, ".%sborder-collapse {\n border-collapse:collapse;\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%sborder-separate {\n border-collapse:separate;\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%stable-auto {\n table-layout:auto;\n}\n\n", class)
	fmt.Fprintf(cg.Output, ".%stable-fixed {\n table-layout:fixed;\n}\n\n", class)
}
