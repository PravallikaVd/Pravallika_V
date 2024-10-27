package api
import (
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	"my_module/dataservice"
	"github.com/IBM/sarama"

)

func SubscribeHandler(db *sql.DB) gin.HandlerFunc {
    return func (c *gin.Context) {
        if err := dataservice.SubscribeUser(db, c); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Subscribed successfully"})
	}


}


func UnsubscribeHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := dataservice.UnsubscribeUser(db,c) ; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
          c.JSON(http.StatusOK, gin.H{"message": "Unsubscribed successfully from selected topics"})
	}

}

func FetchHandler(db *sql.DB, producer sarama.SyncProducer) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := dataservice.FetchUser(db,producer,c) ; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

}

func NotificationHandler(db *sql.DB, producer sarama.SyncProducer) gin.HandlerFunc {
	return func( c *gin.Context) {
		if err := dataservice.SendNotification(db,producer, c); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully"})
	}
}







