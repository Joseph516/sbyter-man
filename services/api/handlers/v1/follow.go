package v1

import "github.com/gin-gonic/gin"

type Follow struct {
}

func NewFollow() Follow {
	return Follow{}
}

func (com Follow) Action(c *gin.Context) {

}

func (com Follow) FollowerList(c *gin.Context) {

}

func (com Follow) FollowList(c *gin.Context) {

}
