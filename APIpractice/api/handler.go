package api

import (
 
	"net/http"
	"database/sql"
    

)

func CreateHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
    
	if err := CreateStudentsLogic(db,w,r) ; err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
	   return 

	}
}
}