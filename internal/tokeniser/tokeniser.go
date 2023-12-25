package tokeniser

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
)

//nolint:revive
func Tokenise(m models.Model) []string {
	tokens := []string{}
	fields := m.Fields()
	for el := fields.Front(); el != nil; el = el.Next() {
		// When the data has other types, this needs to be extended
		switch value := m.ValueAtIdx(el.Value).(type) {
		case string:
			for _, s := range strings.Split(value, " ") {
				if s == "" {
					continue
				}
				tokens = append(tokens, s)
			}
		case []string:
			for _, t := range value {
				for _, s := range strings.Split(t, " ") {
					if s == "" {
						continue
					}
					tokens = append(tokens, s)
				}
			}
		case int:
			tokens = append(tokens, strconv.Itoa(value))
		case bool:
			tokens = append(tokens, strconv.FormatBool(value))
		case uuid.UUID:
			// Skip the zero value. Even random UUID have some non-zero bits.
			if value == uuid.UUID([16]byte{}) {
				continue
			}
			tokens = append(tokens, value.String())
		}
	}
	return tokens
}
