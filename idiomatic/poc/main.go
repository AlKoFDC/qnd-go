package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	//fmt.Println(getRandom())
	fmt.Println(getQuoteJoke())

	http.HandleFunc("/random", serveRandomJoke)
	if err := http.ListenAndServe(":1408", nil); err != nil {
		fmt.Println(err)
	}
}

func serveRandomJoke(w http.ResponseWriter, r *http.Request) {
	_, joke := getRandom()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(joke + "\n"))
}

type entry struct {
	Type  string
	Value struct {
		ID         int
		Joke       string
		Categories []string
	}
}

const (
	icndbURL = "https://api.icndb.com"
)

func getRandom() (id int, joke string) {
	const (
		sfw       = "?exclude=[explicit]"
		randomURL = icndbURL + "/jokes/random" + sfw
	)
	resp, err := http.Get(randomURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var j entry
	json.Unmarshal(body, &j)
	return j.Value.ID, html.UnescapeString(j.Value.Joke)
}

func getQuoteJoke() string {
	var (
		quoteJokes   = []int{24, 34, 394, 264, 426}
		quoteExample = icndbURL + "/jokes/" + strconv.Itoa(quoteJokes[0])
	)
	resp, err := http.Get(quoteExample)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var j entry
	json.Unmarshal(body, &j)
	return html.UnescapeString(j.Value.Joke)
}
