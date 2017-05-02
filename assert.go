package assert

import (
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

// Assert is the wrapper of testing.T
type Assert struct {
	T *testing.T
}

// New is used to create a new Assert object.
func New(t *testing.T) *Assert {
	a := Assert{T: t}
	return &a
}

// Equal is used to check if exp equals to got.
func (a *Assert) Equal(message string, exp, got interface{}) {
	result := reflect.DeepEqual(exp, got)
	if !result {
		_, file, line, _ := runtime.Caller(1)
		a.T.Errorf("\n%s:%d\n%s:\n\t- Expect equal:%s\n\t- Expect:%v\n\t- ButGot:%v\n", file, line, a.T.Name(), message, exp, got)
		a.T.FailNow()
	}
}

// NotEqual is used to check if exp is not equals to got
func (a *Assert) NotEqual(message string, exp, got interface{}) {
	result := reflect.DeepEqual(exp, got)
	if result {
		_, file, line, _ := runtime.Caller(1)
		a.T.Errorf("\n%s:%d\n%s:\n\t- Expect not equal:%s\n\t- Expect:%v\n\t- ButGot:%v\n", file, line, a.T.Name(), message, exp, got)
		a.T.FailNow()
	}
}

//
//// Match is used to check the got is match to the regular expression of exp.
//func (a *Assert) Match(message string, exp string, got string) {
//	regex, err :=  regexp.Compile(exp)
//	if nil != err {
//		a.T.FailNow()
//	}
//
//	if !regex.MatchString(got) {
//		_, file, line, _ := runtime.Caller(1)
//		a.T.Errorf("\n%s:%d\n%s:\n\t- Expect match:%s\n\t- Expect:%v\n\t- ButGot:%v\n", file, line, a.T.Name(), message, exp, got)
//		a.T.FailNow()
//	}
//}

// True is used to check the got be true.
func (a *Assert) True(message string, got bool) {
	result := reflect.DeepEqual(true, got)
	if !result {
		_, file, line, _ := runtime.Caller(1)
		a.T.Errorf("\n%s:%d\n%s:\n\t- Expect true:%s\n\t- Expect:%v\n\t- ButGot:%v\n", file, line, a.T.Name(), message, false, got)
		a.T.FailNow()
	}
}

// False is used to check the got be false.
func (a *Assert) False(message string, got bool) {
	result := reflect.DeepEqual(false, got)
	if !result {
		_, file, line, _ := runtime.Caller(1)
		a.T.Errorf("\n%s:%d\n%s:\n\t- Expect false:%s\n\t- Expect:%v\n\t- ButGot:%v\n", file, line, a.T.Name(), message, false, got)
		a.T.FailNow()
	}
}

// Panic is used to check the fn should give a panic.
func (a *Assert) Panic(message string, fn func()) {
	defer func() {
		recover()
	}()

	fn()

	_, file, line, _ := runtime.Caller(1)
	a.T.Errorf("\n%s:%d\n%s:\n\t- Expect panic:%s\n\t- Expect:Panic\n\t- ButGot:No Panic\n", file, line, a.T.Name(), message)
	a.T.FailNow()
}

// NoPanic is used to check the fn should not give a panic.
func (a *Assert) NoPanic(message string, fn func()) {
	defer func() {
		r := recover()
		_, file, line, _ := runtime.Caller(3)
		a.T.Errorf("\n%s:%d\n%s:\n\t- Expect no panic:%s\n\t- Expect:No Panic\n\t- ButGot:Panic(%v)\n", file, line, a.T.Name(), message, r)
		a.T.FailNow()
	}()

	fn()
}
