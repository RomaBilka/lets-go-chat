package repositories

import (
	"database/sql"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
)

func NewPostgreMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

type MessageRepository struct {
	db *sql.DB
}

func (r *MessageRepository) GetMessages() ([]models.Message, error) {
	messages := []models.Message{}

	rows, err := r.db.Query("SELECT * FROM messages")

	if err != nil && err != sql.ErrNoRows {
		return messages, err
	}

	for rows.Next() {
		message := models.Message{}
		if err := rows.Scan(&message.Id, &message.UserId, &message.Message, &message.CreatedAt); err != nil && err != sql.ErrNoRows {
			return []models.Message{}, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (r *MessageRepository) CreateMessage(message models.Message) (models.MessageId, error) {
	id := 0
	err := r.db.QueryRow("INSERT INTO messages (user_id, message) VALUES ($1, $2)  RETURNING id", message.UserId, message.Message).Scan(&id)
	if err != nil {
		return 0, err
	}

	return models.MessageId(id), nil
}
