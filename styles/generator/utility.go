package generator

func GetPrefix(prefixes ...string) string {
	if len(prefixes) > 0 {
		return prefixes[0]
	}
	return ""
}
