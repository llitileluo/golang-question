package errorx

import (
	"fmt"
	"runtime"
)

type Error interface {
	error
	fmt.Formatter
	Unwrap() error
	Cause() error
	Code() int
	Type() ErrType
	Stack() Stack
}

type ErrType string

const (
	ErrTypeNotFound ErrType = "not_found"
	ErrTypeTimeout  ErrType = "timeout"
	// TODO: add more error types
)

type Frame struct {
	Name string
	File string
	Line int
}

type Stack []Frame

func Wrap(err error) Error {
	if err == nil {
		return nil
	}
	return &CustomError{
		msg:   err.Error(),
		cause: err,
		stack: captureStack(),
	}
}

func New(msg string) Error {
	return &CustomError{
		msg:   msg,
		stack: captureStack(),
	}

}

func C(code int, msg string) Error {
	return &CustomError{
		msg:   msg,
		code:  code,
		stack: captureStack(),
	}
}

func Cf(code int, format string, args ...interface{}) Error {
	return &CustomError{
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		stack: captureStack(),
	}
}

func captureStack() Stack {
	pc := make([]uintptr, 10)
	//获取调用栈信息，跳过前3层（`Callers`本身及其直接调用者）
	n := runtime.Callers(3, pc)
	//获取的程序计数器信息转换成人类可读的函数调用堆栈帧信息
	frames := runtime.CallersFrames(pc[:n])

	stack := Stack{}
	for {
		frame, more := frames.Next()
		stack = append(stack, Frame{Name: frame.Function, File: frame.File, Line: frame.Line})
		if !more {
			break
		}
	}
	return stack
}
