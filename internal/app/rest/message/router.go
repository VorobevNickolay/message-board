package message

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"message-board/internal/app/rest"
	messagepkg "message-board/internal/pkg/message"
	"net/http"
)

type messageStore interface {
	CreateMessage(ctx context.Context, message messagepkg.Message) (messagepkg.Message, error)
	FindMessageById(ctx context.Context, id string) (messagepkg.Message, error)
	GetMessages(ctx context.Context) ([]*messagepkg.Message, error)
	UpdateMessage(ctx context.Context, id, text string) (messagepkg.Message, error)
	DeleteMessage(ctx context.Context, id string) error
}
type Router struct {
	store messageStore
}

func NewRouter(store messageStore) *Router {
	return &Router{store}
}

func (r *Router) SetUpRouter(engine *gin.Engine) {
	engine.GET("/messages", r.getMessages)
	engine.GET("/message/:id", r.getMessageByID)
	engine.POST("/message", rest.AuthMiddleware(), r.postMessage)
	engine.PUT("/message/:id", rest.AuthMiddleware(), r.updateMessage)
	engine.DELETE("/message/:id", rest.AuthMiddleware(), r.deleteMessage)
}

func (r *Router) postMessage(c *gin.Context) {
	var newMessage messagepkg.Message
	if err := c.BindJSON(&newMessage); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}

	newMessage.UserId = c.GetString("userId")
	m, err := r.store.CreateMessage(c, newMessage)
	if err != nil {
		if errors.Is(err, messagepkg.ErrEmptyMessage) {
			c.IndentedJSON(http.StatusBadRequest, rest.ErrorModel{Error: err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: ErrDataBase.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, m)
}

func (r *Router) updateMessage(c *gin.Context) {
	id := c.Param("id")
	userId := c.GetString("userId")
	oldMessage, err := r.store.FindMessageById(c, id)
	if err != nil {
		if errors.Is(err, messagepkg.ErrMessageNotFound) {
			c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: ErrDataBase.Error()})
		}
		return
	}

	if oldMessage.UserId != userId {
		c.IndentedJSON(http.StatusForbidden, rest.ErrorModel{Error: rest.ErrNoAccess.Error()})
		return
	}

	var newMessage messagepkg.Message
	if err := c.BindJSON(&newMessage); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}
	m, err := r.store.UpdateMessage(c, id, newMessage.Text)
	if err != nil {
		if errors.Is(err, messagepkg.ErrMessageNotFound) {
			c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		} else if errors.Is(err, messagepkg.ErrEmptyMessage) {
			c.IndentedJSON(http.StatusBadRequest, rest.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: ErrDataBase.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}

func (r *Router) deleteMessage(c *gin.Context) {
	id := c.Param("id")
	userId := c.GetString("userId")
	oldMessage, err := r.store.FindMessageById(c, id)
	if err != nil {
		if errors.Is(err, messagepkg.ErrMessageNotFound) {
			c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: ErrDataBase.Error()})
		}
		return
	}

	if oldMessage.UserId != userId {
		c.IndentedJSON(http.StatusForbidden, rest.ErrorModel{Error: rest.ErrNoAccess.Error()})
		return
	}

	err = r.store.DeleteMessage(c, id)
	if err != nil {
		if errors.Is(err, messagepkg.ErrMessageNotFound) {
			c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: ErrDataBase.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "message successfully deleted"})
}

func (r *Router) getMessages(c *gin.Context) {
	messages, err := r.store.GetMessages(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: ErrDataBase.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, messages)
}

func (r *Router) getMessageByID(c *gin.Context) {
	id := c.Param("id")

	m, err := r.store.FindMessageById(c, id)
	if err != nil {
		if errors.Is(err, messagepkg.ErrMessageNotFound) {
			c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: ErrDataBase.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, m)
}