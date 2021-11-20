# loglines
Basic timestamped console printing.

## What
This is a very simple package with three functions:
- NowString() returns a string with the current time in a format similar to Apache's common log format. Used by the other two functions.
- Msg() prints a message to stdout.
- Err() prints an error message to stderr.

That's it. It just prints stuff neatly.

## How
Example usage:
```go
import (
	"fmt"
	ll "github.com/grimdork/loglines"
)

func main() {
	ll.Msg("Hello, console!")
	ll.Err("This is an error message.")
}
```
