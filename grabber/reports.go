package grabber

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"grabber/models"
	pb "grabber/pb"
	"grabber/schemas"
	"net/http"
)

func (g *Grabber) FetchReports(academy *models.Academy, in *pb.ListFacultyReportsRequest) ([]*pb.Report, error) {
	reqBody, err := json.Marshal(map[string]string{
		cmb("Facultets"): fmt.Sprint(in.FacultyId),
		cmb("Years"):     in.Years,
	})

	req, err := http.NewRequest(http.MethodPost, academy.Endpoint, bytes.NewBuffer(reqBody))
	// set url
	req.URL.Path = "/Totals/Default.aspx"
	doc, err := g.GetDoc(req)

	if err != nil {
		log.Error(err)
	}

	var reports []*pb.Report

	table := doc.Find("table#ContentPage_GridGroup")

	NewGrid(table, schemas.Report).EachRow(func(i int, row *Row) {
		reports = append(reports, &pb.Report{
			Id:         row.GetSpecificId("id", "group"),
			Name:       row.Get("name"),
			Year:       row.GetInt32("year"),
			Speciality: row.Get("speciality"),
			Count:      row.GetInt32("count"),
			Curricula:  row.Get("curricula"),
		})
	})

	return reports, nil
}
