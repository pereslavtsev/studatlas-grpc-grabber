package schemas

var Curricula = Schema{
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
