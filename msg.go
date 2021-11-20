package loglines

import (
	"fmt"
	"os"
	"strings"
)

// Msg prints formatted messages to stdout, starting with a nicely formatted timestamp.
func Msg(f string, v ...interface{}) {
	var b strings.Builder
	b.WriteString(NowString())
	b.WriteRune(':')
	b.WriteString(fmt.Sprintf(f, v...))
	fmt.Fprintln(os.Stdout, b.String())
}
