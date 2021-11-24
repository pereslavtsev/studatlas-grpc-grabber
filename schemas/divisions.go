package schemas

var Division = Schema{
	"id": {
		Type:   "id",
		Column: "Номер",
	},
	"name": {
		Type:   "text",
		Column: "Название",
	},
	"abbreviation": {
		Type:   "text",
		Column: "Сокращение",
	},
	"head": {
		Type:   "text",
		Column: "ЗавКафедрой",
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
