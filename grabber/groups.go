package grabber

import (
	"grabber/models"

	log "github.com/sirupsen/logrus"
	pb "grabber/pb"
)

var schemaGroup = map[string]*Property{
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
	"specialityId": {
		Type:   "id",
		Column: "Специальность",
	},
	"speciality": {
		Type:   "speciality",
		Column: "Специальность",
	},
	"countAll": {
		Type:   "numeric",
		Column: "Всего",
	},
	"countCommon": {
		Type:   "numeric",
		Column: "ОО",
	},
	"countTargeted": {
		Type:   "numeric",
		Column: "ЦН",
	},
	"countSpecial": {
		Type:   "numeric",
		Column: "СН",
	},
	"curricula": {
		Type:   "text",
		Column: "Учебный План",
	},
}

func (g *Grabber) FetchGroups(academy *models.Academy, params *DictionaryFilter) ([]*pb.Group, error) {
	doc, err := g.DoDictionaryReq(academy, params)

	if err != nil {
		log.Error(err)
	}

	var groups []*pb.Group

	table := doc.Find("table#ContentPage_ucGroups_Grid")

	NewGrid(table, schemaGroup).EachRow(func(i int, row *Row) {
		groups = append(groups, &pb.Group{
			Id:            row.GetId("id"),
			Name:          row.Get("name"),
			Year:          row.GetInt32("year"),
			CountAll:      row.GetInt32("countAll"),
			CountCommon:   row.GetInt32("countCommon"),
			CountTargeted: row.GetInt32("countTargeted"),
			CountSpecial:  row.GetInt32("countSpecial"),
			Curricula:     row.Get("curricula"),
			SpecialityId:  row.GetInt32("specialityId"),
			Speciality:    row.Get("speciality"),
		})
	})

	return groups, nil
}

func (g *Grabber) GetGroups(academy *models.Academy) ([]*pb.Group, error) {
	groups, err := g.FetchGroups(academy, &DictionaryFilter{
		Mode:   "group",
		Filter: "group",
	})
	return groups, err
}

func (g *Grabber) GetFacultyGroups(academy *models.Academy, facultyId int32) ([]*pb.Group, error) {
	groups, err := g.FetchGroups(academy, &DictionaryFilter{
		Mode:   "group",
		Filter: "facultet",
		ID:     facultyId,
	})
	return groups, err
}

func (g *Grabber) GetSpecialityGroups(academy *models.Academy, specialityId int32) ([]*pb.Group, error) {
	groups, err := g.FetchGroups(academy, &DictionaryFilter{
		Mode:   "group",
		Filter: "spets",
		ID:     specialityId,
	})
	return groups, err
}

func (g *Grabber) GetGroupById(academy *models.Academy, id int32) (*pb.Group, error) {
	groups, err := g.FetchGroups(academy, &DictionaryFilter{
		Mode:   "group",
		Filter: "spets",
		ID:     id,
	})
	return groups[0], err
}
