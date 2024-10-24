package icons

const (
	vFill    = "fill"
	vRounded = "rounded"
	vSharp   = "sharp"
)

type IconProps struct {
	Class   string
	Size    string
	Fill    bool
	Color   string
	Variant string
}

func (p *IconProps) checkDefaults() bool {

	if CustomDefaults != nil {
		return p.setCustomDefaults()
	}
	if p.Size == "" {
		p.Size = "1.5rem"
	}

	if p.Color == "" {
		p.Color = "currentColor"
	}
	return true
}

type IconConfig struct {
	Size    string
	Fill    bool
	Color   string
	Variant string
}

var CustomDefaults *IconConfig

func SetCustomDefaults(c *IconConfig) {
	CustomDefaults = c
}

func (p *IconProps) setCustomDefaults() bool {
	p.Size = CustomDefaults.Size
	p.Color = CustomDefaults.Color
	p.Fill = CustomDefaults.Fill
	p.Variant = CustomDefaults.Variant

	return true
}
