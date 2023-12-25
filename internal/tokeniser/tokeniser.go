package tokeniser

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/satrap-illustrations/zs/internal/models"
)

type Token struct {
	Text, Field string
}

// Tokenise extracts tokens from a model
//
//nolint:revive
func Tokenise(m models.Model) []Token {
	tokens := []Token{}
	fields := m.Fields()
	for el := fields.Front(); el != nil; el = el.Next() {
		// When the data has other types, this needs to be extended
		switch value := m.ValueAtIdx(el.Value).(type) {
		case string:
			// allow searching for empty strings
			if value == "" {
				tokens = append(tokens, Token{Text: "", Field: el.Key})
				continue
			}

			for _, s := range strings.Split(value, " ") {
				if s == "" {
					continue
				}
				tokens = append(tokens, Token{
					Text:  normalise(s),
					Field: el.Key,
				})
			}
		case []string:
			for _, t := range value {
				for _, s := range strings.Split(t, " ") {
					if s == "" {
						continue
					}
					tokens = append(tokens, Token{
						Text:  normalise(s),
						Field: el.Key,
					})
				}
			}
		case int:
			tokens = append(tokens, Token{
				Text:  strconv.Itoa(value),
				Field: el.Key,
			})
		case bool:
			tokens = append(tokens, Token{
				Text:  strconv.FormatBool(value),
				Field: el.Key,
			})
		case uuid.UUID:
			// Skip the zero value. Even random UUID have some non-zero bits.
			if value == uuid.UUID([16]byte{}) {
				continue
			}
			tokens = append(tokens, Token{
				Text:  value.String(),
				Field: el.Key,
			})
		}
	}
	return tokens
}

// normalise applies the following transformation to a string:
// 1. Removes leading and traliing punctuation.
func normalise(s string) string {
	return strings.Trim(s, `?!.,;:"'_`)
}
