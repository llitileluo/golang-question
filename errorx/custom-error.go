package errorx

import (
	"fmt"
)

type CustomError struct {
	msg     string
	cause   error
	code    int
	errType ErrType
	stack   Stack
}

func (e *CustomError) Error() string {
	return e.msg
}

func (e *CustomError) Unwrap() error {
	return e.cause
}

func (e *CustomError) Cause() error {
	return e.cause
}

func (e *CustomError) Code() int {
	return e.code
}

func (e *CustomError) Type() ErrType {
	return e.errType
}

func (e *CustomError) Stack() Stack {
	return e.stack
}

func (e *CustomError) Format(s fmt.State, verb rune) {
	switch verb {
	//%v
	case 'v':
		//如果是%+v 就打印堆栈信息
		if s.Flag('+') {
			fmt.Fprintf(s, "Error: %s\n", e.msg)
			for _, frame := range e.stack {
				fmt.Fprintf(s, "%s:%d (in %s)\n", frame.File, frame.Line, frame.Name)
			}
			if e.cause != nil {
				fmt.Fprintf(s, "Caused by: %+v\n", e.cause)
			}
			return
		}
		fallthrough
		//%s
	case 's':
		fmt.Fprint(s, e.Error())
		//%q
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}
