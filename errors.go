package errors

import (
	"errors"
	"fmt"
	"runtime"
)

// ErrUnsupported is wrapper of errors.ErrUnsupported
var ErrUnsupported = errors.ErrUnsupported

// New is wrapper of errors.New
func New(text string) error {
	return errors.New(text)
}

// Join is wrapper of errors.Join
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// Unwrap is wrapper of errors.Unwrap
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Is is wrapper of errors.Is
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// IsAny is super set of errors.Is
// It returns true if err is match any of errs.
func IsAny(err error, errs ...error) bool {
	for _, e := range errs {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}

// As is wrapper of errors.As
func As(err error, target any) bool {
	return errors.As(err, target)
}

// AsAny is super set of errors.As
// It returns true if err is match any of targets.
func AsAny(err error, targets ...any) bool {
	for _, t := range targets {
		if errors.As(err, t) {
			return true
		}
	}
	return false
}

// PrintStack prints stack trace to stdout without carriage return.
func PrintStack() {
	fmt.Print(stack())
}

// PrintlnStack prints stack trace to stdout with carriage return.
func PrintlnStack() {
	fmt.Println(stack())
}

// Stack returns stack trace as error.
func Stack() error {
	return stack()
}

func stack() error {
	numg := runtime.NumGoroutine()
	i := 0
	var buf []byte
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		buf = fmt.Appendf(buf, "%s()\n\t%s:%d\n", fn.Name(), file, line)
		i++
	}
	return fmt.Errorf("=== Stack trace ===\ngoroutine %d [running]:\n%s", numg, buf)
}
