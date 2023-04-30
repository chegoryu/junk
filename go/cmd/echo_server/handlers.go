package main

import (
	"fmt"
	"net/http"
	"sort"
)

func GetHeaders(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	headerCount := 0
	for _, headerValues := range r.Header {
		headerCount += len(headerValues)
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

	sort.Slice(sortedHeaders, func(i int, j int) bool {
		return sortedHeaders[i].Name < sortedHeaders[j].Name || (sortedHeaders[i].Name == sortedHeaders[j].Name && sortedHeaders[i].Value < sortedHeaders[j].Value)
	})

	for _, header := range sortedHeaders {
		fmt.Fprintf(w, "%s: %s\n", header.Name, header.Value)
	}
}
