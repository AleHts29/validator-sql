package validator_sql

import "fmt"

var version = "1.0.0"
var library = "validator-sql"

func Version() string {
	v := fmt.Sprintf("%s_v%s", library, version)
	return v
}
