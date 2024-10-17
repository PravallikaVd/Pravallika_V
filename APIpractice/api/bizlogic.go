package api
import (
	"net/http"
	"my_module/dataservice"
	"database/sql"

)

func CreateStudentsLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.CreateStudent(db,w,r)
}