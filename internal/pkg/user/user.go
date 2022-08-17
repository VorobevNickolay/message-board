package user

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
