package models

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq" //The database driver

	"fmt"
)

type goBlogDb struct {
	db *sql.DB
}

type dbConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	HostIP   string `json:"host_ip"`
	Port     string `json:"port"`
	Sslmode  string `json:"ssl_mode"`
	Dbname   string `json:"dbname"`
}

func getDBConnection(config dbConfig) (*goBlogDb, error) {
	connectionString := makeDBConnectionString(config)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.New("unable to connect to database: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.New("unable to verify connection to database: " + err.Error())
	}

	newGoBlogDB := goBlogDb{
		db: db,
	}

	return &newGoBlogDB, nil
}

func makeDBConnectionString(config dbConfig) string {
	return fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v",
		config.Username,
		config.Password,
		config.HostIP,
		config.Port,
		config.Dbname,
		config.Sslmode,
	)
}

func unWrapStructTags(unwrapStuct interface{}) {

}
