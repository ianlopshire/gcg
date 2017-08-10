package gcg

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// Struct represents a new struct type deleration.
// See https://golang.org/ref/spec#Struct_types.
type Struct struct {
	Name   string
	Fields []StructField
}

// NewStruct creates a new struct with the provided fields.
func NewStruct(name string, fields ...StructField) Struct {
	return Struct{
		Name:   name,
		Fields: fields,
	}
}

// AddField adds a field to a struct.
func (s *Struct) AddField(field StructField) {
	s.Fields = append(s.Fields, field)
}

// MarshalGo generates go code for the struct.
func (s *Struct) MarshalGo() ([]byte, error) {
	if s == nil {
		return nil, errors.New("gcg: nil struct cannot be marshaled")
	}
	buff := bytes.NewBuffer(nil)
	buff.WriteString(fmt.Sprintf("type %s struct {", s.Name))
	for _, field := range s.Fields {
		s, err := field.MarshalGo()
		if err != nil {
			return nil, err
		}
		buff.Write(s)
		buff.WriteString(";")
	}
	buff.WriteString("}")
	return buff.Bytes(), nil
}

// StructField represents a struct field.
// See https://golang.org/ref/spec#Struct_types.
type StructField struct {
	Name string
	Type string
	Tags StructTags
}

// NewStructField creates a new struct field.
func NewStructField(name, typeName string, tags StructTags) StructField {
	return StructField{
		Name: name,
		Type: typeName,
		Tags: tags,
	}
}

// MarshalGo generates go code for the struct field.
// When Name is empty the struct field will be marshaled as an ebedded type.
func (f *StructField) MarshalGo() ([]byte, error) {
	if f == nil {
		return nil, errors.New("gcg: nil struct field cannot be marshaled")
	}
	if f.Type == "" {
		return nil, errors.New("gcg: struct field cannot be marshaled with empty type")
	}
	ts, err := f.Tags.MarshalGo()
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%s %s %s", f.Name, f.Type, ts)), nil
}

// StructTags represent a set of struct tags.
// See https://golang.org/ref/spec#Struct_types.
type StructTags map[string]string

// NewStructTags creates a new set of struct tags.
// FieldValues will be used in sets as key: value.
// If an odd number of fieldValues are provided the last value will be the string zero value ("").
func NewStructTags(fieldValues ...string) StructTags {
	st := StructTags{}

	for i := 0; i < len(fieldValues); i += 2 {
		key := fieldValues[i]
		var value string
		if (i + 1) < len(fieldValues) {
			value = fieldValues[i+1]
		}
		st[key] = value
	}

	return st
}

// MarshalGo generates go code for the struct tags.
// Nil or empty StructTags will be ignored.
func (t *StructTags) MarshalGo() ([]byte, error) {
	if t == nil || len(*t) <= 0 {
		return nil, nil
	}
	keyvals := []string{}
	for key, value := range *t {
		keyvals = append(keyvals, fmt.Sprintf(`%s:"%s"`, key, value))
	}
	return []byte("`" + strings.Join(keyvals, " ") + "`"), nil
}
