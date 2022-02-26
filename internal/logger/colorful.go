package logger

import (
	"fmt"
	"log"
	"os"
)

// Colorful holds 3 types of logger that prints in 3 different colors
type Colorful struct {
	// info prints info messages
	info *log.Logger

	// warning prints warning messages
	warning *log.Logger

	// err prints error messages
	err *log.Logger
}

// Option defines ways of building a logger
type Option func(*Colorful) error

// NewColorful creates a new Colorful
func NewColorful(info, warning, err Color, options ...Option) (*Colorful, error) {
	colorful := &Colorful{
		info:    log.New(os.Stdout, string(info)+"[INFO]\t", log.Ldate|log.Ltime),
		warning: log.New(os.Stdout, string(warning)+"[WARNING]\t", log.Ldate|log.Ltime),
		err:     log.New(os.Stdout, string(err)+"[ERROR]\t", log.Ldate|log.Ltime),
	}

	for _, option := range options {
		err := option(colorful)
		if err != nil {
			return nil, fmt.Errorf("failed to configure logger: %w", err)
		}
	}

	return colorful, nil
}

// Info prints relevant information
func (c *Colorful) Info(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	c.info.Printf(message)
}

// Warning raises a warning
func (c *Colorful) Warning(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	c.info.Printf(message)
}

// Error alerts about an error
func (c *Colorful) Error(format string, v ...interface{}) {
	c.err.Printf(format, v...)
}

// Error alerts about a fatal error and exits the application
func (c *Colorful) Fatal(format string, v ...interface{}) {
	c.err.Printf(format, v...)
	os.Exit(1)
}
