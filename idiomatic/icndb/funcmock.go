package icndb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var getURL = func(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func getRandomJokeWithFunction() (string, error) {
	const (
		sfw           = "?exclude=[explicit]"
		randomJokeURL = "https://api.icndb.com/jokes/random" + sfw
	)
	jokeObject, err := getURL(randomJokeURL)
	if err != nil {
		return "", err
	}

	var j jokeEntry
	if err := json.Unmarshal(jokeObject, &j); err != nil {
		return "", err
	}

	return j.String(), nil
}
