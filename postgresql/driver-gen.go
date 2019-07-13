// Code generated by gdbc/wrapper
// DO NOT EDIT!

package postgresql

import (
	"github.com/identitii/gdbc"
	"database/sql/driver"
	"database/sql"
)

var d = &Driver{}

func init() {
	RegisterDriverName("gdbc-postgresql")
}

func RegisterDriverName(name string) {
	sql.Register(name, d)
}

type Driver struct {
}

func Open(name string) (*Conn, error) {
	c, err := d.Open(name)
	if err != nil {
		return nil, err
	}
	return c.(*Conn), nil
}

func (j Driver) Open(name string) (driver.Conn, error) {

	user, password, jdbcURL, err := gdbc.ParseJDBCURL(name)
	if err != nil {
		return nil, err
	}

	jdbcConn, err := openJdbcConn(jdbcURL, user, password, gdbc.TRANSACTION_READ_COMMITTED) // TODO: Allow transaction isolation to be set. Somehow 🤷‍

	if err != nil {
		return nil, err
	}

	return &Conn{
		Conn: gdbc.NewConn(jdbcConn),
		c: jdbcConn,
	}, nil
}

type Conn struct { // A struct to allow individual drivers to extend functionality
	gdbc.Conn
	c *jdbcConn
}