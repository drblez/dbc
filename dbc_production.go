//+build without_dbc

package dbc

type DBC struct {
}

func New() *DBC {
}

func NewWithLogger(log Logger) *DBC {
}

func (dbc *DBC) Require(condition bool, msgs ...interface{}) {
}

func (dbc *DBC) Ensure(condition bool, msgs ...interface{}) {
}

func (dbc *DBC) Assert(condition bool, msgs ...interface{}) {
}

func (dbc *DBC) Check(condition bool, msgs ...interface{}) {
}

func (dbc *DBC) SimpleValidate(validatedObject SimpleValidator, msgs ...interface{}) {
}

func (dbc *DBC) Validate(validatedObject Validator, msgs ...interface{}) {
}
