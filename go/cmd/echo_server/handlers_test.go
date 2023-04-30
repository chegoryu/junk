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
			"SimpleHeader",
			http.Header{
				"Header": {"Value"},
			},
			"Header: Value\n",
		},
		{
			"MultiValueHeader",
			http.Header{
				"Header": {"Value2", "Value1", "Value3"},
			},
			"Header: Value1\nHeader: Value2\nHeader: Value3\n",
		},
		{
			"MultiHeaderXMultiValue",
			http.Header{
				"Header2":       {"Header2Value2", "Header2Value1", "Header2Value3"},
				"Header1":       {"Header1Value2", "Header1Value1"},
				"ZYXLastHeader": {"AAFirstValue"},
				"OtherHeader":   {"SomeValue"},
			},
			"Header1: Header1Value1\nHeader1: Header1Value2\n" +
				"Header2: Header2Value1\nHeader2: Header2Value2\nHeader2: Header2Value3\n" +
				"OtherHeader: SomeValue\n" +
				"ZYXLastHeader: AAFirstValue\n",
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
				t.Errorf("handler returned unexpected body:\nExpected:\n%s\nGot:\n%s\n", test.ResultBody, body)
			}
		})
	}
}
