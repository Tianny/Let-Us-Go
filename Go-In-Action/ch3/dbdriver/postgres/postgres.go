package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

type PostgresDriver struct {
}

func (dr PostgresDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("unimplemented")
}

var d *PostgresDriver

func init() {
	d = new(PostgresDriver)
	sql.Register("postgres", d)
}
