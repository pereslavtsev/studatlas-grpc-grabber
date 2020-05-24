package grabber

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"grabber/models"
	pb "grabber/pb"
	"net/http"
)

var schemaDivision = map[string]*Property{
	"id": {
		Type:   "id",
		Column: "Номер",
	},
	"name": {
		Type:   "text",
		Column: "Название",
	},
	"abbreviation": {
		Type:   "text",
		Column: "Сокращение",
	},
	"head": {
		Type:   "text",
		Column: "ЗавКафедрой",
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

func (grabber *Grabber) GetDivisionById(academy *models.Academy, id int32) (*pb.Division, error) {
	divisions, err := grabber.GetDivisions(academy, id)

	if err != nil {
		log.Error()
		return nil, err
	}

	if len(divisions) == 0 {
		log.Error()
		return nil, err
	}

	return divisions[0], nil
}

func (grabber *Grabber) GetDivisions(academy *models.Academy, id int32) ([]*pb.Division, error) {
	req, err := http.NewRequest(http.MethodGet, academy.Endpoint, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.URL.Path = "/Dek/Default.aspx"
	q := req.URL.Query()
	q.Add("mode", "kaf")
	q.Add("f", "kaf")
	if id != -1 {
		q.Add("id", fmt.Sprint(id))
	}
	req.URL.RawQuery = q.Encode()

	doc, err := grabber.GetDoc(req)
	if err != nil {
		log.Error(err)
	}

	var divisions []*pb.Division

	table := doc.Find("table#ContentPage_ucKaf_Grid")

	NewGrid(table, schemaDivision).EachRow(func(i int, row *Row) {
		divisions = append(divisions, &pb.Division{
			Id:           row.GetId("id"),
			Name:         row.Get("name"),
			Abbreviation: row.Get("abbreviation"),
			Head:         row.Get("head"),
			Phone:        row.Get("phone"),
			Room:         row.Get("room"),
		})
	})

	return divisions, nil
}
