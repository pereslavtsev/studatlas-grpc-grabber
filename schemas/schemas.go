package schemas

type Property struct {
	Type    string
	Column  string
	Columns []string
}

type Schema map[string]*Property
