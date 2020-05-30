package grabber

import (
	"grabber/models"
	pb "grabber/pb"

	log "github.com/sirupsen/logrus"
)

var schemaSpeciality = map[string]*Property{
	"id": {
		Type:   "id",
		Column: "Специальность",
	},
	"shortName": {
		Type:   "text",
		Column: "Специальность",
	},
	"name": {
		Type:   "text",
		Column: "Название Специальности",
	},
	"facultyId": {
		Type:   "id",
		Column: "Факультет",
	},
	"faculty": {
		Type:   "text",
		Column: "Факультет",
	},
	"divisionId": {
		Type:   "id",
		Column: "Кафедра",
	},
	"division": {
		Type:   "text",
		Column: "Кафедра",
	},
	"code": {
		Type:   "text",
		Column: "Шифр",
	},
	"qualification": {
		Type:   "text",
		Column: "Квалификация",
	},
}

func (g *Grabber) FetchSpecialities(academy *models.Academy, params *DictionaryFilter) ([]*pb.Speciality, error) {
	doc, err := g.DoDictionaryReq(academy, params)

	if err != nil {
		log.Error(err)
	}

	var specialities []*pb.Speciality

	table := doc.Find("table#ContentPage_ucSpets_Grid")

	NewGrid(table, schemaSpeciality).EachRow(func(i int, row *Row) {
		specialities = append(specialities, &pb.Speciality{
			Id:            row.GetId("id"),
			ShortName:     row.Get("shortName"),
			Name:          row.Get("name"),
			Code:          row.Get("code"),
			Qualification: row.Get("qualification"),
			DivisionId:    row.GetId("divisionId"),
			Division:      row.Get("division"),
			FacultyId:     row.GetId("facultyId"),
			Faculty:       row.Get("faculty"),
		})
	})

	return specialities, nil
}

func (g *Grabber) GetSpecialities(academy *models.Academy) ([]*pb.Speciality, error) {
	specialities, err := g.FetchSpecialities(academy, &DictionaryFilter{
		Mode:   "spets",
		Filter: "spets",
	})
	return specialities, err
}

func (g *Grabber) GetFacultySpecialities(academy *models.Academy, facultyId int32) ([]*pb.Speciality, error) {
	specialities, err := g.FetchSpecialities(academy, &DictionaryFilter{
		Mode:   "spets",
		Filter: "facultet",
		ID:     facultyId,
	})
	return specialities, err
}

func (g *Grabber) GetSpecialityById(academy *models.Academy, id int32) (*pb.Speciality, error) {
	specialities, err := g.FetchSpecialities(academy, &DictionaryFilter{
		Mode:   "spets",
		Filter: "spets",
		ID:     id,
	})
	return specialities[0], err
}
