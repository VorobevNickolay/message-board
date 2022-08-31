package message

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetMessages(t *testing.T) {
	t.Run("should return empty list", func(t *testing.T) {
		store := NewInMemoryStore()

		actual, err := store.GetMessages()
		require.NoError(t, err)
		require.Equal(t, 0, len(actual))
	})

	t.Run("should return messages", func(t *testing.T) {
		store := NewInMemoryStore()

		expected1, err := store.CreateMessage(Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		expected2, err := store.CreateMessage(Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		actual, err := store.GetMessages()
		require.NoError(t, err)
		require.Equal(t, 2, len(actual))
		require.Equal(t, actual[0], &expected1)
		require.Equal(t, actual[1], &expected2)
	})
}
func TestFindMessageById(t *testing.T) {
	t.Run("should return ErrMessageNotFound", func(t *testing.T) {
		store := NewInMemoryStore()
		actual, err := store.FindMessageById(uuid.NewString())
		require.Error(t, err, ErrMessageNotFound)
		require.Equal(t, actual, Message{})
	})

	t.Run("should find message", func(t *testing.T) {
		store := NewInMemoryStore()
		_, err := store.CreateMessage(Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		expected, err := store.CreateMessage(Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		_, err = store.CreateMessage(Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		actual, err := store.FindMessageById(expected.ID)
		require.Equal(t, expected, actual)
	})
}
func TestCreateMessage(t *testing.T) {
	t.Run("should create message", func(t *testing.T) {
		store := NewInMemoryStore()

		expected := Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		actual, err := store.CreateMessage(expected)
		require.NoError(t, err)
		require.Equal(t, expected.UserId, actual.UserId)
		require.Equal(t, expected.Text, actual.Text)
		require.NotEmpty(t, actual.ID)
	})
	t.Run("should return ErrEmptyMessage", func(t *testing.T) {
		store := NewInMemoryStore()

		expected := Message{}
		actual, err := store.CreateMessage(expected)
		require.Error(t, err, ErrEmptyMessage)
		require.Empty(t, actual)
	})
}

func TestUpdateMessage(t *testing.T) {
	t.Run("should update message", func(t *testing.T) {
		store := NewInMemoryStore()

		expected := Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		expected, _ = store.CreateMessage(expected)
		actual, err := store.UpdateMessage(expected.ID, "Hi!")
		require.NoError(t, err)
		require.Equal(t, expected.UserId, actual.UserId)
		require.Equal(t, expected.ID, actual.ID)
		require.Equal(t, "Hi!", actual.Text)
		require.NotEmpty(t, actual.ID)
	})
	t.Run("should return ErrEmptyMessage", func(t *testing.T) {
		store := NewInMemoryStore()

		expected := Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		expected, _ = store.CreateMessage(expected)
		actual, err := store.UpdateMessage(expected.ID, "")
		require.Error(t, err, ErrEmptyMessage)
		require.Empty(t, actual)
	})
	t.Run("should return errMessageNotFound", func(t *testing.T) {
		store := NewInMemoryStore()

		expected := Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		expected, _ = store.CreateMessage(expected)
		actual, err := store.UpdateMessage(uuid.NewString(), "Hi!")
		require.Error(t, err, ErrMessageNotFound)
		require.Empty(t, actual)
	})
}

func TestDeleteMessage(t *testing.T) {
	t.Run("should delete message", func(t *testing.T) {
		store := NewInMemoryStore()

		oldMessage := Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		oldMessage, _ = store.CreateMessage(oldMessage)
		err := store.DeleteMessage(oldMessage.ID)
		deleted := store.messages[store.messageIDs[oldMessage.ID]]
		require.NoError(t, err)
		require.Empty(t, deleted)
	})
	t.Run("should return ErrMessageNotFound", func(t *testing.T) {
		store := NewInMemoryStore()

		oldMessage := Message{
			UserId: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		oldMessage, _ = store.CreateMessage(oldMessage)
		err := store.DeleteMessage(uuid.NewString())
		deleted := store.messages[store.messageIDs[oldMessage.ID]]
		require.Error(t, err, ErrMessageNotFound)
		require.Equal(t, *deleted, oldMessage)
	})
}
