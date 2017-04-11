package icndb

import "testing"

func TestShouldFail(t *testing.T) {
	// Run this test without an internet connection, otherwise it will fail!
	joke, err := GetRandomJoke()
	if err == nil {
		t.Errorf("Expected getRandomJoke to fail without internet connection, but got a joke: %s.", joke)
	}
	t.Log(err)
}
