package console

type LogLevel int

const (
	LogLevelUndefined LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelTrace
)

func ToLevel(levelName string) LogLevel {
	switch levelName {
	case "DEBUG":
		return LogLevelDebug
	case "INFO":
		return LogLevelInfo
	case "WARN":
		return LogLevelWarn
	case "ERROR":
		return LogLevelError
	case "FATAL":
		return LogLevelFatal
	case "TRACE":
		return LogLevelTrace
	default:
		return LogLevelUndefined
	}
}

func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	case LogLevelTrace:
		return "TRACE"
	default:
		return "UNDEFINED"
	}
}
