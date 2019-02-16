//+build without_dbc

package dbc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContract_Assert(t *testing.T) {
	contract := New()
	assert.NotPanics(t, func() {
		contract.Assert(false)
	})
	assert.NotPanics(t, func() {
		contract.Assert(true)
	})
}

func TestContract_Check(t *testing.T) {
	contract := New()
	assert.NotPanics(t, func() {
		contract.Check(false)
	})
	assert.NotPanics(t, func() {
		contract.Check(true)
	})
}

func TestContract_Ensure(t *testing.T) {
	contract := New()
	assert.NotPanics(t, func() {
		contract.Ensure(false)
	})
	assert.NotPanics(t, func() {
		contract.Ensure(true)
	})
}

func TestContract_Require(t *testing.T) {
	contract := New()
	assert.NotPanics(t, func() {
		contract.Require(false)
	})
	assert.NotPanics(t, func() {
		contract.Require(true)
	})
}

type T1 struct {
	f bool
}

func (t1 T1) Invariant() bool {
	return t1.f
}

func TestContract_SimpleInvariant(t *testing.T) {
	contract := New()
	assert.NotPanics(t, func() {
		t1 := T1{
			f: false,
		}
		contract.SimpleInvariant(t1)
	})
	assert.NotPanics(t, func() {
		t1 := T1{
			f: true,
		}
		contract.SimpleInvariant(t1)
	})
}

type T2 struct {
	f bool
}

func (t2 T2) Invariant() bool {
	return t2.f
}

func (t2 T2) String() string {
	return fmt.Sprintf("T2: %t", t2.f)
}

func TestContract_Invariant(t *testing.T) {
	contract := New()
	assert.NotPanics(t, func() {
		t2 := T2{
			f: false,
		}
		contract.Invariant(t2)
	})
	assert.NotPanics(t, func() {
		t2 := T2{
			f: true,
		}
		contract.Invariant(t2)
	})
}

func TestAssert(t *testing.T) {
	assert.NotPanics(t, func() {
		Assert(false)
	})
	assert.NotPanics(t, func() {
		Assert(true)
	})
}

func TestCheck(t *testing.T) {
	assert.NotPanics(t, func() {
		Check(false)
	})
	assert.NotPanics(t, func() {
		Check(true)
	})
}

func TestEnsure(t *testing.T) {
	assert.NotPanics(t, func() {
		Ensure(false)
	})
	assert.NotPanics(t, func() {
		Ensure(true)
	})
}

func TestRequire(t *testing.T) {
	assert.NotPanics(t, func() {
		Require(false)
	})
	assert.NotPanics(t, func() {
		Require(true)
	})
}

func TestSimpleInvariant(t *testing.T) {
	assert.NotPanics(t, func() {
		t1 := T1{
			f: false,
		}
		SimpleInvariant(t1)
	})
	assert.NotPanics(t, func() {
		t1 := T1{
			f: true,
		}
		SimpleInvariant(t1)
	})
}

func TestInvariant(t *testing.T) {
	assert.NotPanics(t, func() {
		t2 := T2{
			f: false,
		}
		Invariant(t2)
	})
	assert.NotPanics(t, func() {
		t2 := T2{
			f: true,
		}
		Invariant(t2)
	})
}
