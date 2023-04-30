package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHeaders(t *testing.T) {
	type Header struct {
		Name  string
		Value string
	}

	var tests = []struct {
		Name       string
		Header     http.Header
		ResultBody string
	}{
		{
			"One header",
			http.Header{
				"Header": {"Value"},
			},
			"Header: Value\n",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/headers", nil)
			if err != nil {
				t.Error(err)
			}
			req.Header = test.Header

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetHeaders)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: expected %d, got %d", http.StatusOK, status)
			}

			if body := rr.Body.String(); body != test.ResultBody {
				t.Errorf("handler returned unexpected body:\nExpected:\n %s\nGot:\n %s\n", test.ResultBody, body)
			}
		})
	}
}
