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
}
func (suite *postgresStoreTestSuite) SetupTest() {
	suite.ctx = context.Background()
	databaseUrl := "postgres://vorobevna:message-board@localhost:15432/postgres"

	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	suite.Require().NoError(err)
	suite.pool = dbPool
	suite.store = newPostgresStore(dbPool)
}
func (suite *postgresStoreTestSuite) AfterTest(_, _ string) {
	suite.truncateTable("users")
}
func (suite *postgresStoreTestSuite) truncateTable(tableName string) {
	_, err := suite.pool.Exec(context.Background(), "TRUNCATE TABLE"+tableName+";")
	suite.Require().NoError(err)
}
