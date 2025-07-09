package common

import "github.com/a-h/templ"

type InputProps struct {
	ID          string
	Name        string
	Type        string
	Value       string
	Label       string
	Placeholder string

	Required  bool
	Disabled  bool
	Readonly  bool
	Autofocus bool

	Autocomplete string
	Spellcheck   bool
	Tabindex     int
	Form         string
	Attrs        templ.Attributes
	Classes      string
}

type SelectOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
