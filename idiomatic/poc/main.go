package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// You can call the server by executing: curl localhost:1408/random
func main() {
	errorChannel := make(chan error, 1)
	defer close(errorChannel)

	go func() {
		myHandler := http.NewServeMux()
		myHandler.HandleFunc("/random", serveRandomJoke)

		myServer := &http.Server{
			Addr:    ":1408",
			Handler: myHandler,
		}
		if err := myServer.ListenAndServe(); err != nil {
			errorChannel <- err
		}
	}()

	// Handle SIGINT and SIGTERM for graceful shutdowns.
	systemInterruptChannel := make(chan os.Signal)
	defer close(systemInterruptChannel)
	signal.Notify(systemInterruptChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errorReturn, ok := <-errorChannel:
		if ok {
			fmt.Println(errorReturn)
		} // else: channel closed, we're done successfully
	case interruptSignal := <-systemInterruptChannel:
		fmt.Printf("Exit due to signal %s\n", interruptSignal)
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
