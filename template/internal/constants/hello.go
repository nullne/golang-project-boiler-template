package constants

//go:generate enumer -type=HelloDays -json -trimprefix=HelloDay -transform=lower
type HelloDays int

const (
	HelloDayOne HelloDays = iota + 1
	HelloDayTwo
)
