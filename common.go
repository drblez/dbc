package dbc

import "fmt"

// Common Logger interface for using to logging contact's related panics
type Logger interface {
	Debug(msg string)
}

// Interface for validating objects data
type SimpleValidator interface {
	Validate() bool
}

// Interface for validating objects data with Stringer interface
type Validator interface {
	SimpleValidator
	fmt.Stringer
}
