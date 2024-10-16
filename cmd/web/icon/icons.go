package icon

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
	if p.Size == "" {
		p.Size = "1.5rem"
	}

	if p.Color == "" {
		p.Color = "currentColor"
	}
	return true
}
