package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/stretchr/testify/assert"
)

func createMessage(t *testing.T) {
	userId := createUser(t)

	id, err := testMessageRepository.CreateMessage(models.Message{UserId: userId, Message: message.Message})

	assert.NotNil(t, id)
	assert.NoError(t, err)
}

func TestCreatereateMessage(t *testing.T) {
	createMessage(t)
	err := truncateUsers()
	assert.NoError(t, err)
}

func TestGetMessages(t *testing.T) {
	createMessage(t)

	messages, err := testMessageRepository.GetMessages()
	assert.NoError(t, err)
	assert.Equal(t, message.Message, messages[0].Message)

	err = truncateUsers()
	assert.NoError(t, err)
}

func TestGetMessagesV2(t *testing.T) {
	db, mock := NewMock()
	repo := NewPostgreMessageRepository(db)

	query := regexp.QuoteMeta("SELECT * FROM messages")
	rows := sqlmock.NewRows([]string{"id", "user_id", "message", "created_at"}).AddRow(1, message.UserId, message.Message, time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	messages, err := repo.GetMessages()

	assert.NotNil(t, messages)
	assert.NoError(t, err)
}
