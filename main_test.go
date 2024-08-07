package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// Mock HTTP client for testing
type mockHTTPClient struct {
	doFunc func(req *http.Request) (*http.Response, error)
}

func (c *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.doFunc(req)
}

func TestFetchJoke(t *testing.T) {
	tests := []struct {
		name          string
		mockResponse  *http.Response
		mockError     error
		expectedJoke  *Joke
		expectedError error
	}{
		{
			name: "Successful Joke Fetch",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(`{"id":"123","joke":"Why did the scarecrow win an award? Because he was outstanding in his field!","status":200}`)),
			},
			expectedJoke: &Joke{
				Id:     "123",
				Joke:   "Why did the scarecrow win an award? Because he was outstanding in his field!",
				Status: 200,
			},
			expectedError: nil,
		},
		{
			name: "Joke Not Found",
			mockResponse: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       ioutil.NopCloser(strings.NewReader(`{"status":404}`)),
			},
			expectedJoke:  nil,
			expectedError: errors.New("Joke not found"),
		},
		{
			name:          "Error in HTTP Request",
			mockError:     errors.New("network error"),
			expectedJoke:  nil,
			expectedError: errors.New("network error"),
		},
		{
			name: "Error in Response Parsing",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(`{invalid json`)),
			},
			expectedJoke:  nil,
			expectedError: errors.New("invalid character 'i' looking for beginning of object key string"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockHTTPClient{
				doFunc: func(req *http.Request) (*http.Response, error) {
					return tt.mockResponse, tt.mockError
				},
			}
			joke, err := FetchJoke("https://icanhazdadjoke.com/", client)
			if err != nil && (tt.expectedError == nil || err.Error() != tt.expectedError.Error()) {
				t.Fatalf("Expected error %v, got %v", tt.expectedError, err)
			}
			if joke != nil && (tt.expectedJoke == nil || *joke != *tt.expectedJoke) {
				t.Fatalf("Expected joke %v, got %v", tt.expectedJoke, joke)
			}
		})
	}
}
