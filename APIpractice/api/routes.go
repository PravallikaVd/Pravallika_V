package api

import (
	"net/http"
	"database/sql"

)

func RegisterRoutes(db *sql.DB) {

	http.HandleFunc("/create", CreateHandler(db))
}