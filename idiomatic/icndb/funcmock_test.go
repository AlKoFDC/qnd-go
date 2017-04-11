package icndb

import "testing"

func resetGetURL(was func(string) ([]byte, error)) {
	getURL = was
}

func mockGetURL(string) ([]byte, error) {
	return []byte(`{"value":{"joke":"funny Chuck Norris Joke"}}`), nil
}

func TestShouldReceiveJokeFromICNDBWithFuncMock(t *testing.T) {
	defer resetGetURL(getURL)
	getURL = mockGetURL

	joke, err := GetRandomJoke()
	if err != nil {
		t.Fatal(err)
	}
	if len(joke) <= 0 {
		t.Error("Expected a joke, but didn't get any.")
	}
	t.Log(joke)
}
