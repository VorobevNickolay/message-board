package app

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"message-board/internal/pkg/jwt"
	"message-board/internal/pkg/message"
	"message-board/internal/pkg/user"
	"net/http"
	"strings"
)

type messageStore interface {
	AddMessage(message message.Message) (message.Message, error)
	FindMessageById(id string) (message.Message, error)
	GetMessages() ([]*message.Message, error)
}

type userStore interface {
	AddUser(user user.User) (user.User, error)
	FindUserById(id string) (user.User, error)
	GetUsers() ([]*user.User, error)
	FindUserByName(name string) (user.User, error)
	UserModelById(id string) user.UserModel
}
type Router struct {
	ginContext   *gin.Engine
	messageStore messageStore
	userStore    userStore
}

type ErrorJson struct {
	error string `json:"error"`
}

func NewRouter(messageStore messageStore, userStore userStore) *Router {
	return &Router{gin.Default(), messageStore, userStore}
}

func (r *Router) SetUpRouter() {
	//r.ginContext.POST("/login", login)

	r.ginContext.GET("/messages", r.getMessages)
	r.ginContext.GET("/message/:id", r.getMessageByID)
	r.ginContext.POST("/message", r.postMessage)

	r.ginContext.GET("/users", r.getUsers)
	r.ginContext.GET("/user/:id", r.getUserByID)
	r.ginContext.POST("/user", r.signUp)
	r.ginContext.POST("/user/login", r.login)

}
func (r *Router) Run() {
	r.ginContext.Run("localhost:8080")
}
func (r *Router) postMessage(c *gin.Context) {
	var newMessage message.Message
	if err := c.BindJSON(&newMessage); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "add message error"})
		return
	}
	m, err := r.messageStore.AddMessage(newMessage)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "add message error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, m)
}
func (r *Router) getMessages(c *gin.Context) {
	messages, err := r.messageStore.GetMessages()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "get messages error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"messages": messages})
}
func (r *Router) getUsers(c *gin.Context) {
	users, _ := r.userStore.GetUsers()
	c.IndentedJSON(http.StatusOK, gin.H{"users": users})
}
func (r *Router) getMessageByID(c *gin.Context) {
	id := c.Param("id")

	m, err := r.messageStore.FindMessageById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, m)
}
func (r *Router) getUserByID(c *gin.Context) {
	id := c.Param("id")

	u, err := r.userStore.FindUserById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	um := r.userStore.UserModelById(u.ID)
	c.IndentedJSON(http.StatusOK, um)
}

func (r *Router) login(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	password := u.Password
	u, err := r.userStore.FindUserByName(u.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return

	}

	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)

}
func (r *Router) signUp(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(newUser.Password) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "password can't be empty"})
		return
	}
	newUser.Password = createHash(newUser.Password)
	newUser.Username = strings.ToLower(newUser.Username)
	u, err := r.userStore.AddUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//todo: add validator
	um := r.userStore.UserModelById(u.ID)
	c.IndentedJSON(http.StatusCreated, um)
}
func createHash(s string) string {
	bytePassword := []byte(s)
	hashPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hashPassword)
}
