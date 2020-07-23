package schemas

var Group = Schema{
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
