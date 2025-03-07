package entities

type Notification struct {
	MessageId           int    `json:"message_id"`
	NotificationContent string `json:"notification_content"`
	ClientId            int    `json:"client_id"`
	ClientName          string `json:"client_name"`
}
