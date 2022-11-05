package log

type Level int32

const (
	// DebugLevel is the lowest level of logging.
	// Debug logs are intended for debugging and development purposes.
	DebugLevel Level = iota

	// InfoLevel is used for general informational log messages.
	InfoLevel

	// WarnLevel is used for undesired but relatively expected events,
	// which may indicate a problem.
	WarnLevel

	// ErrorLevel is used for undesired and unexpected events that
	// the program can recover from.
	ErrorLevel

	// FatalLevel is used for undesired and unexpected events that
	// the program cannot recover from.
	FatalLevel
)

// String is part of the fmt.Stringer interface.
//
// Used for testing and debugging purposes.
func (self Level) String() string {

	switch self {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "unknown"
	}
}
