package message

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/suite"
	"testing"
)

type postgresStoreTestSuite struct {
	suite.Suite
	ctx context.Context

	pool  *pgxpool.Pool
	store store
}

func TestPostgresStore(t *testing.T) {
	suite.Run(t, new(postgresStoreTestSuite))
}
func (suite *postgresStoreTestSuite) TestPostgresStore_CreateMessage() {
	suite.Run("should create user", func() {
		m := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}

		m1, err := suite.store.CreateMessage(suite.ctx, m)
		suite.Require().NoError(err)
		suite.NotEmpty(m1.ID)
		suite.Equal(m.UserID, m1.UserID)
		suite.Equal(m.Text, m1.Text)
	})
	suite.Run("should return errEmptyMessage", func() {
		m := Message{
			UserID: uuid.NewString(),
			Text:   "",
		}

		m, err := suite.store.CreateMessage(suite.ctx, m)
		suite.Require().EqualError(err, ErrEmptyMessage.Error())
		suite.Empty(m)
	})
}
func (suite *postgresStoreTestSuite) TestPostgresStore_FindMessageById() {
	suite.Run("should return message", func() {
		m := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		m1 := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		m, err := suite.store.CreateMessage(suite.ctx, m)
		suite.Require().NoError(err)
		suite.Require().NotEmpty(m)
		m1, err = suite.store.CreateMessage(suite.ctx, m1)
		suite.Require().NoError(err)
		suite.Require().NotEmpty(m1)
		actualMessage, err := suite.store.FindMessageByID(suite.ctx, m.ID)
		suite.Require().NoError(err)
		m.CreatedAt = actualMessage.CreatedAt
		suite.Equal(m, actualMessage)
	})
	suite.Run("should return errMessageNotFound", func() {
		m := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}

		m, err := suite.store.CreateMessage(suite.ctx, m)
		suite.Require().NoError(err)
		suite.Require().NotEmpty(m)

		actualMessage, err := suite.store.FindMessageByID(suite.ctx, uuid.NewString())
		suite.Require().EqualError(err, ErrMessageNotFound.Error())
		suite.Empty(actualMessage)
	})
}

func (suite *postgresStoreTestSuite) TestPostgresStore_GetMessages() {
	suite.Run("should return empty array", func() {
		messages, err := suite.store.GetMessages(suite.ctx)
		suite.Require().NoError(err)
		suite.Empty(messages)
	})
	suite.Run("should return messages", func() {
		m := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		m1 := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		m, err := suite.store.CreateMessage(suite.ctx, m)
		suite.Require().NoError(err)
		suite.Require().NotEmpty(m)
		m1, err = suite.store.CreateMessage(suite.ctx, m1)
		suite.Require().NoError(err)
		suite.Require().NotEmpty(m1)
		messages, err := suite.store.GetMessages(suite.ctx)
		suite.Require().NoError(err)
		m.CreatedAt = messages[0].CreatedAt
		m1.CreatedAt = messages[1].CreatedAt
		suite.Equal(m, *messages[0])
		suite.Equal(m1, *messages[1])
	})
}

func (suite *postgresStoreTestSuite) TestPostgresStore_UpdateMessage() {
	suite.Run("should return message", func() {
		m := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}

		m, err := suite.store.CreateMessage(suite.ctx, m)
		updated := Message{
			ID:     m.ID,
			UserID: m.UserID,
			Text:   uuid.NewString(),
		}

		suite.Require().NoError(err)
		suite.Require().NotEmpty(m)
		actualMessage, err := suite.store.UpdateMessage(suite.ctx, updated)
		suite.Require().NoError(err)
		suite.Equal(m.ID, actualMessage.ID)
		suite.Equal(m.UserID, actualMessage.UserID)
		suite.Equal(updated.Text, actualMessage.Text)
	})
}

func (suite *postgresStoreTestSuite) TestPostgresStore_DeleteMessage() {
	suite.Run("should return message", func() {
		m := Message{
			UserID: uuid.NewString(),
			Text:   uuid.NewString(),
		}
		m, err := suite.store.CreateMessage(suite.ctx, m)
		suite.Require().NoError(err)
		suite.Require().NotEmpty(m)
		err = suite.store.DeleteMessage(suite.ctx, m.ID)
		suite.Require().NoError(err)
		users, err := suite.store.GetMessages(suite.ctx)
		suite.Require().NoError(err)
		suite.Empty(users)
	})
	suite.Run("should return errMessageNotFound", func() {
		err := suite.store.DeleteMessage(suite.ctx, uuid.NewString())
		suite.Require().Error(err, ErrMessageNotFound)
		users, err := suite.store.GetMessages(suite.ctx)
		suite.Require().NoError(err)
		suite.Empty(users)
	})
}
func (suite *postgresStoreTestSuite) SetupTest() {
	suite.ctx = context.Background()
	databaseUrl := "postgres://vorobevna:message-board@localhost:15432/postgres"

	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	suite.Require().NoError(err)
	suite.pool = dbPool
	suite.store = NewPostgresStore(dbPool)
}
func (suite *postgresStoreTestSuite) AfterTest(_, _ string) {
	suite.truncateTable("messages")
}
func (suite *postgresStoreTestSuite) truncateTable(tableName string) {
	_, err := suite.pool.Exec(context.Background(), "TRUNCATE TABLE"+" "+tableName+";")
	suite.Require().NoError(err)
}
