# rest-client

[![Build](https://img.shields.io/github/actions/workflow/status/no-src/rest-client/go.yml?branch=main)](https://github.com/no-src/rest-client/actions)
[![License](https://img.shields.io/github/license/no-src/rest-client)](https://github.com/no-src/rest-client/blob/main/LICENSE)

A command line tool for sending HTTP requests and displaying the response.

## Installation

The first need [Go](https://go.dev/doc/install) installed (**version 1.22+ is required**), then you can use the below
command to install `rest-client`.

```bash
go install github.com/no-src/rest-client/...@latest
```

## Quick Start

### Configuration

Create a `conf.yaml` file to define the custom variables, it is optional.

```yaml
host: http://127.0.0.1
secret: 123456
```

### Request

Create a `request.http` file to define the HTTP requests, it is required.

```text
### Test POST HTTP Request
POST {{host}}/say
Content-Type: application/json

{
  "content": "hello",
  "secret": "{{secret}}"
}

### Test GET HTTP Request
GET {{host}}/info
```

### Show Requests

```bash
rc -conf=conf.yaml -http=request.http
```

### Send Request

```bash
rc -conf=conf.yaml -http=request.http -send -id=1
```