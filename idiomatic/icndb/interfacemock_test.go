package icndb

import (
	"errors"
	"testing"
)

type mockedServerFunc func() (string, error)

func (srv mockedServerFunc) getURL(string) ([]byte, error) {
	b, e := srv()
	return []byte(b), e
}

var _ jokeGetter = (*mockedServerFunc)(nil)

func TestShouldGetMockedURL(t *testing.T) {
	for _, testIO := range []struct {
		desc          string
		jsonJoke      string
		expectedJoke  string
		expectedError error
	}{
		{desc: "joke", jsonJoke: "funny mocked Chuck Norris Joke", expectedJoke: "funny mocked Chuck Norris Joke"},
		{desc: "quoted", jsonJoke: "&quot;funny&quot; mocked Chuck Norris Joke", expectedJoke: `"funny" mocked Chuck Norris Joke`},
		{desc: "error", expectedError: errors.New("some error")},
	} {
		t.Run(testIO.desc, func(t *testing.T) {
			var getter mockedServerFunc = func() (string, error) {
				return `{"value":{"joke":"` + testIO.jsonJoke + `"}}`, testIO.expectedError
			}
			joke, err := getRandomJokeWithInterface(getter)
			if err != testIO.expectedError {
				t.Fatal(err)
			}
			if testIO.expectedError != nil {
				return
			}
			if len(joke) <= 0 {
				t.Error("Expected a joke, but didn't get any.")
			}
			if joke != testIO.expectedJoke {
				t.Errorf("Expected '%s', but got '%s'.", testIO.expectedJoke, joke)
			}
		})
	}
}
