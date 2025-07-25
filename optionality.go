package faraday

type optionality uint8

const (
	Required optionality = iota
	Optional
	Conditional
	Unused
)

func fromString(s string) optionality {
	switch s {
	case "R":
		return Required
	case "C":
		return Conditional
	case "X":
		return Unused
	case "O":
		return Required
	default:
		return Optional
	}
}

func canInt8(n int) bool {
	return n >= 0 && n < 256
}
