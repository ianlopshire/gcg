package gcg

import (
	"strings"
)

// Comment represents a comment or comment block.
// See https://golang.org/ref/spec#Comments.
type Comment string

// NewComment creates a new comment based on the provided string.
func NewComment(s string) Comment {
	return Comment(s)
}

// MarshalGo generates go code for the comment.
// If the comment is multiple lines each line will be prefixed with `//`
func (c *Comment) MarshalGo() ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}
	lines := strings.Split(string(*c), "\n")
	return []byte(`// ` + strings.Join(lines, "\n"+`// `)), nil
}
