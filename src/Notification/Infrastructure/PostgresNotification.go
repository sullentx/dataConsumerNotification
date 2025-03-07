package infraestructure

import (
	entities "api-consumer/Notification/Domain/Entities"
	"database/sql"
)

type PostgresNotification struct {
	db *sql.DB
}

func NewPostgresNotificationRepository(db *sql.DB) *PostgresNotification {
	return &PostgresNotification{db: db}
}

func (r *PostgresNotification) Save(notification entities.Notification) error {
	_, err := r.db.Exec(
		"INSERT INTO notifications (notification_content, client_id, client_name) VALUES ($1, $2, $3)",
		notification.NotificationContent,
		notification.ClientId,
		notification.ClientName,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresNotification) GetMessages() ([]entities.Notification, error) {
	rows, err := r.db.Query("SELECT message_id, notification_content, client_id, client_name FROM notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []entities.Notification
	for rows.Next() {
		var notification entities.Notification
		var clientName string
		if err := rows.Scan(&notification.MessageId, &notification.NotificationContent, &notification.ClientId, &clientName); err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}
