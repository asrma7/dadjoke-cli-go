package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// Define the interface for making HTTP requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Define the constants
const usage = `
Fetch random dad jokes from icanhazdadjoke.com

Usage of dadjokes:

dadjokes [flags]
-j, --joke Fetch a dad joke by ID
-h, --help prints help information
`

// Joke represents the structure of the joke response
type Joke struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// FetchJoke fetches a joke from the given URL using the provided HTTP client
func FetchJoke(url string, client HTTPClient) (*Joke, error) {
	userAgent := "dadjokes-cli (https://github.com/asrma7/dadjoke-cli-go)"
	acceptHeader := "application/json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", acceptHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		return nil, err
	}

	if joke.Status == 404 {
		return nil, errors.New("Joke not found")
	}

	return &joke, nil
}

// Main function to handle command-line arguments and fetch jokes
func main() {
	var jokeId string
	var showId bool
	flag.StringVar(&jokeId, "joke", "", "Fetch a dad joke by ID")
	flag.StringVar(&jokeId, "j", "", "Fetch a dad joke by ID")
	flag.BoolVar(&showId, "id", false, "Show joke id")
	flag.BoolVar(&showId, "i", false, "Show joke id")
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	var url string
	if jokeId != "" {
		url = fmt.Sprintf("https://icanhazdadjoke.com/j/%s", jokeId)
	} else {
		url = "https://icanhazdadjoke.com/"
	}

	client := &http.Client{}
	joke, err := FetchJoke(url, client)
	if err != nil {
		fmt.Println("Error fetching joke: ", err)
		os.Exit(1)
	}

	if showId {
		fmt.Printf("%v\n\nJoke ID for reference: %v\n", joke.Joke, joke.Id)
	} else {
		fmt.Println(joke.Joke)
	}
}
