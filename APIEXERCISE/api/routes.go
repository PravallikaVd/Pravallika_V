package api

import (
	"database/sql"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB, producer sarama.SyncProducer)  {

	router.POST("/subscribe", SubscribeHandler(db))
	router.POST("/unsubscribe", UnsubscribeHandler(db))
	router.GET("/subscriptions/:user_id", FetchHandler(db,producer))
	router.POST("/notifications/send", NotificationHandler(db,producer))


}