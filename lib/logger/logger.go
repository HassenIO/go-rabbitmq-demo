package logger

import "log"

// OnError logs a fatal log message if any error
func OnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}
