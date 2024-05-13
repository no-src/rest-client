package internal

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

var (
	errRequestNameNotFound = errors.New("request name not found")
	errInvalidRequestLine  = errors.New("invalid request line")
	errInvalidHeader       = errors.New("invalid header")
	errInvalidHTTPMethod   = errors.New("invalid http method")
	errInvalidURL          = errors.New("invalid http url")
)

func parseHttp(text string) (requests []*Request, err error) {
	var request *Request
	step := stepInit
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		var ok bool
		ok, request, step, err = parse(line, request, step)
		if err != nil {
			return nil, err
		}
		if ok && request != nil {
			requests = append(requests, request)
			request = nil
			ok, request, step, err = parse(line, request, step)
			if err != nil {
				return nil, err
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	} else {
		if step >= stepRequestLine && request != nil {
			requests = append(requests, request)
		}
	}
	return requests, nil
}

func parse(line string, request *Request, step int) (bool, *Request, int, error) {
	ok := false
	switch step {
	case stepInit:
		name, err := parseName(line)
		if err == nil {
			request = newRequest(name)
		} else {
			return ok, request, step, err
		}
		step = stepName
	case stepName:
		method, url, protocol, err := parseRequestLine(line)
		if err == nil {
			request.Method = method
			request.URL = url
			request.Protocol = protocol
			step = stepRequestLine
		} else {
			return ok, request, step, err
		}
	case stepRequestLine, stepHeader:
		isEnd, key, value, err := parseHeader(line)
		if err == nil {
			if isEnd {
				step = stepHeaderEnd
			} else {
				request.Header.Add(key, value)
				step = stepHeader
			}
		} else {
			return ok, request, step, err
		}
	case stepHeaderEnd, stepBody:
		isEnd, body, err := parseBody(line)
		if err == nil {
			if isEnd {
				step = stepInit
				ok = true
			} else {
				request.Body += body + "\n"
				step = stepBody
			}
		} else {
			return ok, request, step, err
		}
	}
	return ok, request, step, nil
}

func parseName(line string) (string, error) {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "###") {
		return strings.TrimSpace(strings.TrimLeft(line, "#")), nil
	}
	return "", errRequestNameNotFound
}

func parseRequestLine(line string) (method string, url string, protocol string, err error) {
	line = initLine(line)
	args := strings.Split(line, " ")
	length := len(args)
	if length < 2 || length > 3 {
		err = errInvalidRequestLine
		return
	}
	if !isLegalHTTPMethod(args[0]) {
		err = fmt.Errorf("%w => %s", errInvalidHTTPMethod, args[0])
		return
	}
	if !isLegalURL(args[1]) {
		err = fmt.Errorf("%w => %s", errInvalidURL, args[1])
		return
	}
	method = args[0]
	url = args[1]
	protocol = defaultHTTPProtocol

	if length == 3 {
		protocol = args[2]
	}
	return
}

func parseHeader(line string) (isEnd bool, key, value string, err error) {
	line = initLine(line)
	if len(line) == 0 {
		isEnd = true
		return
	}
	args := strings.Split(line, ": ")
	if len(args) != 2 {
		err = errInvalidHeader
		return
	}
	key = args[0]
	value = args[1]
	return
}

func parseBody(line string) (isEnd bool, body string, err error) {
	line = initLine(line)
	if strings.HasPrefix(line, "###") {
		return true, "", nil
	}
	return false, line, nil
}

func initLine(line string) string {
	line = strings.TrimSpace(line)
	return parseVariables(line)
}
