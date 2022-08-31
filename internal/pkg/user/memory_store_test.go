package user

/*
func TestGetUsers(t *testing.T) {
	t.Run("should return empty list", func(t *testing.T) {
		store := NewInMemoryStore()

		actual, err := store.GetUsers()
		require.NoError(t, err)
		require.Equal(t, 0, len(actual))
	})

	t.Run("should return users", func(t *testing.T) {
		store := NewInMemoryStore()
		exp1, err := store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		exp2, err := store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		actual, err := store.GetUsers()

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
		actual, err := store.FindUserById(uuid.NewString())
		require.Error(t, err, ErrUserNotFound)
		require.Equal(t, actual, User{})
	})

	t.Run("should find user", func(t *testing.T) {
		store := NewInMemoryStore()
		_, err := store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		expected, err := store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		_, err = store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		actual, err := store.FindUserById(expected.ID)
		require.Equal(t, expected, actual)
	})
}
func TestFindUserByName(t *testing.T) {
	t.Run("should return ErrUserNotFound", func(t *testing.T) {
		store := NewInMemoryStore()
		actual, err := store.findUserByName(uuid.NewString())
		require.Error(t, err, ErrUserNotFound)
		require.Equal(t, actual, User{})
	})

	t.Run("should find user", func(t *testing.T) {
		store := NewInMemoryStore()
		_, err := store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		expected, err := store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		_, err = store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		actual, err := store.findUserByName(expected.Username)
		require.Equal(t, expected, actual)
	})
}
func TestFindUserByNameAndPassword(t *testing.T) {
	t.Run("should return ErrUserNotFound, no users in store", func(t *testing.T) {
		store := NewInMemoryStore()
		actual, err := store.FindUserByNameAndPassword(uuid.NewString(), uuid.NewString())
		require.Error(t, err, ErrUserNotFound)
		require.Equal(t, actual, User{})
	})
	t.Run("should return ErrUserNotFound, wrong password", func(t *testing.T) {
		store := NewInMemoryStore()
		userName := uuid.NewString()
		_, err := store.CreateUser(userName, uuid.NewString())
		actual, err := store.FindUserByNameAndPassword(userName, uuid.NewString())
		require.Error(t, err, ErrUserNotFound)
		require.Equal(t, actual, User{})
	})

	t.Run("should find user", func(t *testing.T) {
		store := NewInMemoryStore()
		_, err := store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		password := uuid.NewString()
		expected, err := store.CreateUser(uuid.NewString(), password)
		require.NoError(t, err)

		_, err = store.CreateUser(uuid.NewString(), uuid.NewString())
		require.NoError(t, err)

		actual, err := store.FindUserByNameAndPassword(expected.Username, password)
		require.Equal(t, expected, actual)
	})
}
func TestCreateUser(t *testing.T) {
	t.Run("should return errUsedUserName", func(t *testing.T) {
		store := NewInMemoryStore()

		username := uuid.NewString()
		_, err1 := store.CreateUser(username, uuid.NewString())
		actual, err2 := store.CreateUser(username, uuid.NewString())
		require.NoError(t, err1)
		require.Error(t, err2, ErrUsedUsername)
		require.Equal(t, actual, User{})
	})
	t.Run("should create user", func(t *testing.T) {
		store := NewInMemoryStore()

		expected := User{
			Username: uuid.NewString(),
			ID:       uuid.NewString(),
			Password: uuid.NewString(),
		}
		actual, err := store.CreateUser(expected.Username, expected.Password)
		require.NoError(t, err)
		require.Equal(t, expected.Username, actual.Username)
		require.NoError(t, bcrypt.CompareHashAndPassword([]byte(actual.Password), []byte(expected.Password)))
		require.NotEmpty(t, actual.ID)
	})
}
*/
