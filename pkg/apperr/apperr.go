package apperr

import (
	"fmt"
	"runtime"
)

type AppErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func New(message string) *AppErr {
	return &AppErr{
		Message: message,
	}
}

func (e *AppErr) Error() string {
	return e.Message
}

func (e *AppErr) WithStatus(status int) *AppErr {
	e.Status = status
	return e
}

type InternalError struct {
	err   error
	stack []runtime.Frame
}

func NewInternalError(err error) *InternalError {
	pcs := make([]uintptr, 10)
	n := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs[:n])
	var stack []runtime.Frame
	for {
		frame, more := frames.Next()
		stack = append(stack, frame)
		if !more {
			break
		}
	}
	return &InternalError{
		err:   err,
		stack: stack,
	}
}

func (e *InternalError) Error() string {
	out := fmt.Sprintf("Error: %v\nStack Trace:\n", e.err)
	for _, frame := range e.stack {
		out += fmt.Sprintf("  %s:%d %s\n", frame.File, frame.Line, frame.Function)
	}
	return out
}

func (e *InternalError) StackTrace(formats ...func(file, function string, line int) string) string {
	var out string
	format := func(file, function string, line int) string {
		return fmt.Sprintf("%s:%d %s", file, line, function)
	}
	if len(formats) > 0 {
		format = formats[0]
	}

	for _, frame := range e.stack {
		out += format(frame.File, frame.Function, frame.Line) + "\n"
	}

	return out
}

func (e *InternalError) Unwarp() error {
	return e.err
}
