package icndb

import "encoding/json"

type jokeGetter interface {
	getURL(string) ([]byte, error)
}

type icndbServer struct{}

func (srv icndbServer) getURL(url string) ([]byte, error) {
	return getURL(url)
}

func getRandomJokeWithInterface(srv jokeGetter) (string, error) {
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
