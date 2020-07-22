package grabber

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net/http"
)

func fixCharset(res *http.Response) (*http.Response, error) {
	utf8, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
	if err != nil {
		log.Error("IO error:", err)
		return nil, err
	}

	// Fix charset
	body, err := ioutil.ReadAll(utf8)
	if err != nil {
		log.Error("IO error:", err)
		return nil, err
	}

	// Replace a response body
	res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return res, nil
}
