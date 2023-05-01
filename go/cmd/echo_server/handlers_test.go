package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func CheckHttpResponseCode(t *testing.T, rr *httptest.ResponseRecorder, expectedStatus int) {
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: expected %d, got %d", expectedStatus, status)
	}
}

func CheckHttpResponseBody(t *testing.T, rr *httptest.ResponseRecorder, expectedBody string) {
	if body := rr.Body.String(); body != expectedBody {
		t.Errorf("handler returned unexpected body:\nExpected:\n%s\nGot:\n%s\n", expectedBody, body)
	}
}

func CheckHttpResponseHeaderValue(t *testing.T, rr *httptest.ResponseRecorder, headerName string, expectedHeaderValue string) {
	if headerValue := rr.Header().Get(headerName); headerValue != expectedHeaderValue {
		t.Errorf("handler returned unexpected '%s' header value: expected '%s', got '%s'", headerName, expectedHeaderValue, headerValue)
	}
}

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)

	handler.ServeHTTP(rr, req)

	CheckHttpResponseCode(t, rr, http.StatusOK)
	CheckHttpResponseHeaderValue(t, rr, "Content-Type", "text/plain")
	CheckHttpResponseBody(t, rr, "pong\n")
}

func TestGetHeaders(t *testing.T) {
	type Header struct {
		Name  string
		Value string
	}

	var tests = []struct {
		Name         string
		Host         string
		Header       http.Header
		ExpectedBody string
	}{
		{
			"SimpleHeader",
			"",
			http.Header{
				"Header": {"Value"},
			},
			"Header: Value\n",
		},
		{
			"MultiValueHeader",
			"",
			http.Header{
				"Header": {"Value2", "Value1", "Value3"},
			},
			"Header: Value1\nHeader: Value2\nHeader: Value3\n",
		},
		{
			"MultiHeaderXMultiValue",
			"",
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
		{
			"HostHeader",
			"somehost.com:1234",
			http.Header{},
			"Host: somehost.com:1234\n",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/headers", nil)
			if err != nil {
				t.Error(err)
			}
			req.Host = test.Host
			req.Header = test.Header

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetHeaders)

			handler.ServeHTTP(rr, req)

			CheckHttpResponseCode(t, rr, http.StatusOK)
			CheckHttpResponseHeaderValue(t, rr, "Content-Type", "text/plain")
			CheckHttpResponseBody(t, rr, test.ExpectedBody)
		})
	}
}
