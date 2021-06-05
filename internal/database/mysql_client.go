package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julian1le9cuero/course-phones-review/internal/logs"
)

type MySqlClient struct {
	*sql.DB
}

// Return db client
func NewSqlClient(source string) *MySqlClient {
	// Metodo open para abrir conexion
	// driverName = "mysql"
	// dataSourceName = source argument
	db, err := sql.Open("mysql", source)

	if err != nil {
		logs.Log().Errorf("cannot create db: %s", err.Error())
		// Panic stops the execution
		panic(err)
	}
	// Ping verifies the connection
	err = db.Ping()

	if err != nil {
		logs.Log().Warn("cannot connect to mysql")
	}

	return &MySqlClient{db}
}

func (c *MySqlClient) ViewStats() sql.DBStats {
	return c.Stats()
}
