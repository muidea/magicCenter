package net

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime"
	"net/http"

	"muidea.com/magicCenter/foundation/util"
)

type maxBytesReader struct {
	res http.ResponseWriter
	req io.ReadCloser // underlying reader
	n   int64         // max bytes remaining
	err error         // sticky error
}

func (l *maxBytesReader) tooLarge() (n int, err error) {
	l.err = errors.New("http: request body too large")
	return 0, l.err
}

func (l *maxBytesReader) Read(p []byte) (n int, err error) {
	if l.err != nil {
		return 0, l.err
	}
	if len(p) == 0 {
		return 0, nil
	}
	// If they asked for a 32KB byte read but only 5 bytes are
	// remaining, no need to read 32KB. 6 bytes will answer the
	// question of the whether we hit the limit or go past it.
	if int64(len(p)) > l.n+1 {
		p = p[:l.n+1]
	}
	n, err = l.req.Read(p)

	if int64(n) <= l.n {
		l.n -= int64(n)
		l.err = err
		return n, err
	}

	n = int(l.n)
	l.n = 0

	// The server code and client code both use
	// maxBytesReader. This "requestTooLarge" check is
	// only used by the server code. To prevent binaries
	// which only using the HTTP Client code (such as
	// cmd/go) from also linking in the HTTP server, don't
	// use a static type assertion to the server
	// "*response" type. Check this interface instead:
	type requestTooLarger interface {
		requestTooLarge()
	}
	if res, ok := l.res.(requestTooLarger); ok {
		res.requestTooLarge()
	}
	l.err = errors.New("http: request body too large")
	return n, l.err
}

func (l *maxBytesReader) Close() error {
	return l.req.Close()
}

// ParsePostJSON 解析http post请求提交的json数据
func ParsePostJSON(req *http.Request, param interface{}) error {
	util.ValidataPtr(param)

	if req.Body == nil {
		return errors.New("missing form body")
	}

	contentType, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	switch {
	case contentType == "application/json":
		var reader io.Reader = req.Body
		maxFormSize := int64(1<<63 - 1)
		if _, ok := req.Body.(*maxBytesReader); !ok {
			maxFormSize = int64(10 << 20) // 10 MB is a lot of text.
			reader = io.LimitReader(req.Body, maxFormSize+1)
		}

		payload, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		if int64(len(payload)) > maxFormSize {
			return errors.New("http: POST too large")
		}

		err = json.Unmarshal(payload, param)
		if err != nil {
			return err
		}

	default:
		return errors.New("invalid contentType, contentType:" + contentType)
	}

	return nil
}
