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

func (row *Row) GetInt32(key string) int32 {
	i, _ := strconv.ParseInt(row.Get(key), 10, 32)
	return int32(i)
}

func (row *Row) GetSpecificId(key string, param string) int32 {
	cell := row.Columns.Eq(row.Grid.ColumnsOrder[key])
	val, exists := cell.Find("a").Attr("href")
	if !exists {
		//log.Warnf(`Link hasn't founded, used text value "%s" instead`, cell.Text())
		i, _ := strconv.ParseInt(cell.Text(), 10, 32)
		return int32(i)
	}
	u, err := url.Parse(val)
	if err != nil || u == nil {
		log.Error()
		return -1
	}

	if u.Query().Get("group") == "" {
		log.Warnf(`No ID in the url: %s`, u.RequestURI())
	}

	i, _ := strconv.ParseInt(u.Query().Get("group"), 10, 32)
	return int32(i)
}

func (row *Row) Get(key string) string {
	return strings.TrimSpace(row.Columns.Eq(row.Grid.ColumnsOrder[key]).Text())
}
