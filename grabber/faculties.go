package grabber

import (
	"net/http"

	"grabber/models"
	pb "grabber/pb"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

func (grabber *Grabber) GetFaculties(academy *models.Academy) ([]*pb.Faculty, error) {
	req, err := http.NewRequest(http.MethodGet, academy.Endpoint, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.URL.Path = "/Dek/Default.aspx"
	q := req.URL.Query()
	q.Add("mode", "facultet")
	req.URL.RawQuery = q.Encode()

	res, err := grabber.Do(req)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var faculties []*pb.Faculty

	table := doc.Find("table#ContentPage_ucFacultets_Grid").First()

	schema := map[string]*Property{
		"id": {
			Type:    "id",
			Columns: []string{"Факультет", "Институт"},
		},
		"name": {
			Type:    "text",
			Columns: []string{"Факультет", "Институт"},
		},
		"abbreviation": {
			Type:   "text",
			Column: "Сокращение",
		},
		"head": {
			Type:   "text",
			Column: "Декан",
		},
		"phone": {
			Type:   "text",
			Column: "Телефон",
		},
		"room": {
			Type:   "text",
			Column: "Аудитория",
		},
	}

	NewGrid(table, schema).EachRow(func(i int, row *Row) {
		faculties = append(faculties, &pb.Faculty{
			Id:           row.GetId("id"),
			Name:         row.Get("name"),
			Abbreviation: row.Get("abbreviation"),
			Head:         row.Get("head"),
			Phone:        row.Get("phone"),
			Room:         row.Get("room"),
		})
	})

	return faculties, nil
}
