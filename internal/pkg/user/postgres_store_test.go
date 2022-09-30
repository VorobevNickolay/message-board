package user

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
	store Store
}

func TestPostgresStore(t *testing.T) {
	suite.Run(t, new(postgresStoreTestSuite))
}
func (suite *postgresStoreTestSuite) TestPostgresStore_CreateUser() {
	suite.Run("should create user", func() {
		name := uuid.NewString()
		password := uuid.NewString()

		user, err := suite.store.CreateUser(suite.ctx, name, password)
		suite.Require().NoError(err)
		suite.NotEmpty(user.ID)
		suite.Equal(name, user.Username)
	})
	suite.Run("should return errUsedUsername", func() {
		name := uuid.NewString()
		password := uuid.NewString()
		user, err := suite.store.CreateUser(suite.ctx, name, password)
		suite.Require().NoError(err)
		suite.Require().NotEmpty(user)
		user, err = suite.store.CreateUser(suite.ctx, name, password)
		suite.Require().Error(err, ErrUsedUsername)
		suite.Empty(user)
	})
	suite.Run("should return errEmptyPassword", func() {
		name := ""
		password := ""
		user, err := suite.store.CreateUser(suite.ctx, name, password)
		suite.Require().Error(err, ErrEmptyPassword)
		suite.Empty(user)
	})
}
func (suite *postgresStoreTestSuite) TestPostgresStore_GetUsers() {
	suite.Run("should return empty users", func() {
		users, err := suite.store.GetUsers(suite.ctx)
		suite.Require().NoError(err)
		suite.Empty(users)
	})
	suite.Run("should return users", func() {

		user1, _ := suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		user2, _ := suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		users, err := suite.store.GetUsers(suite.ctx)
		suite.Require().NoError(err)
		suite.Equal(user1, *users[0])
		suite.Equal(user2, *users[1])
	})
}
func (suite *postgresStoreTestSuite) TestPostgresStore_FindUserById() {
	suite.Run("should return user", func() {
		name := uuid.NewString()
		password := uuid.NewString()

		user, err := suite.store.CreateUser(suite.ctx, name, password)
		suite.Require().NoError(err)
		suite.Require().NotEmpty(user)
		u, err := suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		suite.Require().NoError(err)
		suite.Require().NotEmpty(u)
		actualUser, err := suite.store.FindUserByID(suite.ctx, user.ID)
		suite.Require().NoError(err)
		suite.Equal(user, actualUser)
	})
	suite.Run("should return error", func() {
		u, err := suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		suite.Require().NoError(err)
		suite.Require().NotEmpty(u)
		u, err = suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		suite.Require().NoError(err)
		suite.Require().NotEmpty(u)
		actualUser, err := suite.store.FindUserByID(suite.ctx, uuid.NewString())

		suite.Require().EqualError(err, ErrUserNotFound.Error())
		suite.Empty(actualUser)
	})
}

func (suite *postgresStoreTestSuite) TestPostgresStore_FindUserByNameAndPassword() {
	suite.Run("should return user", func() {
		name := uuid.NewString()
		password := uuid.NewString()

		user, _ := suite.store.CreateUser(suite.ctx, name, password)
		u, err := suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		suite.Require().NoError(err)
		suite.Require().NotEmpty(u)
		actualUser, err := suite.store.FindUserByNameAndPassword(suite.ctx, user.Username, password)
		suite.Require().NoError(err)
		suite.Equal(user, actualUser)
	})
	suite.Run("should return error", func() {
		u, err := suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		suite.Require().NoError(err)
		suite.Require().NotEmpty(u)
		u, err = suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		suite.Require().NoError(err)
		suite.Require().NotEmpty(u)
		actualUser, err := suite.store.FindUserByNameAndPassword(suite.ctx, uuid.NewString(), uuid.NewString())

		suite.Require().EqualError(err, ErrUserNotFound.Error())
		suite.Empty(actualUser)
	})
	suite.Run("should return error", func() {
		name := uuid.NewString()
		u, err := suite.store.CreateUser(suite.ctx, name, uuid.NewString())
		suite.Require().NoError(err)
		suite.Require().NotEmpty(u)
		u, err = suite.store.CreateUser(suite.ctx, uuid.NewString(), uuid.NewString())
		suite.Require().NoError(err)
		suite.Require().NotEmpty(u)
		actualUser, err := suite.store.FindUserByNameAndPassword(suite.ctx, name, uuid.NewString())

		suite.Require().EqualError(err, ErrUserNotFound.Error())
		suite.Empty(actualUser)
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
	suite.truncateTable("users")
}
func (suite *postgresStoreTestSuite) truncateTable(tableName string) {
	_, err := suite.pool.Exec(context.Background(), "TRUNCATE TABLE"+" "+tableName+";")
	suite.Require().NoError(err)
}
