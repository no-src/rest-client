package internal

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/no-src/log"
)

type handler func(id int, request *Request) (err error)

func sendHandler(userRequestId int) handler {
	return func(id int, request *Request) (err error) {
		if id != userRequestId {
			return
		}
		body, err := send(request)
		if err == nil {
			log.Log("[OK] send http request success\n\nrequest:\n%s\nresponse:\n%s", renderRequest(request), renderBody(body))
		}
		return err
	}
}

func showHandler(userRequestId int) handler {
	return func(id int, request *Request) (err error) {
		if userRequestId > 0 && id != userRequestId {
			return nil
		}
		log.Log("http request [%d]\n"+renderRequest(request), id)
		return nil
	}
}

func renderRequest(request *Request) string {
	return fmt.Sprintf("### %s\n%s %s %s\n%s\n%s\n", request.Name, request.Method, request.URL, request.Protocol, showHeader(request.Header), renderBody([]byte(request.Body)))
}

func renderBody(body []byte) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, body, "", "  ")
	if err != nil {
		return string(body)
	}
	return buf.String()
}
