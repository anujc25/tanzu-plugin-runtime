package component

import (
	"io"
	"strings"
)

// terminalEscaper replaces ANSI escape sequences and other terminal special
// characters to avoid terminal escape character attacks (issue #101695).
// Add "\x1b", "^[" to the `NewReplacer` params to scape color
var terminalEscaper = strings.NewReplacer("\r", "\\r")

// WriteEscaped replaces unsafe terminal characters with replacement strings
// and writes them to the given writer.
func WriteEscaped(writer io.Writer, output string) error {
	_, err := terminalEscaper.WriteString(writer, output)
	return err
}

// EscapeTerminal escapes terminal special characters in a human readable (but
// non-reversible) format.
func EscapeTerminal(in string) string {
	return terminalEscaper.Replace(in)
}
