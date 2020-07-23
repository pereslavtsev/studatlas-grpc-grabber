package schemas

var Report = Schema{
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
	"speciality": {
		Type:   "speciality",
		Column: "Специальность",
	},
	"count": {
		Type:   "numeric",
		Column: "Студентов",
	},
	"curricula": {
		Type:   "text",
		Column: "Учебный План",
	},
}
