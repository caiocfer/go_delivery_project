package db

import (
	"database/sql"
	"fmt"

	"github.com/caiocfer/go_delivery_project/common"
	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	fmt.Println("Connecting to Postgresql")
	DBConnection := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		common.DB_HOST, common.DB_PORT, common.DB_USER, common.DB_PASSWORD, common.DB_NAME)

	DBMS := "postgres"
	db, err := sql.Open(DBMS, DBConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
