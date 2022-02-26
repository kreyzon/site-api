package logger

// Color defines a color type for loggers
type Color string

// Color codes to change the color of the log message on terminal
// Using WHITE(DEFAULT TERMINAL COLOR) color to show INFO LOGS
// Using YELLOW color to show WARNING LOGS
// Using RED color to show ERROR LOGS
const (
	White  Color = "\033[0m"
	Yellow Color = "\033[33m"
	Red    Color = "\033[31m"
)
