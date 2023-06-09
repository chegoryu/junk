package handlers

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/chegoryu/junk/go/pkg/buildinfo"
	"github.com/chegoryu/junk/go/pkg/caseinsensitivecmp"
)

func AddHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/version", version)

	mux.HandleFunc("/headers", headers)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	io.WriteString(w, "pong\n")
}

func version(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "%s\n", buildinfo.ProgramVersion)
}

func headers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	headerCount := 0
	for _, headerValues := range r.Header {
		headerCount += len(headerValues)
	}
	if len(r.Host) > 0 {
		headerCount++
	}

	type Header struct {
		Name  string
		Value string
	}
	sortedHeaders := make([]Header, 0, headerCount)

	for headerName, headerValues := range r.Header {
		for _, headerValue := range headerValues {
			sortedHeaders = append(sortedHeaders, Header{
				Name:  headerName,
				Value: headerValue,
			})
		}
	}
	if len(r.Host) > 0 {
		sortedHeaders = append(sortedHeaders, Header{
			Name:  "host",
			Value: r.Host,
		})
	}

	sort.Slice(sortedHeaders, func(i int, j int) bool {
		arr := sortedHeaders[:]
		return caseinsensitivecmp.Less(arr[i].Name, arr[j].Name) ||
			(caseinsensitivecmp.Equal(arr[i].Name, arr[j].Name) && arr[i].Value < arr[j].Value)
	})

	for _, header := range sortedHeaders {
		fmt.Fprintf(w, "%s: %s\n", strings.ToLower(header.Name), header.Value)
	}
}
