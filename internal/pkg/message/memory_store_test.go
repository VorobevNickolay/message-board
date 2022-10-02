package message

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetMessages(t *testing.T) {
	t.Run("should return empty list", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		actual, err := store.GetMessages(g)
		require.NoError(t, err)
		require.Equal(t, 0, len(actual))
	})

	t.Run("should return messages", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		expected1, err := store.CreateMessage(g, Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		expected2, err := store.CreateMessage(g, Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		actual, err := store.GetMessages(g)
		require.NoError(t, err)
		require.Equal(t, 2, len(actual))
		require.Equal(t, actual[0], &expected1)
		require.Equal(t, actual[1], &expected2)
	})
}
func TestFindMessageById(t *testing.T) {
	t.Run("should return ErrMessageNotFound", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		actual, err := store.FindMessageByID(g, uuid.NewString())
		require.Error(t, err, ErrMessageNotFound)
		require.Equal(t, actual, Message{})
	})

	t.Run("should find message", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		_, err := store.CreateMessage(g, Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		expected, err := store.CreateMessage(g, Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		_, err = store.CreateMessage(g, Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		})
		require.NoError(t, err)

		actual, err := store.FindMessageByID(g, expected.ID)
		require.Equal(t, expected, actual)
	})
}
func TestCreateMessage(t *testing.T) {
	t.Run("should create message", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		expected := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		actual, err := store.CreateMessage(g, expected)
		require.NoError(t, err)
		require.Equal(t, expected.UserID, actual.UserID)
		require.Equal(t, expected.Text, actual.Text)
		require.NotEmpty(t, actual.ID)
	})
}

func TestUpdateMessage(t *testing.T) {
	t.Run("should update message", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		expected := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		expected, _ = store.CreateMessage(g, expected)

		updated := Message{
			ID:     expected.ID,
			UserID: expected.UserID,
			Text:   "Hi!",
		}
		actual, err := store.UpdateMessage(g, updated)

		require.NoError(t, err)
		require.Equal(t, expected.UserID, actual.UserID)
		require.Equal(t, expected.ID, actual.ID)
		require.Equal(t, "Hi!", actual.Text)
		require.NotEmpty(t, actual.ID)
	})
	t.Run("should return errMessageNotFound", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		expected := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		updated := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		expected, _ = store.CreateMessage(g, expected)
		actual, err := store.UpdateMessage(g, updated)
		require.Error(t, err, ErrMessageNotFound)
		require.Empty(t, actual)
	})
}

func TestDeleteMessage(t *testing.T) {
	t.Run("should delete message", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		oldMessage := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		oldMessage, _ = store.CreateMessage(g, oldMessage)
		err := store.DeleteMessage(g, oldMessage.ID)
		deleted, err2 := store.FindMessageByID(g, oldMessage.ID)
		require.NoError(t, err)
		require.Error(t, err2)
		require.Empty(t, deleted)
	})
	t.Run("should return ErrMessageNotFound", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		oldMessage := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		oldMessage, _ = store.CreateMessage(g, oldMessage)
		err := store.DeleteMessage(g, uuid.NewString())
		deleted, _ := store.FindMessageByID(g, oldMessage.ID)
		require.Error(t, err, ErrMessageNotFound)
		require.Equal(t, deleted, oldMessage)
	})
}
