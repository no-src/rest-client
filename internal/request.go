package internal

import "net/http"

type Request struct {
	Name     string
	Method   string
	URL      string
	Protocol string
	Header   http.Header
	Body     string
}

func newRequest(name string) *Request {
	req := &Request{
		Name:   name,
		Header: make(http.Header),
	}
	return req
}
