package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

var ErrFieldNotFound = fmt.Errorf("field not found")

type Model interface {
	// Fields returns a set of the fields in the Model.
	Fields() map[string]bool

	// UnsafeValueAt returns the value of the field in the Model.
	// It may panic if the field does not exist.
	UnsafeValueAt(field string) any

	// ValueAt returns the value of the field in the Model.
	// It returns ErrFieldNotFound if the field does not exist.
	ValueAt(field string) (any, error)
}

// StringOf returns a string representation of the Model.
func StringOf(t Model) (string, error) {
	var out strings.Builder
	for field := range t.Fields() {
		value := t.UnsafeValueAt(field)
		buf, err := json.Marshal(value)
		if err != nil {
			return "", fmt.Errorf("failed to marshal value: %w", err)
		}
		_, _ = fmt.Fprintf(&out, "%20s%s\n", field, buf)
	}
	return out.String(), nil
}

// FieldSlice returns a slice of the fields in the Model.
func FieldSlice(t Model) []string {
	fieldSet := t.Fields()
	fieldSlice := make([]string, 0, len(fieldSet))
	for field := range fieldSet {
		fieldSlice = append(fieldSlice, field)
	}
	return fieldSlice
}
