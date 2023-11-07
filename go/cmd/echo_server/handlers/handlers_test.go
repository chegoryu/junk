package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chegoryu/junk/go/pkg/buildinfo"
	"github.com/stretchr/testify/require"
)

func checkHttpResponseHeaderValue(t *testing.T, rr *httptest.ResponseRecorder, headerName string, expectedHeaderValue string) {
	require.Equalf(
		t,
		expectedHeaderValue, rr.Header().Get(headerName),
		"handler returned unexpected '%s' header value", headerName,
	)
}

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ping)

	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	checkHttpResponseHeaderValue(t, rr, "Content-Type", "text/plain")
	require.Equal(t, "pong\n", rr.Body.String())
}

func TestVersion(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(version)

	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	checkHttpResponseHeaderValue(t, rr, "Content-Type", "text/plain")
	require.Equal(t, fmt.Sprintf("%s\n", buildinfo.ProgramVersion), rr.Body.String())
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
			require.NoError(t, err)

			req.Host = test.Host
			req.Header = test.Header

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(headers)

			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)
			checkHttpResponseHeaderValue(t, rr, "Content-Type", "text/plain")
			require.Equal(t, test.ExpectedBody, rr.Body.String())
		})
	}
}
