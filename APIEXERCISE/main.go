package main 
import (
	"fmt"
	"database/sql"
    _"github.com/go-sql-driver/mysql"
	"my_module/api"
	"log"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Connecting to MYSQL Database")
	db, err := sql.Open("mysql","root:kUqqo1-fyvwob@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		log.Fatal(err)
	}
    
	defer db.Close()
    fmt.Println("Successfully Connected to MYSQL Database ")

	if err := db.Ping() ; err != nil {
		log.Fatal(err)
	}

    fmt.Println("Initializing Kafka Producer...")
	producer, err := InitKafkaProducer()
	if err != nil {
	    log.Fatal("Error initializing Kafka producer: ", err)
	}
	defer producer.Close()


	r := gin.Default()
	api.RegisterRoutes(r,db,producer)
	fmt.Println("Server starting on port 8082...")
	log.Fatal(r.Run(":8082"))
}
	func InitKafkaProducer()(sarama.SyncProducer, error) {
		brokerlist := []string{"localhost:9092"}
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
	    config.Producer.Retry.Max = 5
	    config.Producer.Return.Successes = true

		producer, err := sarama.NewSyncProducer(brokerlist,config)
	    if err != nil {
		    return nil, err
	    }

	return producer, nil
} 



