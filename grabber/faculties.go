package grabber

import (
	"grabber/models"
	pb "grabber/pb"
	"grabber/schemas"

	log "github.com/sirupsen/logrus"
)

func (g *Grabber) FetchFaculties(academy *models.Academy, params *DictionaryFilter) ([]*pb.Faculty, error) {
	doc, err := g.DoDictionaryReq(academy, params)

	if err != nil {
		log.Error(err)
	}

	var faculties []*pb.Faculty

	table := doc.Find("table#ContentPage_ucFacultets_Grid")

	NewGrid(table, schemas.Faculty).EachRow(func(i int, row *Row) {
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

func (g *Grabber) GetFacultyById(academy *models.Academy, id int32) (*pb.Faculty, error) {
	faculties, err := g.FetchFaculties(academy, &DictionaryFilter{
		Mode:   "facultet",
		Filter: "facultet",
		ID:     id,
	})

	if err != nil {
		log.Error()
		return nil, err
	}

	if len(faculties) == 0 {
		log.Errorf(`Faculty with ID "%i" has not found`, id)
		return nil, err
	}

	return faculties[0], err
}

func (g *Grabber) GetFaculties(academy *models.Academy) ([]*pb.Faculty, error) {
	faculties, err := g.FetchFaculties(academy, &DictionaryFilter{
		Mode:   "facultet",
		Filter: "facultet",
	})
	return faculties, err
}
