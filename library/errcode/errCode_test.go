package errcode

import (
	"errors"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	testErr := &Errno{
		Code:    0,
		Message: "test",
	}
	err := errors.New("test error")
	got := New(testErr, err)
	want := &Err{
		Code:    0,
		Message: "test",
		ErrCord: err,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("new errCode fail")
	}
}

func TestDecodeErr(t *testing.T) {
	err1 := &Err{
		Code:    0,
		Message: "test1",
		ErrCord: nil,
	}
	code1, msg1 := DecodeErr(err1)
	if code1 != 0 || msg1 != "test1" {
		t.Errorf("err decode Err")
	}
	err2 := &Errno{
		Code:    0,
		Message: "test2",
	}
	code2, msg2 := DecodeErr(err2)
	if code2 != 0 || msg2 != "test2" {
		t.Errorf("err decode Errno")
	}
	err3 := errors.New("test err")
	code3, msg3 := DecodeErr(err3)
	if code3 != SystemError.Code || msg3 != err3.Error() {
		t.Errorf("err decode err")
	}
}

func TestErr_Add(t *testing.T) {
	err := &Err{
		Code:    0,
		ErrCord: nil,
	}
	err = err.Add("test err")
	if err.Message != " test err" {
		t.Errorf("err addHttp-transport-grpc-enpoint Err message")
	}
}

func TestErr_AddF(t *testing.T) {
	err := &Err{
		Code:    0,
		Message: "test",
		ErrCord: nil,
	}
	err = err.AddF("%s %s", "test1", "test2")
	if err.Message != "test test1 test2" {
		t.Errorf("err addHttp-transport-grpc-enpoint Err message")
	}
}
