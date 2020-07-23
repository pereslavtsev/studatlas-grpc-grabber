package grabber

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"grabber/models"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

var contextKey = "grabber"

// Grabber is the database operations wrapper
type Grabber struct {
	ctx    context.Context
	client *http.Client
}

type DictionaryFilter struct {
	Mode   string
	Filter string
	ID     int32
}

// Init initializes the grabber wrapper and create a new value context with the object in it.
func Init(ctx context.Context) context.Context {
	log.Infof(`Init %s...`, contextKey)

	return context.WithValue(ctx, contextKey, &Grabber{
		client: &http.Client{},
	})
}

// FromContext returns the grabber wrapper from a given context.
func FromContext(ctx context.Context) *Grabber {
	g, ok := ctx.Value(contextKey).(*Grabber)
	if !ok {
		log.Panic("Calling grabber.FromContext from a non-grabber context")
	}
	return g
}

func (g *Grabber) Do(r *http.Request) (*http.Response, error) {
	// log a http.Request instance
	log.WithFields(log.Fields{
		"method": r.Method,
		"query":  r.URL.RawQuery,
		"body":   r.Body,
	}).Debugf(`Sending a request to "%s"...`, r.URL.Hostname())

	switch r.Method {
	case http.MethodPost:
		{
			preReq, _ := http.NewRequestWithContext(r.Context(), http.MethodGet, r.URL.String(), nil)
			doc, _ := g.GetDoc(preReq)

			var data map[string]string

			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				log.Error("json.Unmarshal: ", err)
			}

			// скорректированное имя инпута для текущего сайта
			// без этого фикса никакие POST-запросы не будут проходить
			newData := make(map[string]string)
			for k, v := range data {
				fixedKey, exists := doc.Find(fmt.Sprintf(`[name*="%s"]`, k)).Attr("name")

				if !exists {
					log.Warnf(`Parameter "$s" hasn't corrected, a default value "%s" used instead.'`, k, v)
					continue
				}

				log.Debugf(`Field transformation: "%s" -> "%s"`, k, fixedKey)

				newData[fixedKey] = v
			}

			doc.Find("input[type=hidden]").Each(func(i int, s *goquery.Selection) {
				nameVal, nameExists := s.Attr("name")
				valVal, _ := s.Attr("value")

				if nameExists {
					newData[nameVal] = valVal
				}
			})

			postData := url.Values{}
			for k, v := range newData {
				postData.Set(k, v)
			}

			r, _ = http.NewRequest(http.MethodPost, r.URL.String(), strings.NewReader(postData.Encode()))
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			log.Debugf("Request body has been corrected to", postData)

			break
		}
	}

	// receive a response
	res, err := g.client.Do(r)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// TODO: get err
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	return fixCharset(res)
}

func (g *Grabber) GetDoc(req *http.Request) (*goquery.Document, error) {
	res, err := g.Do(req)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return goquery.NewDocumentFromReader(res.Body)
}

func (g *Grabber) DoDictionaryReq(academy *models.Academy, params *DictionaryFilter) (*goquery.Document, error) {
	req, err := http.NewRequest(http.MethodGet, academy.Endpoint, nil)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	req.URL.Path = "/Dek/Default.aspx"
	q := req.URL.Query()

	q.Add("mode", params.Mode)
	q.Add("f", params.Filter)
	q.Add("id", fmt.Sprint(params.ID))

	req.URL.RawQuery = q.Encode()

	return g.GetDoc(req)
}
