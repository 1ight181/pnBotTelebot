package units

import "fmt"

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

func ParseFrequencyUnit(s string) (FrequencyUnit, error) {
	switch s {
	case "s":
		return Seconds, nil
	case "m":
		return Minutes, nil
	case "h":
		return Hours, nil
	default:
		return 0, fmt.Errorf("недопустимое значение для FrequencyUnit: %s", s)
	}
}
