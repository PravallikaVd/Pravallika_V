package model

type Subscription struct {
	UserID              string               `json:"user_id"`
	Topics              []string             `json:"topics"`
	NotificationChannels NotificationChannels `json:"notification_channels"`
}
type NotificationChannels struct {
	Email             string `json:"email"`
	SMS               string `json:"sms"`
	PushNotifications bool   `json:"push_notifications"`
}
type Notification struct {
	Topic   string  `json:"topic"`
	Event   Event   `json:"event"`
	Message Message `json:"message"`
}
type Event struct {
	EventID   string                 `json:"event_id"`
	Timestamp string                 `json:"timestamp"`
	Details   map[string]interface{} `json:"details"`
}

type UnsubscribeRequest struct {
	UserID              string               `json:"user_id"`
	Topics              []string             `json:"topics"`

}

type Message struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}


type SubscriptionResponse struct {
    Topic    string
    Channels struct {
        Email            string
        SMS              string
        PushNotifications bool
    }
}
