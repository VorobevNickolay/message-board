package app

import (
	"github.com/gin-gonic/gin"
	"message-board/internal/pkg/message"
	"message-board/internal/pkg/user"
	"net/http"
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
}
type Router struct {
	ginContext   *gin.Engine
	messageStore messageStore
	userStore    userStore
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

}
func (r *Router) Run() {
	r.ginContext.Run("localhost:8080")
}
func (r *Router) postMessage(c *gin.Context) {
	var newMessage message.Message
	if err := c.BindJSON(&newMessage); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "add message error")
		return
	}
	m, err := r.messageStore.AddMessage(newMessage)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "add message error")
		return
	}
	c.IndentedJSON(http.StatusCreated, m)
}
func (r *Router) getMessages(c *gin.Context) {
	messages, err := r.messageStore.GetMessages()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "get messages error")
		return
	}
	c.IndentedJSON(http.StatusOK, messages)
}
func (r *Router) getUsers(c *gin.Context) {
	users, _ := r.userStore.GetUsers()
	c.IndentedJSON(http.StatusOK, users)
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
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, u)
}

/*func login(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	isUserFound := user.LoginUser(u)
	if !isUserFound {
		c.JSON(http.StatusUnauthorized, "wrong login or password")
	}
	token, err := jwt.CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, token)
} */
func (r *Router) signUp(c *gin.Context) {
	var newUser user.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	u, err := r.userStore.AddUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "add user error")
		return
	}

	c.IndentedJSON(http.StatusCreated, u)
}
