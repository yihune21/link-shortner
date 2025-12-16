package utils

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/yihune21/link-shortner/internal/database"
)

func ConnectDb(db_url string) (*database.Queries, error ) {
	conn , err := sql.Open("postgres",db_url)
	if err != nil {
		return nil, errors.New("Can't connect")
	}
	
	db_conn := database.New(conn)
    return db_conn , nil
}