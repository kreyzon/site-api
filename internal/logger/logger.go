package logger

// Logger defines methods for logging useful information
type Logger interface {
	// Info prints relevant information
	Info(format string, v ...interface{})

	// Warning raises a warning
	Warning(format string, v ...interface{})

	// Error alerts about an error
	Error(format string, v ...interface{})

	// Error alerts about a fatal error and exits the application
	Fatal(format string, v ...interface{})
}
