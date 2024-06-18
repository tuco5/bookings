package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var tests = []struct {
	name     string
	expected string
	status   int
}{
	{"home page", "This is a home page", http.StatusOK},
	{"about page", "This is the about page", http.StatusOK},
}

func TestHandlers(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Create a request to pass to the handler
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Repo.Home)

			// Call the handler
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}

			// Check the response body
			if rr.Body.String() != tt.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expected)
			}
		})
	}
}
