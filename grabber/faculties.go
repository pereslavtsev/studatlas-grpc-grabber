package grabber

import (
	"fmt"

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
	params := map[string]string{
		"mode": "facultet",
		"f":    "facultet",
	}

	if id != -1 {
		params["id"] = fmt.Sprint(id)
	}

	doc, err := grabber.DoDictionaryReq(academy, params)

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
