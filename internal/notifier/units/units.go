package telegram

type FrequencyUnit int

const (
	Seconds FrequencyUnit = iota
	Minutes
	Hours
)

func (fu FrequencyUnit) String() string {
	switch fu {
	case Seconds:
		return "s"
	case Minutes:
		return "m"
	case Hours:
		return "h"
	default:
		return ""
	}
}
