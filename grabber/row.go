package grabber

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

type Row struct {
	Columns *goquery.Selection
	Grid    *Grid
	Node    *goquery.Selection
}

func (row *Row) GetId(key string) int32 {
	return row.GetSpecificId(key, "id")
}

func (row *Row) GetSpecificId(key string, param string) int32 {
	cell := row.Columns.Eq(row.Grid.ColumnsOrder[key])
	val, exists := cell.Find("a").Attr("href")
	if !exists {

	}
	u, err := url.ParseQuery(val)
	if err != nil || u == nil {
		log.Error()
		return -1
	}

	i, _ := strconv.ParseInt(u.Get(param), 10, 32)
	return int32(i)
}

func (row *Row) Get(key string) string {
	return strings.TrimSpace(row.Columns.Eq(row.Grid.ColumnsOrder[key]).Text())
}