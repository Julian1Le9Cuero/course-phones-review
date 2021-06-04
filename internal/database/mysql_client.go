package database

import (
	"database/sql"
	// "github.com/julian1le9cuero"
)

type MySqlClient struct {
	*sql.DB
}

// Return db client
func NewSqlClient(source string) *sql.DB {
	// Metodo open para abrir conexion
	// driverName = "mysql"
	// dataSourceName = source argument
	db, err := sql.Open("mysql", source)

	if err != nil {
		logs.log().Errorf("cannot create db: %s", err.Error())
		// Panic stops the execution
		panic(err)
	}
	// Ping verifies the connection
	err = db.Ping()

	if err != nil {
		logs.log().Warn("cannot connect to mysql")
	}

	return &MySqlClient{db}
}

func (c *MySqlClient) ViewStats() sql.DBStats {
	return c.Stats()
}
