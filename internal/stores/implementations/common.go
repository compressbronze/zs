package implementations

import (
	"encoding/json"
	"errors"
	"os"
)

var ErrInvalidDocType = errors.New("invalid document type")

func readJSONFile(path string, v any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(v)
}
