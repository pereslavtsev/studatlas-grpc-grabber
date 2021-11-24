package schemas

var GroupWorkloadItem = Schema{
	"year": {
		Type:   "numeric",
		Column: "Курс",
	},
	"group": {
		Type:   "text",
		Column: "Группа",
	},
	"groupId": {
		Type:   "id",
		Column: "Группа",
	},
	"speciality": {
		Type:   "text",
		Column: "Специальность",
	},
	"countAll": {
		Type:   "numeric",
		Column: "Студентов",
	},
	"curriculum": {
		Type:   "text",
		Column: "Учебный план",
	},
}

var TeacherWorkloadItem = Schema{
	"teacherQuery": {
		Type:   "id",
		Column: "Преподаватель",
	},
	"teacher": {
		Type:   "text",
		Column: "Преподаватель",
	},
}
