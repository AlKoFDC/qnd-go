package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(serveRandomJoke))
	defer srv.Close()

	res, err := http.Get(srv.URL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	t.Log(string(body))
}
