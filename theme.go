package templ_components

import (
	"fmt"
	"strings"
)

type DesignToken struct {
	Name  string
	Value string
}

type UtilityGenerator struct {
	NamePrefix  string
	CSSProperty string
	Tokens      []DesignToken
}

// GenerateCSS generates utility CSS rules based on the generator's configuration.
func (ug UtilityGenerator) GenerateCSS() string {
	var sb strings.Builder
	for _, token := range ug.Tokens {
		// Example: .text-blue { color: #3498db; }
		sb.WriteString(fmt.Sprintf(".%s-%s {\n", ug.NamePrefix, token.Name))
		sb.WriteString(fmt.Sprintf("  %s: %s;\n", ug.CSSProperty, token.Value))
		sb.WriteString("}\n\n")
	}
	return sb.String()
}

var DefColors = []DesignToken{
	{"primary", "#0c1a2e"},
	{"secondary", "#545454"},
	{"tertiary", "#"},
}
