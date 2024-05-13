package internal

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

const (
	headerContentType        = "Content-Type"
	defaultHeaderContentType = "application/json"
	defaultHTTPProtocol      = "HTTP/1.1"
)

func send(request *Request) (body []byte, err error) {
	if request == nil {
		return body, errors.New("invalid request")
	}
	req, err := http.NewRequest(request.Method, request.URL, strings.NewReader(request.Body))
	if err != nil {
		return body, err
	}
	req.Header = request.Header
	if len(req.Header.Get(headerContentType)) == 0 {
		req.Header.Set(headerContentType, defaultHeaderContentType)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func showHeader(header http.Header) string {
	b := strings.Builder{}
	for key, values := range header {
		for _, v := range values {
			b.WriteString(fmt.Sprintf("%s: %s\n", key, v))
		}
	}
	return b.String()
}

func isLegalHTTPMethod(method string) bool {
	allows := []string{
		http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace,
	}
	return slices.Contains(allows, method)
}

func isLegalURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
