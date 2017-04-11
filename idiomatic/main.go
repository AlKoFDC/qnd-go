package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(getRandomJoke())
}

var getURL = func(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

type icndbServer struct{}

func (srv icndbServer) getURL(url string) ([]byte, error) {
	return getURL(url)
}

func getRandomJoke() (string, error) {
	return getRandomJokeWithGetter(icndbServer{})
}

type jokeGetter interface {
	getURL(string) ([]byte, error)
}

func getRandomJokeWithGetter(srv jokeGetter) (string, error) {
	const (
		sfw           = "?exclude=[explicit]"
		randomJokeURL = "https://api.icndb.com/jokes/random" + sfw
	)
	jokeObject, err := srv.getURL(randomJokeURL)
	if err != nil {
		return "", err
	}

	var j jokeEntry
	if err := json.Unmarshal(jokeObject, &j); err != nil {
		return "", err
	}

	return j.String(), nil
}

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
