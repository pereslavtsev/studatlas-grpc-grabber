package grabber

import (
	log "github.com/sirupsen/logrus"
	"grabber/models"
	pb "grabber/pb"
	"net/http"
)

var schemaReport = map[string]*Property{
	"id": {
		Type:   "id",
		Column: "Группа",
	},
	"name": {
		Type:   "text",
		Column: "Группа",
	},
	"year": {
		Type:   "numeric",
		Column: "Курс",
	},
	"speciality": {
		Type:   "speciality",
		Column: "Специальность",
	},
	"count": {
		Type:   "numeric",
		Column: "Студентов",
	},
	"curricula": {
		Type:   "text",
		Column: "Учебный План",
	},
}

func (g *Grabber) FetchReports(academy *models.Academy, params *DictionaryFilter) ([]*pb.Report, error) {
	req, err := http.NewRequest(http.MethodPost, academy.Endpoint, nil)
	req.URL.Path = "/Totals/Default.aspx"

	doc, err := g.GetDoc(req)

	if err != nil {
		log.Error(err)
	}

	var reports []*pb.Report

	table := doc.Find("table#ContentPage_GridGroup")

	//log.Debug(table.Html())

	NewGrid(table, schemaReport).EachRow(func(i int, row *Row) {
		reports = append(reports, &pb.Report{
			Id:         row.GetSpecificId("id", "group"),
			Name:       row.Get("name"),
			Year:       row.Get("year"),
			Speciality: row.Get("speciality"),
			Count:      row.GetInt32("count"),
			Curricula:  row.Get("curricula"),
		})
	})

	return reports, nil
}

func (g *Grabber) GetFacultyReports(academy *models.Academy, facultyId int32) ([]*pb.Report, error) {
	return g.FetchReports(academy, &DictionaryFilter{
		Mode:   "group",
		Filter: "facultet",
		ID:     facultyId,
	})
}
