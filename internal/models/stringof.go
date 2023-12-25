package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Fielder interface {
	Fields() []string
}

// String returns a string representation of the Organization.
func StringOf(t Fielder) (string, error) {
	fields := t.Fields()

	j, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	tAsMap := make(map[string]any, len(fields))
	if err := json.Unmarshal(j, &tAsMap); err != nil {
		return "", err
	}

	var out strings.Builder
	for _, field := range fields {
		_, _ = fmt.Fprintf(&out, "%20s%v\n", field, tAsMap[field])
	}

	return out.String(), nil
}
