//+build !without_dbc

package dbc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBC_Assert(t *testing.T) {
	contract := New()
	assert.Panics(t, func() {
		contract.Assert(false)
	})
	assert.NotPanics(t, func() {
		contract.Assert(true)
	})
}

func TestDBC_Check(t *testing.T) {
	contract := New()
	assert.Panics(t, func() {
		contract.Check(false)
	})
	assert.NotPanics(t, func() {
		contract.Check(true)
	})
}

func TestDBC_Ensure(t *testing.T) {
	contract := New()
	assert.Panics(t, func() {
		contract.Ensure(false)
	})
	assert.NotPanics(t, func() {
		contract.Ensure(true)
	})
}

func TestDBC_Require(t *testing.T) {
	contract := New()
	assert.Panics(t, func() {
		contract.Require(false)
	})
	assert.NotPanics(t, func() {
		contract.Require(true)
	})
}

type T1 struct {
	f bool
}

func (t1 T1) Validate() bool {
	return t1.f
}

func TestDBC_SimpleValidate(t *testing.T) {
	contract := New()
	assert.Panics(t, func() {
		t1 := T1{
			f: false,
		}
		contract.SimpleValidate(t1)
	})
	assert.NotPanics(t, func() {
		t1 := T1{
			f: true,
		}
		contract.SimpleValidate(t1)
	})
}

type T2 struct {
	f bool
}

func (t2 T2) Validate() bool {
	return t2.f
}

func (t2 T2) String() string {
	return fmt.Sprintf("T2: %t", t2.f)
}

func TestDBC_Validate(t *testing.T) {
	contract := New()
	assert.Panics(t, func() {
		t2 := T2{
			f: false,
		}
		contract.Validate(t2)
	})
	assert.NotPanics(t, func() {
		t2 := T2{
			f: true,
		}
		contract.Validate(t2)
	})
}
