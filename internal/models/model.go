package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/google/uuid"
)

var ErrFieldNotFound = fmt.Errorf("field not found")

type Model interface {
	// Fields returns a the fields in the Model mapped to an index.
	Fields() *orderedmap.OrderedMap[string, int]

	// ValueAtIdx returns the value of the field in the Model at i.
	ValueAtIdx(i int) any

	// ValueAt returns the value of the field in the Model.
	// It returns ErrFieldNotFound if the field does not exist.
	ValueAt(field string) (any, error)
}

// StringOf returns a string representation of the Model.
func StringOf(t Model) (string, error) {
	var out strings.Builder
	m := t.Fields()
	for el := m.Front(); el != nil; el = el.Next() {
		value := t.ValueAtIdx(el.Value)
		buf, err := json.Marshal(value)
		if err != nil {
			return "", fmt.Errorf("failed to marshal value %v: %w", value, err)
		}
		_, _ = fmt.Fprintf(&out, "%-20s\t%s\n", el.Key, buf)
	}
	return out.String(), nil
}

// FieldSlice returns a slice of the fields in the Model.
func FieldSlice(t Model) []string {
	fieldSet := t.Fields()
	fieldSlice := make([]string, 0, fieldSet.Len())
	fieldSlice = append(fieldSlice, fieldSet.Keys()...)
	return fieldSlice
}

func ValueContains(val any, query string) bool {
	switch value := val.(type) {
	case string:
		return value == query

	case []string:
		return slices.Contains(value, query)

	case bool:
		return strconv.FormatBool(value) == query

	case int:
		parsed, err := strconv.Atoi(query)
		if err != nil {
			return false
		}
		return parsed == value

	case uuid.UUID:
		parsed, err := uuid.Parse(query)
		if err != nil {
			return false
		}
		return parsed == value

	// other types don't appear in the data, extend this when they do
	default:
		return false
	}
}
