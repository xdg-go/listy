package listy

// This file is for methods for containers of string-like things, for which it
// makes no sense to auto-generate for all types.

import (
	"strings"
)

// Join concatenates elements with a separator interspersed, returning a new
// native element.
func (xs S) Join(sep string) string {
	return strings.Join([]string(xs), sep)
}
