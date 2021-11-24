package schemas

var Faculty = Schema{
	"id": {
		Type:    "id",
		Columns: []string{"Факультет", "Институт"},
	},
	"name": {
		Type:    "text",
		Columns: []string{"Факультет", "Институт"},
	},
	"abbreviation": {
		Type:   "text",
		Column: "Сокращение",
	},
	"head": {
		Type:   "text",
		Column: "Декан",
	},
	"phone": {
		Type:   "text",
		Column: "Телефон",
	},
	"room": {
		Type:   "text",
		Column: "Аудитория",
	},
}
