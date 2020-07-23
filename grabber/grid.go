package grabber

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"grabber/schemas"
	"strings"
)

type Grid struct {
	ColumnsOrder map[string]int
	Nodes        *goquery.Selection
	Schema       map[string]*schemas.Property
}

func NewGrid(Nodes *goquery.Selection, Schema map[string]*schemas.Property) *Grid {
	columns := Nodes.Find(".TblHead > td,th").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})

	columnsOrder := make(map[string]int, len(columns))
	for k, v := range Schema {
		matched := false
		if v.Column != "" {
			for i, c := range columns {
				if strings.ToLower(v.Column) == strings.ToLower(c) {
					columnsOrder[k] = i
					matched = true
					break
				}
			}
		} else if len(v.Columns) != 0 {
			for _, t := range v.Columns {
				for i, c := range columns {
					if strings.ToLower(t) == strings.ToLower(c) {
						columnsOrder[k] = i
						matched = true
						break
					}
				}
				if matched {
					break
				}
			}
		}

		if !matched {
			log.Warnf("Column for the property \"%s\" hasn't been found", k)
		}
	}

	log.WithFields(log.Fields{
		"order": columnsOrder,
	}).Debugf("Schema properties")

	return &Grid{
		ColumnsOrder: columnsOrder,
		Nodes:        Nodes,
		Schema:       Schema,
	}
}

func (grid *Grid) EachRow(f func(int, *Row)) *goquery.Selection {
	return grid.Nodes.Find(".TblText, .TblhiText").Each(func(i int, s *goquery.Selection) {
		cols := s.Find("td")
		f(i, &Row{
			Grid:    grid,
			Node:    s,
			Columns: cols,
		})
	})
}
