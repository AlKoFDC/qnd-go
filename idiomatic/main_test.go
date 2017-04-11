package main

import (
	"errors"
	"html"
	"testing"
)

func resetGetURL(was func(string) ([]byte, error)) {
	getURL = was
}

func mockGetURL(string) ([]byte, error) {
	return []byte(`{"value":{"joke":"funny Chuck Norris Joke"}}`), nil
}

func TestShouldReceiveJokeFromICNDB(t *testing.T) {
	defer resetGetURL(getURL)
	getURL = mockGetURL

	joke, err := getRandomJoke()
	if err != nil {
		t.Fatal(err)
	}
	if len(joke) <= 0 {
		t.Error("Expected a joke, but didn't get any.")
	}
	t.Log(joke)
}

func TestShouldFail(t *testing.T) {
	joke, err := getRandomJoke()
	if err == nil {
		t.Errorf("Expected getRandomJoke to fail without internet connection, but got a joke: %s.", joke)
	}
	t.Log(err)
}

type mockedServerFunc func() (string, error)

func (srv mockedServerFunc) getURL(string) ([]byte, error) {
	b, e := srv()
	return []byte(b), e
}

var _ jokeGetter = (*mockedServerFunc)(nil)

func TestShouldGetMockedURL(t *testing.T) {
	for _, testIO := range []struct {
		desc          string
		expectedJoke  string
		expectedError error
	}{
		{expectedJoke: "funny mocked Chuck Norris Joke", desc: "joke1"},
		{expectedJoke: "&quot;funny&quot; mocked Chuck Norris Joke", desc: "joke3"},
		{expectedError: errors.New("some error"), desc: "joke2"},
	} {
		t.Run(testIO.desc, func(t *testing.T) {
			var getter mockedServerFunc = func() (string, error) {
				return `{"value":{"joke":"` + testIO.expectedJoke + `"}}`, testIO.expectedError
			}
			joke, err := getRandomJokeWithGetter(getter)
			if err != testIO.expectedError {
				t.Fatal(err)
			}
			if testIO.expectedError != nil {
				return
			}
			if len(joke) <= 0 {
				t.Error("Expected a joke, but didn't get any.")
			}
			if joke != html.UnescapeString(testIO.expectedJoke) {
				t.Errorf("Expected '%s', but got '%s'.", testIO.expectedJoke, joke)
			}
		})
	}
}
