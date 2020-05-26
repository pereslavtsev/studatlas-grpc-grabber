package grabber

import (
	log "github.com/sirupsen/logrus"
	"grabber/models"
	pb "grabber/pb"
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

func (g *Grabber) FetchDivisions(academy *models.Academy, params *DictionaryFilter) ([]*pb.Division, error) {
	doc, err := g.DoDictionaryReq(academy, params)

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

func (g *Grabber) GetDivisionById(academy *models.Academy, id int32) (*pb.Division, error) {
	divisions, err := g.FetchDivisions(academy, &DictionaryFilter{
		Mode:   "kaf",
		Filter: "kaf",
		ID:     id,
	})

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

func (g *Grabber) GetDivisions(academy *models.Academy) ([]*pb.Division, error) {
	doc, err := g.DoDictionaryReq(academy, &DictionaryFilter{
		Mode:   "kaf",
		Filter: "kaf",
	})

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
