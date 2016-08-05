package checks

import (
	"fmt"
	"os"
)

type ErrReporter interface {
	ErrorStr(s string)
	FatalErrorStr(s string)
	ReportedErr() bool
}

// StdErrReporter implements ErrReporter by sending errors to StdErr
type StdErrReporter struct {
	reportedErr bool
}

func (r *StdErrReporter) ErrorStr(s string) {
	r.errImpl(s, false)
}

func (r *StdErrReporter) FatalErrorStr(s string) {
	r.errImpl(s, false)
}

func (r *StdErrReporter) errImpl(s string, shouldExit bool) {
	r.reportedErr = true
	fmt.Fprintf(os.Stderr, s)
	if shouldExit {
		os.Exit(1)
	}
}

func (r *StdErrReporter) ReportedErr() bool {
	return r.reportedErr
}
