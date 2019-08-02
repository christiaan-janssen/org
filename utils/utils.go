package utils

import (
	"log"
)

// LogError logs errors
func LogError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
