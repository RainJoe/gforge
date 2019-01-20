package schema

const (
	cUnderScore = "_"
)

var (
	typeWrappers = []typeWrapper{i64TypeWrapper, byteTypeWrapper, intTypeWrapper, float64TypeWrapper, float32TypeWrapper, stringTypeWrapper, timeTypeWrapper, jsonTypeWrapper}
)

// Column stands for a column of a table
type column struct {
	Name string `json:"column_name"`
	Type string `json:"data_type"`
}

//GetType returns which built in type the column should be in generated go code
func (c *column) GetType() (string, error) {
	t := getType(c.Type)
	if "" == t {
		return "", errUnknownType(c.Name, c.Type)
	}
	return t, nil
}

//GetName returns the Cammel Name of the struct
func (c *column) GetName() string {
	return convertUnderScoreToCammel(c.Name)
}

func getType(t string) string {
	for _, wrapper := range typeWrappers {
		typer := wrapper(t)
		if typer.Match() {
			return typer.Type()
		}
	}
	return ""
}
