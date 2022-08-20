package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username" gorm:"unique"`
}
