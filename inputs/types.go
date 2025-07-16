package inputs

import (
	"time"

	"github.com/a-h/templ"
)

type textProps struct {
	MinLength int
	MaxLength int
	Pattern   string
}

type numberProps struct {
	Min  float64
	Max  float64
	Step float64
}

type emailProps struct {
	AcceptedDomains []string
}

type addressProps struct {
	DefCountry string
	DefCity    string
	DefState   string
}

type option struct {
	Label string
	Value string
}

func NewOption(label, value string) option {
	return option{
		Label: label,
		Value: value,
	}
}

type phoneProps struct {
	International bool
}

type optionProps struct {
	Options  []option
	Multiple bool
}

type selectProps struct {
	ShowOther  bool
	ShowClear  bool
	Searchable bool
}

type textareaProps struct {
	Rows    int
	Cols    int
	ResizeY bool
	ResizeX bool
}

type dateProps struct {
	MinDate time.Time
	MaxDate time.Time
}

type timeProps struct {
	MinTime time.Time
	MaxTime time.Time
}

type fileProps struct {
	Accept  string
	Capture string
}

type rangeProps struct {
	ListID string
	CMP    templ.Component
}
