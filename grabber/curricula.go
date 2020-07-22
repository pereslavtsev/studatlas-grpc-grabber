package grabber

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"grabber/models"
	pb "grabber/pb"
	"net/http"
)

var schemaCurricula = map[string]*Property{
	"id": {
		Type:    "id",
		Columns: []string{"Имя плана", "Веб-версия"},
	},
	"name": {
		Type:    "text",
		Columns: []string{"Имя плана", "Веб-версия"},
	},
	"speciality": {
		Type:   "text",
		Column: "Специальность",
	},
}

func (g *Grabber) GetFacultyCurricula(academy *models.Academy, in *pb.ListFacultyCurriculaRequest) ([]*pb.Curricula, error) {
	reqBody, err := json.Marshal(map[string]string{
		cmb("Facultets"): fmt.Sprint(in.FacultyId),
		cmb("Years"):     in.Years,
	})

	req, err := http.NewRequest(http.MethodPost, academy.Endpoint, bytes.NewBuffer(reqBody))
	// set url
	req.URL.Path = "/Plans/Default.aspx"
	doc, err := g.GetDoc(req)

	if err != nil {
		log.Error(err)
	}

	var curricula []*pb.Curricula
	table := doc.Find("table#ContentPage_Grid")

	if table == nil {
		log.Error("Таблица с учебными планами не найдена.")
		return nil, nil
	}

	NewGrid(table, schemaCurricula).EachRow(func(i int, row *Row) {
		curricula = append(curricula, &pb.Curricula{
			Id:         row.GetId("id"),
			Name:       row.Get("name"),
			Speciality: row.Get("speciality"),
		})
	})

	return curricula, nil
}
