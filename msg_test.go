package loglines_test

import (
	"testing"

	"github.com/grimdork/loglines"
)

func TestMsg(t *testing.T) {
	loglines.Msg("This goes to stdout.")
	loglines.Err("This goes to stderr.")
}
