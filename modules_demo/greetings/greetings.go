package greetings

import "fmt"

// Hello returns a greeting
func Hello(name string) string {
	if name == "" {
		return "Did you forget something?"
	}
	message := fmt.Sprintf("Howdy, %v! Welcome to the party!", name)
	return message
}
