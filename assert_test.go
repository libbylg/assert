package assert

import "testing"

func TestAssert_Equal(t *testing.T) {
	expect := &Assert{T: t}
	expect.Equal("检查是否相等", "123", "456")
}

func TestAssert_NotEqual(t *testing.T) {
	expect := New(t)
	expect.NotEqual("检查是否相等", "123", "123")
}

func TestAssert_True(t *testing.T) {
	expect := New(t)
	expect.True("检查是否相等", "123" == "456")
}

func TestAssert_False(t *testing.T) {
	expect := New(t)
	expect.False("检查是否相等", "123" == "123")
}

func TestAssert_Panic(t *testing.T) {
	expect := New(t)
	expect.Panic("期望抛出异常", func() {
		// Do nothing.
	})
}

func throwPanic() {
	panic(123)
}

func TestAssert_NoPanic(t *testing.T) {
	expect := New(t)
	expect.NoPanic("期望不抛出异常", func() {
		throwPanic()
	})
}
