package dbc

import "fmt"

type Logger interface {
	Debug(msg string)
}

type SimpleValidator interface {
	Validate() bool
}

type Validator interface {
	SimpleValidator
	fmt.Stringer
}


