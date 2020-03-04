package databases

import (
	"database/sql"
	"fmt"
)

//Dbmodel ...
type Dbmodel struct {
	Database *sql.DB
}

//CreateDatabase ...
func CreateDatabase() (*sql.DB, error) {
	serverName := "db4free.net:3306"
	user := "lagash"
	password := ".a@vMztxya.hS6b"
	dbName := "marketplacelag"
	proveedor := "mysql"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open(proveedor, connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
