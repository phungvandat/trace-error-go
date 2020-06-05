package trace_err

import (
	"fmt"
	"log"
	"runtime"
)

func TraceError(err error) error {
	log.Printf("\nErr: %v \nTrace: %v", err.Error(), getStackTrace(5))
	return err
}

func TraceErrWithStackBuf(err error, bufNum int) error {
	log.Printf("\nErr: %v \nTrace: %v", err.Error(), getStackTrace(bufNum))
	return err
}

func getStackTrace(bufNum int) string {
	stackBuf := make([]uintptr, bufNum)
	length := runtime.Callers(3, stackBuf[:])
	stack := stackBuf[:length]

	trace := ""

	frames := runtime.CallersFrames(stack)

	for {
		frame, more := frames.Next()
		trace = trace + fmt.Sprintf("\n\tFile: %s, Function: %s, Line: %d. ", frame.File, frame.Function, frame.Line)
		if !more {
			break
		}
	}

	return trace
}
