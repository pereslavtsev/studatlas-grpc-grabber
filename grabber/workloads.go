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

type WorkloadsMode string

const (
	Groups   WorkloadsMode = "Group"
	Teachers               = "Preps"
)

type WorkloadsInputs struct {
	View    string
	Faculty string
	Years   string
}

var inputs = &WorkloadsInputs{
	View:    op("View"),
	Faculty: cmb("Facultets"),
	Years:   cmb("Years"),
}

func (g *Grabber) GetGroupsWorkloads(academy *models.Academy, in *pb.ListGroupsWorkloadsRequest) ([]*pb.GroupWorkloadItem, error) {
	reqBody, err := json.Marshal(map[string]string{
		inputs.View:    "Group",
		inputs.Faculty: fmt.Sprint(in.FacultyId),
		inputs.Years:   in.Years,
	})

	req, err := http.NewRequest(http.MethodPost, academy.Endpoint, bytes.NewBuffer(reqBody))
	req.URL.Path = "/Nagr/Default.aspx"

	doc, err := g.GetDoc(req)

	if err != nil {
		log.Error(err)
	}

	var items []*pb.GroupWorkloadItem
	table := doc.Find("table#ContentPage_GridGroup")

	NewGrid(table, schemas.GroupWorkloadItem).EachRow(func(i int, row *Row) {
		items = append(items, &pb.GroupWorkloadItem{
			Year:       row.GetInt32("year"),
			Group:      row.Get("group"),
			GroupId:    row.GetSpecificId("groupId", "group"),
			Speciality: row.Get("speciality"),
			CountAll:   row.GetInt32("countAll"),
			Curriculum: row.Get("curriculum"),
		})
	})

	return items, nil
}

// GetTeachersWorkloads Учебная нагрузка преподавателей
func (g *Grabber) GetTeachersWorkloads(academy *models.Academy, in *pb.ListTeachersWorkloadsRequest) ([]*pb.TeacherWorkloadItem, error) {
	var items []*pb.TeacherWorkloadItem
	return items, nil
}
