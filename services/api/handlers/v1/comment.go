package v1

import "github.com/gin-gonic/gin"

type Comment struct {
}

func NewComment() Comment {
	return Comment{}
}

func (com Comment) List(c *gin.Context) {

}

func (com Comment) Action(c *gin.Context) {

}
