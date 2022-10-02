package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUsers(t *testing.T) {
	t.Run("should return empty list", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}
		actual, err := store.GetUsers(g)
		require.NoError(t, err)
		require.Equal(t, 0, len(actual))
	})

	t.Run("should return users", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}
		exp1, err := store.CreateUser(g, uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		exp2, err := store.CreateUser(g, uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		actual, err := store.GetUsers(g)

		require.Equal(t, 2, len(actual))
		if exp1 == *actual[0] {
			require.Equal(t, exp1, *actual[0])
			require.Equal(t, exp2, *actual[1])
		} else {
			require.Equal(t, exp1, *actual[1])
			require.Equal(t, exp2, *actual[0])
		}
	})
}

func TestFindUserById(t *testing.T) {
	t.Run("should return ErrUserNotFound", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		actual, err := store.FindUserByID(g, uuid.NewString())
		require.Error(t, err, ErrUserNotFound)
		require.Equal(t, actual, User{})
	})

	t.Run("should find user", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		_, err := store.CreateUser(g, uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		expected, err := store.CreateUser(g, uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		_, err = store.CreateUser(g, uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		actual, err := store.FindUserByID(g, expected.ID)
		require.Equal(t, expected, actual)
	})
}
func TestFindUserByName(t *testing.T) {
	t.Run("should return ErrUserNotFound", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		actual, err := store.findUserByName(g, uuid.NewString())
		require.Error(t, err, ErrUserNotFound)
		require.Equal(t, actual, User{})
	})

	t.Run("should find user", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		_, err := store.CreateUser(g, uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		expected, err := store.CreateUser(g, uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		_, err = store.CreateUser(g, uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		actual, err := store.findUserByName(g, expected.Username)
		require.Equal(t, expected, actual)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("should create user", func(t *testing.T) {
		store := NewInMemoryStore()
		g := &gin.Context{}

		expected := User{
			Username: uuid.NewString(),
			ID:       uuid.NewString(),
			Password: uuid.NewString(),
		}
		actual, err := store.CreateUser(g, expected.Username, expected.Password)
		require.NoError(t, err)
		require.Equal(t, expected.Username, actual.Username)
		require.Equal(t, expected.Password, actual.Password)
		require.NotEmpty(t, actual.ID)
	})
}
