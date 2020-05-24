package grabber

import (
	"fmt"
	"net/http"

	"grabber/models"
	pb "grabber/pb"

	log "github.com/sirupsen/logrus"
)

var schema = map[string]*Property{
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

func (grabber *Grabber) GetFacultyById(academy *models.Academy, id int32) (*pb.Faculty, error) {
	faculties, err := grabber.GetFaculties(academy, id)

	if err != nil {
		log.Error()
		return nil, err
	}

	if len(faculties) == 0 {
		log.Error()
		return nil, err
	}

	return faculties[0], nil
}

func (grabber *Grabber) GetFaculties(academy *models.Academy, id int32) ([]*pb.Faculty, error) {
	req, err := http.NewRequest(http.MethodGet, academy.Endpoint, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.URL.Path = "/Dek/Default.aspx"
	q := req.URL.Query()
	q.Add("mode", "facultet")
	q.Add("f", "facultet")
	if id != -1 {
		q.Add("id", fmt.Sprint(id))
	}
	req.URL.RawQuery = q.Encode()

	doc, err := grabber.GetDoc(req)
	if err != nil {
		log.Error(err)
	}

	var faculties []*pb.Faculty

	table := doc.Find("table#ContentPage_ucFacultets_Grid")

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
