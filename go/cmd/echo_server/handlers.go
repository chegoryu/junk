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
			Name:  "Host",
			Value: r.Host,
		})
	}

	sort.Slice(sortedHeaders, func(i int, j int) bool {
		arr := sortedHeaders[:]
		return arr[i].Name < arr[j].Name || (arr[i].Name == arr[j].Name && arr[i].Value < arr[j].Value)
	})

	for _, header := range sortedHeaders {
		fmt.Fprintf(w, "%s: %s\n", header.Name, header.Value)
	}
}
