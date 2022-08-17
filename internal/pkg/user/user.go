package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
