//+build !without_dbc

package dbc

import (
	"errors"
	"fmt"
	"runtime"
)

type DBC struct {
	log Logger
}

func New() *DBC {
	return &DBC{
		log: nil,
	}
}

func NewWithLogger(log Logger) *DBC {
	return &DBC{
		log: log,
	}
}

func _panic(dbc *DBC, function string, condition bool, msgs ...interface{}) {
	if condition {
		return
	}
	caller := ""
	pc, file, line, ok := runtime.Caller(2)
	funcInfo := runtime.FuncForPC(pc)
	if ok {
		caller = fmt.Sprintf("func %s in %s at %d",
			funcInfo.Name(),
			file,
			line,
		)
	} else {
		caller = "unknown caller"
	}
	str := fmt.Sprintf("%s: %s", function, caller)
	if dbc.log != nil {
		dbc.log.Debug(str)
	}
	for i := range msgs {
		s := fmt.Sprintf("[%d] %+v", i, msgs[i])
		str += "\n" + s
		if dbc.log != nil {
			dbc.log.Debug(s)
		}
	}
	panic(errors.New(str))
}

func (dbc *DBC) Require(condition bool, msgs ...interface{}) {
	_panic(dbc, "Require", condition, msgs)
}

func (dbc *DBC) Ensure(condition bool, msgs ...interface{}) {
	_panic(dbc, "Ensure", condition, msgs)
}

func (dbc *DBC) Assert(condition bool, msgs ...interface{}) {
	_panic(dbc, "Assert", condition, msgs)
}

func (dbc *DBC) Check(condition bool, msgs ...interface{}) {
	_panic(dbc, "Check", condition, msgs)
}

func (dbc *DBC) SimpleValidate(validatedObject SimpleValidator, msgs ...interface{}) {
	_panic(dbc, "Validate", validatedObject.Validate(), msgs)
}

func (dbc *DBC) Validate(validatedObject Validator, msgs ...interface{}) {
	_panic(dbc, "Validate", validatedObject.Validate(),
		append(msgs, validatedObject))
}
