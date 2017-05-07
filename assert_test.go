package assert

import (
	"testing"
)

func TestAssert_Equal(t *testing.T) {
	expect := &Assert{T: t}
	expect.Equal("Expect-the-values-is-equal", "123", "456")
}

func TestAssert_NotEqual(t *testing.T) {
	_, expect := New(t)
	expect.NotEqual("Expect-the-values-is-not-equal", "123", "123")
}

func TestAssert_True(t *testing.T) {
	_, expect := New(t)
	expect.True("Expect-the-expresion-is-true", "123" == "456")
}

func TestAssert_False(t *testing.T) {
	_, expect := New(t)
	expect.False("Expect-the-expresion-is-false", "123" == "123")
}

func TestAssert_Panic(t *testing.T) {
	_, expect := New(t)
	expect.Panic("Expect-the-func-throw-a-panic", func() { /* Do nothing. */ })
}

func throwPanic() {
	panic(123)
}

func TestAssert_NoPanic(t *testing.T) {
	_, expect := New(t)
	expect.NoPanic("Expect-the-func-do-not-throw-a-panic", func() {
		throwPanic()
	})
}

func TestAssert_NoPanic_Use_Orign(t *testing.T) {
	_, expect := New(t)
	expect.Assert("Expect-the func-do-not-throw-a-panic", &NoPanic{func() { throwPanic() }})
}

func TestAssert_NoEqual_Use_Orign1(t *testing.T) {
	assert, _ := New(t)
	assert.Assert("Expect-not-equal", &NotEqual{"444", "444"})
}

func TestAssert_NoEqual_Use_Orign2(t *testing.T) {
	assert, _ := New(t)
	assert.Assert(&NotEqual{"444", "444"}, "Expect-not-equal")
}

func TestAssert_NoEqual_Use_Orign3(t *testing.T) {
	assert, _ := New(t)
	assert.Assert(&NotEqual{444, 444, "Expect-not-equal"})
}

func TestAssert_NoEqual_Use_Orign4(t *testing.T) {
	assert, _ := New(t)
	assert(&NotEqual{444, 444, "Expect-not-equal"})
}
