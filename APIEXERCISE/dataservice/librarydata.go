package dataservice
import (
	"database/sql"
	"my_module/model"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/IBM/sarama"
	"fmt"
	"encoding/json"

	
)

func SubscribeUser(db *sql.DB, c *gin.Context) error {
	var subscription model.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		return err
	}
    
	topicsCSV := strings.Join(subscription.Topics, ",")

	query := `INSERT INTO subscriptions(user_id, topics, email, sms, push_notifications) VALUES(?,?,?,?,?)`
	_, err := db.Exec(query,subscription.UserID, topicsCSV, subscription.NotificationChannels.Email, subscription.NotificationChannels.SMS, subscription.NotificationChannels.PushNotifications)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, gin.H{"userdata": subscription})
    return nil
}

func UnsubscribeUser(db *sql.DB, c *gin.Context) error {
	var unsubscribeRequest model.UnsubscribeRequest

	if err := c.ShouldBindJSON(&unsubscribeRequest); err != nil {
		return err
	}

	var currentTopics string
	query := `SELECT topics FROM subscriptions WHERE user_id = ?`
	err := db.QueryRow(query, unsubscribeRequest.UserID).Scan(&currentTopics)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "No subscriptions found for user"})
			return nil
		}
		return err
	}

	topicsList := strings.Split(currentTopics, ",")

	var remainingTopics []string
	for _, topic := range topicsList {

		if !contains(unsubscribeRequest.Topics, topic) {
			remainingTopics = append(remainingTopics, topic)
		}
	}

	if len(remainingTopics) == 0 {
		deleteQuery := `DELETE FROM subscriptions WHERE user_id = ?`
		_, err := db.Exec(deleteQuery, unsubscribeRequest.UserID)
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, gin.H{"message": "Successfully unsubscribed from all topics"})
		return nil
	}

	updatedTopics := strings.Join(remainingTopics, ",")
	updateQuery := `UPDATE subscriptions SET topics = ? WHERE user_id = ?`
	_, err = db.Exec(updateQuery, updatedTopics, unsubscribeRequest.UserID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"user_id":unsubscribeRequest.UserID , "remaining_topics": remainingTopics})
	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}


func FetchUser(db *sql.DB, producer sarama.SyncProducer, c *gin.Context) error {
   userID := c.Param("user_id")
   query := `SELECT topics, email, sms, push_notifications FROM subscriptions WHERE user_id = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var subscriptions []model.SubscriptionResponse
	for rows.Next() {
		var sub model.SubscriptionResponse
		err := rows.Scan(&sub.Topic, &sub.Channels.Email, &sub.Channels.SMS, &sub.Channels.PushNotifications)
		if err != nil {
			return err
		}
		subscriptions = append(subscriptions, sub)
	}

	message := &sarama.ProducerMessage{
		Topic: "user_subscription_fetch",
		Value: sarama.StringEncoder(fmt.Sprintf("User %s fetched their subscriptions", userID)),
	}
	_, _, err = producer.SendMessage(message)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "subscriptions": subscriptions})
	return nil
}


func SendNotification(db *sql.DB, producer sarama.SyncProducer, c *gin.Context) error {
	
	var payload model.Notification 

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return err
	}

	if err := storeNotificationInDB(db, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store notification in the database"})
		return err
	}

	topic := payload.Topic
	switch topic {
	case "signup":
		topic = "signup_notifications"
	case "purchase":
		topic = "purchase_notifications"
	case "sms":
		topic = "sms_notifications"
	}

	messageContent, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize notification payload"})
		return err
	}

	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(messageContent),
	}

	_, _, err = producer.SendMessage(kafkaMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to Kafka"})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Notification sent successfully to %s", topic)})
	return nil 

}
   
func storeNotificationInDB(db *sql.DB, payload *model.Notification) error {


	query := `INSERT INTO notifications (topic, event_id, message_title, message_body, user_id) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(query, payload.Topic, payload.Event.EventID,  payload.Message.Title, payload.Message.Body, payload.Event.Details["user_id"])
	if err != nil {
		return err
	}
	return nil
}	