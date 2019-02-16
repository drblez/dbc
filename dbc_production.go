//+build without_dbc

package dbc

type DBC struct {
}

func New() *DBC {
	return &DBC{}
}

func NewWithLogger(log Logger) *DBC {
	return &DBC{}
}

func (dbc *DBC) Require(condition bool, msgs ...interface{}) {
}

func (dbc *DBC) Ensure(condition bool, msgs ...interface{}) {
}

func (dbc *DBC) Assert(condition bool, msgs ...interface{}) {
}

func (dbc *DBC) Check(condition bool, msgs ...interface{}) {
}

func (dbc *DBC) SimpleInvariant(validatedObject SimpleInvariantValidator, msgs ...interface{}) {
}

func (dbc *DBC) Invariant(validatedObject InvariantValidator, msgs ...interface{}) {
}

func Require(condition bool, msgs ...interface{}) {
}

func Ensure(condition bool, msgs ...interface{}) {
}

func Assert(condition bool, msgs ...interface{}) {
}

func Check(condition bool, msgs ...interface{}) {
}

func SimpleInvariant(validatedObject SimpleInvariantValidator, msgs ...interface{}) {
}

func Invariant(validatedObject InvariantValidator, msgs ...interface{}) {
}
