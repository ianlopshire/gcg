package gcg

import "fmt"
import "github.com/pkg/errors"

// NamedType represents a type declaration.
// See https://golang.org/ref/spec#Types.
type NamedType struct {
	Name           string
	UnderlyingType string
}

// NewNamedType creates an new Named Type
func NewNamedType(name, underlyingType string) NamedType {
	return NamedType{
		Name:           name,
		UnderlyingType: underlyingType,
	}
}

// MarshalGo generates go code for the named type.
func (t *NamedType) MarshalGo() ([]byte, error) {
	if t == nil {
		return nil, errors.New("gcg: cannot marshal nil NamedType")
	}
	if t.Name == "" {
		return nil, errors.New("gcg: cannot marshal NamedType with empty name")
	}
	if t.UnderlyingType == "" {
		return nil, errors.New("gcg: cannot marshal NamedType with empty underlying type")
	}
	return []byte(fmt.Sprintf("type %s %s\n", t.Name, t.UnderlyingType)), nil
}
