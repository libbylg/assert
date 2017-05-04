package assert

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"runtime"
)

type Assertor interface {
	Assert() error
}

type Equal struct {
	Expect interface{}
	Actual interface{}
}

func (r *Equal) Assert() error {
	if !reflect.DeepEqual(r.Expect, r.Actual) {
		return fmt.Errorf("Expect:%v, Actual:%v", r.Expect, r.Actual)
	}
	
	return nil
}

type NotEqual struct {
	Expect interface{}
	Actual interface{}
}

func (r *NotEqual) Assert() error {
	if reflect.DeepEqual(r.Expect, r.Actual) {
		return fmt.Errorf("Expect:%v, Actual:%v", r.Expect, r.Actual)
	}
	
	return nil
}

type True struct {
	Actual bool
}

func (r *True) Assert() error {
	if true != r.Actual {
		return fmt.Errorf("Expect:%v, Actual:%v", true, r.Actual)
	}
	
	return nil
}

type False struct {
	Actual bool
}

func (r *False) Assert() error {
	if false != r.Actual {
		return fmt.Errorf("Expect:%v, Actual:%v", false, r.Actual)
	}
	
	return nil
}

type Panic struct {
	F func()
}

func (r *Panic) Assert() (err error) {
	// 先对 err 赋值,占据一个位置
	err = fmt.Errorf("")
	
	// 如果fn抛出panic,那么逻辑会进入这里
	defer func() {
		recover()
		if nil == err {
			err = fmt.Errorf("Expect panic, but no panic catched")
		}
	}()
	
	r.F()
	
	// 如果程序的逻辑走到这里说明没有碰到任何panic
	err = nil
	return
}

type NoPanic struct {
	F func()
}

func (r *NoPanic) Assert() (err error) {
	// 先对 err 赋值,占据一个位置
	err = fmt.Errorf("")
	
	// 如果fn抛出panic,那么逻辑会进入这里
	defer func() {
		ret := recover()
		if nil != err {
			err = fmt.Errorf("Expect no panic, but panic catched:%v", ret)
		}
	}()
	
	r.F()
	
	err = nil
	return
}

type Match struct {
	Regexp string
	Actual string
}

func (r *Match) Assert() error {
	regex, _ := regexp.Compile(r.Regexp)
	if !regex.MatchString(r.Actual) {
		return fmt.Errorf("Expect match:`%s`, but actual `%s`", r.Regexp, r.Actual)
	}
	
	return nil
}

type NotMatch struct {
	Regexp string
	Actual string
}

func (r *NotMatch) Assert() error {
	regex, _ := regexp.Compile(r.Regexp)
	if regex.MatchString(r.Actual) {
		return fmt.Errorf("Expect not match:`%s`, but actual `%s`", r.Regexp, r.Actual)
	}
	
	return nil
}

type Nil struct {
	Actual interface{}
}

func (r *Nil) Assert() error {
	if nil != r.Actual {
		return fmt.Errorf("Expect nil, but actual not nil:%v", r.Actual)
	}
	
	return nil
}

type NotNil struct {
	Actual interface{}
}

func (r *NotNil) Assert() error {
	if nil != r.Actual {
		return fmt.Errorf("Expect not nil, but actual nil")
	}
	
	return nil
}

// Assert is the wrapper of testing.T
type Assert struct {
	T *testing.T
	F bool //  true: Fail(); false: FailNow()
}

// New is used to create a new Assert object.
func New(t *testing.T) (*Assert, *Assert) {
	return &Assert{T: t, F: false}, &Assert{T: t, F: true}
}

func (a *Assert) Assert(message string, assertor Assertor) {
	a.AssertInner(message, assertor, 2)
}

func (a *Assert) AssertInner(message string, assertor Assertor, callerSkip int) {
	if err := assertor.Assert(); nil != err {
		_, file, line, _ := runtime.Caller(callerSkip)
		a.T.Errorf("\n%s:%d\n%s\n%s\n", file, line, message, err.Error())
		if a.F {
			a.T.Fail()
		} else {
			a.T.FailNow()
		}
	}
}

// PassValue is used to check if exp equals to got.
func (a *Assert) Equal(message string, exp, got interface{}) {
	a.AssertInner(message, &Equal{Expect: exp, Actual: got}, 2)
}

// NotEqual is used to check if exp is not equals to got
func (a *Assert) NotEqual(message string, exp, got interface{}) {
	a.AssertInner(message, &NotEqual{Expect: exp, Actual: got}, 2)
}

// True is used to check the got be true.
func (a *Assert) True(message string, got bool) {
	a.AssertInner(message, &True{Actual: got}, 2)
}

// False is used to check the got be false.
func (a *Assert) False(message string, got bool) {
	a.AssertInner(message, &False{Actual: got}, 2)
}

// Panic is used to check the fn should give a panic.
func (a *Assert) Panic(message string, fn func()) {
	a.AssertInner(message, &Panic{fn}, 2)
}

// NoPanic is used to check the fn should not give a panic.
func (a *Assert) NoPanic(message string, fn func()) {
	a.AssertInner(message, &NoPanic{fn}, 2)
}

// Match is used to check the got is match to the regular expression of exp.
func (a *Assert) Match(message string, regex string, got string) {
	a.AssertInner(message, &Match{Regexp: regex, Actual: got}, 2)
}

func (a *Assert) NotMatch(message string, regex string, got string) {
	a.AssertInner(message, &NotMatch{Regexp: regex, Actual: got}, 2)
}

func (a *Assert) Nil(message string, got interface{}) {
	a.AssertInner(message, &Nil{Actual: got}, 2)
}

func (a *Assert) NotNil(message string, got interface{}) {
	a.AssertInner(message, &NotNil{Actual: got}, 2)
}

func (a *Assert) In(message string, got interface{}) {
}

func (a *Assert) OneOf() {

}

func (a *Assert) Empty(message string, got interface{}) {

}

func (a *Assert) NotEmpty() {

}
