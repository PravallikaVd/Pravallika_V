package main
import (
	"fmt"
	"database/sql"
	"log"
	"my_module/api"
	"net/http"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connecting to MYSQL Database ")
    db, err := sql.Open("mysql","root:kUqqo1-fyvwob@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Successfully Connected to MYSQL Database ")

	if err := db.Ping() ; err != nil {
		log.Fatal(err)

	}


	api.RegisterRoutes(db)
	fmt.Println("Server starting on port 8082...")
	log.Fatal(http.ListenAndServe(":8082",nil))
}
