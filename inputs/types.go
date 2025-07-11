package inputs

import "time"

type textProps struct {
	MinLength int
	MaxLength int
	Pattern   string
}

type addressProps struct {
	DefCountry string
	DefCity    string
	DefState   string
}

type checkboxProps struct {
	Multiple bool
}

type emailProps struct {
	AcceptedDomains []string
}

type numberProps struct {
	Min   float64
	Max   float64
	Value float64
	Step  float64
}

type Option struct {
	Label string
	Value string
}

func NewOption(label, value string) Option {
	return Option{
		Label: label,
		Value: value,
	}
}

type phoneProps struct {
	International bool
}

type optionProps struct {
	Options      []Option
	DefaultValue string
}

type radioProps struct {
}

type selectProps struct {
	ShowOther  bool
	ShowClear  bool
	Searchable bool
}

type textareaProps struct {
	Rows int
	Cols int
}

type dateProps struct {
	MinDate time.Time
	MaxDate time.Time
}

type timeProps struct {
	MinTime time.Time
	MaxTime time.Time
}
