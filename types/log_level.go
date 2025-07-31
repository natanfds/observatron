package types

type LogLevel string

const (
	Debug LogLevel = "DEBUG"
	Info  LogLevel = "INFO"
	Warn  LogLevel = "WARN"
	Error LogLevel = "ERROR"
)

func (l LogLevel) IsValid() bool {
	switch l {
	case Debug, Info, Warn, Error:
		return true
	default:
		return false
	}
}
