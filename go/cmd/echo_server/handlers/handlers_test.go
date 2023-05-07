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
			"EmptyHeaders",
			"",
			http.Header{},
			"",
		},
		{
			"SimpleHeader",
			"",
			http.Header{
				"Header": {
					"Value",
				},
			},
			"header: Value\n",
		},
		{
			"MultiValueHeader",
			"",
			http.Header{
				"Header": {
					"Value2",
					"Value1",
					"Value3",
				},
			},
			"header: Value1\n" +
				"header: Value2\n" +
				"header: Value3\n",
		},
		{
			"MultiHeaderXMultiValue",
			"",
			http.Header{
				"Header2": {
					"Header2Value2",
					"Header2Value1",
					"Header2Value3",
				},
				"Header1": {
					"Header1Value2",
					"Header1Value1",
				},
				"ZYXLastHeader": {
					"AAFirstValue",
				},
				"OtherHeader": {
					"SomeValue",
				},
			},
			"header1: Header1Value1\n" +
				"header1: Header1Value2\n" +
				"header2: Header2Value1\n" +
				"header2: Header2Value2\n" +
				"header2: Header2Value3\n" +
				"otherheader: SomeValue\n" +
				"zyxlastheader: AAFirstValue\n",
		},
		{
			"HostHeader",
			"somehost.com:1234",
			http.Header{},
			"host: somehost.com:1234\n",
		},
		{
			"CaseInsensitiveHeaderNameCmp",
			"",
			http.Header{
				"aHeader": {
					"aHeaderValue",
				},
				"BHeader": {
					"BHeaderValue",
				},
				"cHeader": {
					"cHeaderValue",
				},
				"DHeader": {
					"DHeaderValue",
				},
			},
			"aheader: aHeaderValue\n" +
				"bheader: BHeaderValue\n" +
				"cheader: cHeaderValue\n" +
				"dheader: DHeaderValue\n",
		},
		{
			"CaseSensitiveHeaderValueCmp",
			"",
			http.Header{
				"Header": {
					"aHeaderValue",
					"BHeaderValue",
					"cHeaderValue",
					"DHeaderValue",
				},
			},
			"header: BHeaderValue\n" +
				"header: DHeaderValue\n" +
				"header: aHeaderValue\n" +
				"header: cHeaderValue\n",
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
