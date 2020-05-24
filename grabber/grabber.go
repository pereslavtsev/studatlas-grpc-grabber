package grabber

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
)

type contextKeyType string

var contextKey = contextKeyType("grabber")

// Grabber is the database operations wrapper
type Grabber struct {
	ctx    context.Context
	client *http.Client
}

// Init initializes the grabber wrapper and create a new value context with the object in it.
func Init(ctx context.Context) context.Context {
	log.Infof(`Init %s...`, contextKey)

	return context.WithValue(ctx, contextKey, &Grabber{
		client: &http.Client{},
	})
}

// FromContext returns the database wrapper from a given context.
func FromContext(ctx context.Context) *Grabber {
	db, ok := ctx.Value(contextKey).(*Grabber)
	if !ok {
		log.Panic("Calling grabber.FromContext from a non-grabber context")
	}
	return db
}

func (grabber *Grabber) Do(req *http.Request) (*http.Response, error) {
	// log a http.Request instance
	log.WithFields(log.Fields{
		"method": req.Method,
		"query":  req.URL.RawQuery,
		"body":   req.Body,
	}).Debugf("Sending a request to \"%s\"...", req.URL.Hostname())

	// receive a response
	res, err := grabber.client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// TODO: get err
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Fix charset
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