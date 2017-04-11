package icndb

import (
	"encoding/json"
	"html"
)

type jokeEntry string

var _ json.Unmarshaler = (*jokeEntry)(nil)

func (je jokeEntry) String() string {
	return string(je)
}

func (je *jokeEntry) UnmarshalJSON(value []byte) error {
	var jsonEntry struct {
		Value struct {
			Joke string
		}
	}
	if err := json.Unmarshal(value, &jsonEntry); err != nil {
		return err
	}
	*je = jokeEntry(html.UnescapeString(jsonEntry.Value.Joke))
	return nil
}
