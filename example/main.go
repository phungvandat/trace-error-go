package main

import (
	"fmt"

	"github.com/phungvandat/trace_err"
)

func main() {
	err := fmt.Errorf("Custom error")
	trace_err.TraceError(err)
}
