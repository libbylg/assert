package assert

import "testing"

func TestAssert_Equal(t *testing.T) {
	expect := &Assert{T: t}
	expect.Equal("Expect the values is equal", "123", "456")
}

func TestAssert_NotEqual(t *testing.T) {
	expect := New(t)
	expect.NotEqual("Expect the values is not equal", "123", "123")
}

func TestAssert_True(t *testing.T) {
	expect := New(t)
	expect.True("Expect the expresion is true", "123" == "456")
}

func TestAssert_False(t *testing.T) {
	expect := New(t)
	expect.False("Expect the expresion is false", "123" == "123")
}

func TestAssert_Panic(t *testing.T) {
	expect := New(t)
	expect.Panic("Expect the func throw a panic", func() {
		// Do nothing.
	})
}

func throwPanic() {
	panic(123)
}

func TestAssert_NoPanic(t *testing.T) {
	expect := New(t)
	expect.NoPanic("Expect the func do not throw a panic", func() {
		throwPanic()
	})
}
