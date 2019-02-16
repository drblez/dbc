//+build !without_dbc

package dbc

import (
	"errors"
	"fmt"
	"runtime"
)

// Contract structure to hold info about Logger
type Contract struct {
	log Logger
}

// Create new contract without logging
func New() *Contract {
	return &Contract{
		log: nil,
	}
}

// Create new contract with logging
func NewWithLogger(log Logger) *Contract {
	return &Contract{
		log: log,
	}
}

func _panic(contract *Contract, function string, condition bool, msgs ...interface{}) {
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
	if contract.log != nil {
		contract.log.Debug(str)
	}
	for i := range msgs {
		s := fmt.Sprintf("[%d] %+v", i, msgs[i])
		str += "\n" + s
		if contract.log != nil {
			contract.log.Debug(s)
		}
	}
	panic(errors.New(str))
}

// Requires check preconditions
//
//    ...
//    contract := dbc.New()
//    ...
//
//    func f(parm1 int, parm2 int) {
//        contract.Require(parm1 != parm2)
//        ...
//    }
//
func (contract *Contract) Require(condition bool, msgs ...interface{}) {
	_panic(contract, "Require", condition, msgs)
}

// Requires check postconditions
//
//    ...
//    contract := dbc.New()
//    ...
//
//    func f() string {
//       ..
//       value, ok := someMap[key]
//       contract.Ensure(ok)
//       return value
//    }
//
func (contract *Contract) Ensure(condition bool, msgs ...interface{}) {
	_panic(contract, "Ensure", condition, msgs)
}

// Simple assert
func (contract *Contract) Assert(condition bool, msgs ...interface{}) {
	_panic(contract, "Assert", condition, msgs)
}

// Simple assert, but "Check" name
func (contract *Contract) Check(condition bool, msgs ...interface{}) {
	_panic(contract, "Check", condition, msgs)
}

// SimpleInvariant use SimpleInvariantValidator interface for check conditions.
/*
   ...
   type T1 struct {
       f bool
   }
   func (t1 T1) Invariant() bool {
       return t1.f
   }
   ...
   contract := New()
   ...
   t1 := T1{
       f: false,
   }
   ...
   contract.SimpleInvariant(t1)
*/
func (contract *Contract) SimpleInvariant(validatedObject SimpleInvariantValidator, msgs ...interface{}) {
	_panic(contract, "Invariant", validatedObject.Invariant(), msgs)
}

// Invariant use Validator interface (with Stringer) for check conditions.
/*
   ...
   type T1 struct {
       f bool
   }
   func (t1 T1) Invariant() bool {
       return t1.f
   }
   func (t1 T1) String() string {
       return fmt.Sprintf("T1.f: %t", t1.f)
   }
   ...
   contract := New()
   ...
   t1 := T1{
       f: false,
   }
   ...
   contract.Invariant(t1)

*/
func (contract *Contract) Invariant(validatedObject InvariantValidator, msgs ...interface{}) {
	_panic(contract, "Invariant", validatedObject.Invariant(),
		append(msgs, validatedObject))
}

var defaultContract = New()

func Require(condition bool, msgs ...interface{}) {
	_panic(defaultContract, "Require", condition, msgs)
}

func Ensure(condition bool, msgs ...interface{}) {
	_panic(defaultContract, "Ensure", condition, msgs)
}

func Assert(condition bool, msgs ...interface{}) {
	_panic(defaultContract, "Assert", condition, msgs)
}

func Check(condition bool, msgs ...interface{}) {
	_panic(defaultContract, "Check", condition, msgs)
}

func SimpleInvariant(validatedObject SimpleInvariantValidator, msgs ...interface{}) {
	_panic(defaultContract, "Invariant", validatedObject.Invariant(), msgs)
}

func Invariant(validatedObject InvariantValidator, msgs ...interface{}) {
	_panic(defaultContract, "Invariant", validatedObject.Invariant(),
		append(msgs, validatedObject))
}
