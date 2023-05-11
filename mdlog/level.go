package mdlog

type Level int

const (
	Fatal Level = iota
	Error Level = iota
	Warn  Level = iota
	Info  Level = iota
	Debug Level = iota
)

var levels = map[Level]string{
	Fatal: "fatal",
	Error: "error",
	Warn:  "warn",
	Info:  "info",
	Debug: "debug",
}

func Levels() map[Level]string {
	return levels
}

func (l Level) String() string {
	return levels[l]
}
