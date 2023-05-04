package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chegoryu/junk/go/pkg/buildinfo"
)

func checkHttpResponseCode(t *testing.T, rr *httptest.ResponseRecorder, expectedStatus int) {
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: expected %d, got %d", expectedStatus, status)
	}
}

func checkHttpResponseBody(t *testing.T, rr *httptest.ResponseRecorder, expectedBody string) {
	if body := rr.Body.String(); body != expectedBody {
		t.Errorf("handler returned unexpected body:\nExpected:\n%s\nGot:\n%s\n", expectedBody, body)
	}
}

func checkHttpResponseHeaderValue(t *testing.T, rr *httptest.ResponseRecorder, headerName string, expectedHeaderValue string) {
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
	handler := http.HandlerFunc(ping)

	handler.ServeHTTP(rr, req)

	checkHttpResponseCode(t, rr, http.StatusOK)
	checkHttpResponseHeaderValue(t, rr, "Content-Type", "text/plain")
	checkHttpResponseBody(t, rr, "pong\n")
}

func TestVersion(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(version)

	handler.ServeHTTP(rr, req)

	checkHttpResponseCode(t, rr, http.StatusOK)
	checkHttpResponseHeaderValue(t, rr, "Content-Type", "text/plain")
	checkHttpResponseBody(t, rr, fmt.Sprintf("%s\n", buildinfo.ProgramVersion))
}

func TestHeaders(t *testing.T) {
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
			handler := http.HandlerFunc(headers)

			handler.ServeHTTP(rr, req)

			checkHttpResponseCode(t, rr, http.StatusOK)
			checkHttpResponseHeaderValue(t, rr, "Content-Type", "text/plain")
			checkHttpResponseBody(t, rr, test.ExpectedBody)
		})
	}
}
