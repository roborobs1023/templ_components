package types

type SelectOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ListByLabel []SelectOption

func (l ListByLabel) Len() int {
	return len(l)
}

func (list ListByLabel) Less(i, j int) bool {
	if list[i].Label != list[j].Label {
		return list[i].Label < list[j].Label
	}
	return list[i].Label < list[j].Label
}

func (list ListByLabel) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (o SelectOption) GetValue() string {
	return o.Value
}

func (o SelectOption) GetLabel() string {
	return o.Label
}

func GetIndexByValue(value string, list []SelectOption) int {
	for i, opt := range list {
		if opt.Value == value {
			return i
		}
	}
	return -1
}

func GetIndexByLabel(Label string, list []SelectOption) int {
	for i, opt := range list {
		if opt.Label == Label {
			return i
		}
	}
	return -1
}

func SortByValue(list []SelectOption) []SelectOption {

	return []SelectOption{}
}

func GetValueByLabel(label string, list []SelectOption) string {
	for _, opt := range list {
		if opt.Label == label {
			return opt.Value
		}
	}
	return ""
}

func GetLabelByValue(Value string, list []SelectOption) string {
	for _, opt := range list {
		if opt.Value == Value {
			return opt.Label
		}
	}
	return ""
}
