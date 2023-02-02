package schema_helper

import (
	"fmt"
)

type PostgresType string

const (
	BIGINT     PostgresType = "BIGINT"
	TIME_STAMP PostgresType = "timestamp"
	JSON       PostgresType = "json"
	BOOLEAN    PostgresType = "BOOLEAN"
	TEXT       PostgresType = "text"
	DATE       PostgresType = "date"
	INTEGER    PostgresType = "integer"
)

func Varchar(size int) string {
	return fmt.Sprintf("varchar(%v)", size)
}

func Numeric(params ...int) string {
	switch len(params) {
	case 0:
		return "numeric()"
	case 1:
		return fmt.Sprintf("numeric(%v)", params[0])
	}
	precision := params[0]
	scale := params[1]
	return fmt.Sprintf("numeric(%v,%v)", precision, scale)
}
