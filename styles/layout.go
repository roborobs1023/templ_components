package styles

import "strings"

func Transition(properties ...string) string {
	text := "transition: "
	if !(len(properties) > 0) {
		return "transition-property: all"
	}

	ps := strings.Join(properties, ", ")

	text += ps

	return text
}
