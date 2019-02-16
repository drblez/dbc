package dbc

import "fmt"

// Common Logger interface for using to logging contact's related panics
type Logger interface {
	Debug(msg string)
}

// Interface for validating invariant of object
type SimpleInvariantValidator interface {
	Invariant() bool
}

// Interface for validating invariant of object with Stringer interface
type InvariantValidator interface {
	SimpleInvariantValidator
	fmt.Stringer
}
