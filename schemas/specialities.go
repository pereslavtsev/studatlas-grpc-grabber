package schemas

var Speciality = Schema{
	"id": {
		Type:   "id",
		Column: "Специальность",
	},
	"shortName": {
		Type:   "text",
		Column: "Специальность",
	},
	"name": {
		Type:   "text",
		Column: "Название Специальности",
	},
	"facultyId": {
		Type:   "id",
		Column: "Факультет",
	},
	"faculty": {
		Type:   "text",
		Column: "Факультет",
	},
	"divisionId": {
		Type:   "id",
		Column: "Кафедра",
	},
	"division": {
		Type:   "text",
		Column: "Кафедра",
	},
	"code": {
		Type:   "text",
		Column: "Шифр",
	},
	"qualification": {
		Type:   "text",
		Column: "Квалификация",
	},
}
