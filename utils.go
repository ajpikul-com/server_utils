package server_utils

import (
	"github.com/ajpikul-com/ilog"
)

// Set a global logger for the library
var defaultLogger ilog.LoggerInterface

// Establish a default logger
func init() {
	if defaultLogger == nil {
		defaultLogger = new(ilog.EmptyLogger)
	}
}

// Allow calling program to change default logger
func SetDefaultLogger(newLogger ilog.LoggerInterface) {
	defaultLogger = newLogger
	defaultLogger.Info("Default Logger Set")
}
