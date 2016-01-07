package symbol

import "strings"

type MySQLConvert struct{}

func (mysql MySQLConvert) Convert(db_type string) string {
	kv := strings.SplitN(strings.ToLower(db_type), "(", 2)
	switch kv[0] {
	case "int", "tinyint":
		return "int64"
	case "date", "datetime", "timestamp":
		return "time.Time"
	case "float", "decimal", "double":
		return "float64"
	default:
		return "string"
	}
}
