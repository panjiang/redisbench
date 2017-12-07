package wares

import (
	"fmt"
	"log"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + string(bytes))
}

// InitLogger : Initial logger with a new formatting
func InitLogger() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}
