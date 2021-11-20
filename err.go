package logline

import (
	"fmt"
	"os"
	"strings"
)

// Err prints formatted messages to stderr, starting with a nicely formatted timestamp.
func Err(f string, v ...interface{}) {
	var b strings.Builder
	b.WriteString(NowString())
	b.WriteRune(':')
	b.WriteString(fmt.Sprintf(f, v...))
	b.WriteString("\n")
	fmt.Fprintf(os.Stderr, b.String())
}
